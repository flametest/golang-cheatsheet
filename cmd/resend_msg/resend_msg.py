
import requests
import json
import random
from google.cloud import storage, spanner
from matplotlib import pyplot as plt

STORAGE = 'axinan_psp_prod'
PROJECT = 'axinan-prod'
SLACK_CHANNEL_PSP_CS_URL = ''
storage_client = storage.Client(project=PROJECT,)
bucket = storage_client.bucket(STORAGE)
# storage = storage.GoogleCloudStorage(bucket_name=settings.PSP_STORAGE)


def upload(file_name):
    blob = bucket.blob(file_name)
    blob.upload_from_filename(file_name)


def send_msg(activation_id, user_id, phone_model, bi, photo):
    spanner_client = spanner.Client(
        project='axinan-prod',
    )
    instance = spanner_client.instance('axinan-1')
    database = instance.database('insurance-prod')

    instance1 = spanner_client.instance('axinan-1')
    database1 = instance1.database('ax-customer-account-prod')

    activation = None
    user = None
    with database.snapshot() as snapshot:
        results = snapshot.execute_sql(
            "select * from PolicyActivation where Id = '%s'" % activation_id
        )
        for row in results:
            activation = row
            break
    with database1.snapshot() as snapshot1:
        results1 = snapshot1.execute_sql(
            "select * from User where UserId = '%s'" % user_id
        )
        for row1 in results1:
            user = row1
            break
    print(user)
    print(activation)
    images = []
    images.append('https://storage.googleapis.com/{}/croped/{}'.format(STORAGE, photo))
    headers = {'Content-type': 'application/json'}
    reject_reasons = {
        'Blur moved': 'Please hold your phone steady and re-take.',
        'Blur dark': 'The photo looks blur. Please re-take with better lighting condition.',
        'Blur reflection': 'There are reflections on the screen. Please re-take with better lighting condition.',
        'Blur dirty': 'Please clean your mirror and re-take.',
        'Cover': 'Part of the screen is covered. Please re-take and submit again.',
        'Partial': 'We need a picture of the whole phone. Please re-take and submit again.',
        'Likely Crack': 'Something looks like a crack. Please re-take and submit again.',
        'Obvious Crack': 'Sorry, we do not cover cracked screens.',
    }
    data = {
        "text": "Received new activation request.\n "
                "UserId: {}\n "
                "Name: {} {}\n "
                "PremiumType: {}\n "
                "Brand: {}\n "
                "Model: {}\n "
                "CrackOrNot: {}\n"
                "Country: {}\n"
                "Use Promotion: {}\n"
                "Promotion Percentage: {}\n"
                "boot lasting time: {}\n"
                "".format(
            user[0],
            user[5],
            user[6],
            'yearly',
            'Samsung',
            phone_model,
            'unknown',
            user[4],
            '',
            0,
            '',
        ),
        'attachments': [],
    }
    device_info = json.loads(bi)
    if 'wifi' in device_info:
        data['text'] += "-- WIFI: {}\n".format(device_info['wifi'])
    if 'bluetooth' in device_info:
        data['text'] += "-- BlueTooth: {}\n".format(device_info['bluetooth'])
    if 'gps' in device_info:
        data['text'] += "-- GPS: {}\n".format(device_info['gps'])

    battery_analyse_result = []
    if 'battery' in device_info and 'states' in device_info['battery']:
        battery = device_info['battery']['states']
        battery_percentage_list = [z['batteryLevel'] * 100 for z in battery]
        timestamp = [int(z['timestamp']) for z in battery]
        timestamp = [(z-min(timestamp))/1000*1.0/60.0 for z in timestamp]
        image_name = '{}.png'.format(random.randint(100, 10000000))

        def analyse_battery_info(percentage_list, time_list):
            # list in result [start_min, end_min, battery_start, battery_end, drop_ration]
            # e.g. battery drop from 6% to 2% in 4 mins, the data will be [10, 14, 6.0, 2.0, 1.0]
            MAX_TIME_GAP_IN_MIN = 30
            MIN_VALID_TIME_DURATION = 10
            result = []
            if len(percentage_list) != len(time_list):
                return []

            tmp_result = []
            for idx in range(len(percentage_list)):
                percentage = percentage_list[idx]
                current_time = time_list[idx]
                if len(tmp_result) == 0 or current_time - tmp_result[-1][1] >= MAX_TIME_GAP_IN_MIN or percentage > tmp_result[-1][3]:
                    tmp_result.append([current_time, current_time, percentage, percentage])
                else:
                    tmp_result[-1][1] = current_time
                    tmp_result[-1][3] = percentage

            for r in tmp_result:
                if r[1] - r[0] < MIN_VALID_TIME_DURATION:
                    continue
                if r[2] - r[3] <= 0.1:
                    continue
                result.append(r + [(r[2]-r[3])*1.0/(r[1]-r[0])])
            return result

        battery_analyse_result = analyse_battery_info(battery_percentage_list, timestamp)
        print('battery_analyse_result', battery_analyse_result)
        fig, ax = plt.subplots(nrows=1, ncols=1)  # create figure & 1 axis
        ax.plot(timestamp, battery_percentage_list, '+')
        ax.set(xlabel='time in mins', ylabel='battery level')
        ax.set_ylim([-20, 120])
        fig.savefig(image_name)  # save the figure to file
        plt.close(fig)  # close the figure

        with open(image_name, "rb") as in_file:  # opening for [r]eading as [b]inary
            image_data = in_file.read()  # if you only wanted to read 512 bytes, do .read(512)
            upload(image_name)
            images = ['https://storage.googleapis.com/{}/{}'.format(STORAGE, image_name)] + images
            print(images)

    data['text'] += '-- Found {} valid battery drop.\n'.format(len(battery_analyse_result))
    for idx, battery_info in enumerate(battery_analyse_result):
        data['text'] += '-- -- #{}: {} percent/min starting from {}\n'.format(idx+1, round(battery_info[4], 3), round(battery_info[0], 1))
    for idx, image in enumerate(images):
        data['attachments'].append({
            "image_url": image,
            "color": "#3AA3E3",
            "fallback": "You are unable to do this, please contact our developer.",
            "attachment_type": "default",
            "callback_id": "psp-activation",
            "actions": [],
        })
        if idx == 0:
            data['attachments'][-1]['title'] = 'Battery Condition'
        if idx == 1:
            data['attachments'][-1]['title'] = 'Back Camera'
        if idx == 1 and len(images) < 3:
            data['attachments'][-1]['title'] = 'Front Camera'
        if idx == 2:
            data['attachments'][-1]['title'] = 'Front Camera'

    def get_new_attachment():
        return {
            "color": "#3AA3E3",
            "fallback": "You are unable to do this, please contact our developer.",
            "attachment_type": "default",
            "callback_id": "psp-activation",
            "actions": [],
        }
    attachment = get_new_attachment()
    attachment['actions'].append(
        {
            "name": "Accept",
            "text": "Accept",
            "type": "button",
            "value": "accept|{}".format(activation_id),
            "style": "primary",
        })
    count = 0
    for k, v in reject_reasons.items():
        count += 1
        attachment['actions'].append({
            "name": "Reject",
            "text": k,
            "type": "button",
            "style": "danger",
            "value": "reject|{}|{}".format(activation_id, v),
        })
        if count >= 4:
            count = 0
            data['attachments'].append(attachment)
            attachment = get_new_attachment()
    if count > 0:
        data['attachments'].append(attachment)
    print(json.dumps(data))
    r = requests.post(url=SLACK_CHANNEL_PSP_CS_URL, data=json.dumps(data), headers=headers)
    print(r.status_code)
    print(r.text)


if __name__ == '__main__':
    print('test')
    activation_id = "34046"
    user_id = '12272'
    phone_model = 'Galaxy-Note10+'
    photo = '2022/08/a67e3/photo.jpg'
    bi = '''{"battery":{"states":[{"batteryLevel":0.6200000047683716,"isCharging":false,"timestamp":1622862827749},{"batteryLevel":0.6200000047683716,"isCharging":false,"timestamp":1622862857762},{"batteryLevel":0.6200000047683716,"isCharging":false,"timestamp":1622862940466},{"batteryLevel":0.6100000143051147,"isCharging":false,"timestamp":1622862970471},{"batteryLevel":0.6100000143051147,"isCharging":false,"timestamp":1622863000475},{"batteryLevel":0.6100000143051147,"isCharging":false,"timestamp":1622863030514},{"batteryLevel":0.6100000143051147,"isCharging":false,"timestamp":1622863060492},{"batteryLevel":0.6100000143051147,"isCharging":false,"timestamp":1622863090503},{"batteryLevel":0.6100000143051147,"isCharging":false,"timestamp":1622863120508},{"batteryLevel":0.6100000143051147,"isCharging":false,"timestamp":1622863150518},{"batteryLevel":0.6100000143051147,"isCharging":false,"timestamp":1622863180526},{"batteryLevel":0.6100000143051147,"isCharging":false,"timestamp":1622863210537},{"batteryLevel":0.6000000238418579,"isCharging":false,"timestamp":1622863240549},{"batteryLevel":0.6000000238418579,"isCharging":false,"timestamp":1622863362005},{"batteryLevel":0.30000001192092896,"isCharging":false,"timestamp":1622873213737},{"batteryLevel":0.30000001192092896,"isCharging":false,"timestamp":1622873243749},{"batteryLevel":0.30000001192092896,"isCharging":false,"timestamp":1622873273755},{"batteryLevel":0.7799999713897705,"isCharging":true,"timestamp":1622886364844},{"batteryLevel":0.7799999713897705,"isCharging":true,"timestamp":1622886394849},{"batteryLevel":0.20999999344348907,"isCharging":true,"timestamp":1658248459266},{"batteryLevel":0.2199999988079071,"isCharging":true,"timestamp":1658248521519},{"batteryLevel":0.15000000596046448,"isCharging":false,"timestamp":1661138704781},{"batteryLevel":0.15000000596046448,"isCharging":false,"timestamp":1661138734782},{"batteryLevel":0.15000000596046448,"isCharging":false,"timestamp":1661138764784},{"batteryLevel":0.14000000059604645,"isCharging":false,"timestamp":1661138848435},{"batteryLevel":0.14000000059604645,"isCharging":false,"timestamp":1661138878416},{"batteryLevel":0.14000000059604645,"isCharging":false,"timestamp":1661138913546},{"batteryLevel":0.36000001430511475,"isCharging":false,"timestamp":1661154766265}]},"wifi":{"enabled":false},"bluetooth":{"enabled":false},"gps":{"enabled":true},"location":{"alt":4.300000190734863,"lat":6.9972983,"long":100.5050438},"uptime":86196}	'''
    send_msg(activation_id, user_id, phone_model, bi, photo)