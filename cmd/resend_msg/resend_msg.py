
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
    requests.post(url=SLACK_CHANNEL_PSP_CS_URL, data=json.dumps(data), headers=headers)


if __name__ == '__main__':
    print('test')
    activation_id = "34140"
    user_id = '14011'
    phone_model = 'Galaxy-M31'
    photo = '2022/08/8f892/photo.jpg'
    bi = '''{"battery":{"states":[{"batteryLevel":0.1899999976158142,"isCharging":false,"timestamp":1661706948376},{"batteryLevel":0.1899999976158142,"isCharging":false,"timestamp":1661706948380},{"batteryLevel":0.1899999976158142,"isCharging":false,"timestamp":1661706948385},{"batteryLevel":0.1899999976158142,"isCharging":false,"timestamp":1661706969211},{"batteryLevel":0.18000000715255737,"isCharging":false,"timestamp":1661706974349},{"batteryLevel":0.18000000715255737,"isCharging":false,"timestamp":1661706978391},{"batteryLevel":0.18000000715255737,"isCharging":false,"timestamp":1661706978398},{"batteryLevel":0.18000000715255737,"isCharging":false,"timestamp":1661706978406},{"batteryLevel":0.18000000715255737,"isCharging":false,"timestamp":1661706999313},{"batteryLevel":0.18000000715255737,"isCharging":false,"timestamp":1661707004380},{"batteryLevel":0.18000000715255737,"isCharging":false,"timestamp":1661707008404},{"batteryLevel":0.18000000715255737,"isCharging":false,"timestamp":1661707008411},{"batteryLevel":0.18000000715255737,"isCharging":false,"timestamp":1661707008421},{"batteryLevel":0.18000000715255737,"isCharging":false,"timestamp":1661707029274},{"batteryLevel":0.18000000715255737,"isCharging":false,"timestamp":1661707034361},{"batteryLevel":0.18000000715255737,"isCharging":false,"timestamp":1661707038435},{"batteryLevel":0.18000000715255737,"isCharging":false,"timestamp":1661707038453},{"batteryLevel":0.18000000715255737,"isCharging":false,"timestamp":1661707038458},{"batteryLevel":0.18000000715255737,"isCharging":false,"timestamp":1661707084045},{"batteryLevel":0.18000000715255737,"isCharging":false,"timestamp":1661707084053},{"batteryLevel":0.18000000715255737,"isCharging":false,"timestamp":1661707084063},{"batteryLevel":0.18000000715255737,"isCharging":false,"timestamp":1661707084073},{"batteryLevel":0.18000000715255737,"isCharging":false,"timestamp":1661707084080},{"batteryLevel":0.18000000715255737,"isCharging":false,"timestamp":1661707114065},{"batteryLevel":0.18000000715255737,"isCharging":false,"timestamp":1661707114071},{"batteryLevel":0.18000000715255737,"isCharging":false,"timestamp":1661707114080},{"batteryLevel":0.18000000715255737,"isCharging":false,"timestamp":1661707114082},{"batteryLevel":0.18000000715255737,"isCharging":false,"timestamp":1661707114091},{"batteryLevel":0.18000000715255737,"isCharging":false,"timestamp":1661707144061},{"batteryLevel":0.18000000715255737,"isCharging":false,"timestamp":1661707144065},{"batteryLevel":0.18000000715255737,"isCharging":false,"timestamp":1661707144069},{"batteryLevel":0.18000000715255737,"isCharging":false,"timestamp":1661707144075},{"batteryLevel":0.18000000715255737,"isCharging":false,"timestamp":1661707144080},{"batteryLevel":0.17000000178813934,"isCharging":false,"timestamp":1661707174098},{"batteryLevel":0.17000000178813934,"isCharging":false,"timestamp":1661707174104},{"batteryLevel":0.17000000178813934,"isCharging":false,"timestamp":1661707174119},{"batteryLevel":0.17000000178813934,"isCharging":false,"timestamp":1661707174125},{"batteryLevel":0.17000000178813934,"isCharging":false,"timestamp":1661707174133},{"batteryLevel":0.17000000178813934,"isCharging":false,"timestamp":1661707204160},{"batteryLevel":0.17000000178813934,"isCharging":false,"timestamp":1661707204168},{"batteryLevel":0.17000000178813934,"isCharging":false,"timestamp":1661707204174},{"batteryLevel":0.17000000178813934,"isCharging":false,"timestamp":1661707204179},{"batteryLevel":0.17000000178813934,"isCharging":false,"timestamp":1661707204183},{"batteryLevel":0.17000000178813934,"isCharging":false,"timestamp":1661707215886},{"batteryLevel":0.17000000178813934,"isCharging":false,"timestamp":1661707219514},{"batteryLevel":0.17000000178813934,"isCharging":false,"timestamp":1661707234100},{"batteryLevel":0.17000000178813934,"isCharging":false,"timestamp":1661707234110},{"batteryLevel":0.17000000178813934,"isCharging":false,"timestamp":1661707234123},{"batteryLevel":0.17000000178813934,"isCharging":false,"timestamp":1661707234132},{"batteryLevel":0.17000000178813934,"isCharging":false,"timestamp":1661707234137},{"batteryLevel":0.17000000178813934,"isCharging":false,"timestamp":1661707245949},{"batteryLevel":0.17000000178813934,"isCharging":false,"timestamp":1661707249799},{"batteryLevel":0.17000000178813934,"isCharging":false,"timestamp":1661707264121},{"batteryLevel":0.17000000178813934,"isCharging":false,"timestamp":1661707264126},{"batteryLevel":0.17000000178813934,"isCharging":false,"timestamp":1661707264137},{"batteryLevel":0.17000000178813934,"isCharging":false,"timestamp":1661707264143},{"batteryLevel":0.17000000178813934,"isCharging":false,"timestamp":1661707264151},{"batteryLevel":0.17000000178813934,"isCharging":false,"timestamp":1661707275958},{"batteryLevel":0.17000000178813934,"isCharging":false,"timestamp":1661707279502},{"batteryLevel":0.1599999964237213,"isCharging":false,"timestamp":1661707339553},{"batteryLevel":0.1599999964237213,"isCharging":false,"timestamp":1661707339556},{"batteryLevel":0.1599999964237213,"isCharging":false,"timestamp":1661707339558},{"batteryLevel":0.1599999964237213,"isCharging":false,"timestamp":1661707339562},{"batteryLevel":0.1599999964237213,"isCharging":false,"timestamp":1661707339566},{"batteryLevel":0.1599999964237213,"isCharging":false,"timestamp":1661707339571},{"batteryLevel":0.1599999964237213,"isCharging":false,"timestamp":1661707339573},{"batteryLevel":0.1599999964237213,"isCharging":false,"timestamp":1661707369559},{"batteryLevel":0.1599999964237213,"isCharging":false,"timestamp":1661707369564},{"batteryLevel":0.1599999964237213,"isCharging":false,"timestamp":1661707369568},{"batteryLevel":0.1599999964237213,"isCharging":false,"timestamp":1661707369573},{"batteryLevel":0.1599999964237213,"isCharging":false,"timestamp":1661707369577},{"batteryLevel":0.1599999964237213,"isCharging":false,"timestamp":1661707369582},{"batteryLevel":0.1599999964237213,"isCharging":false,"timestamp":1661707369585},{"batteryLevel":1,"isCharging":false,"timestamp":1661768980650},{"batteryLevel":1,"isCharging":false,"timestamp":1661769010423},{"batteryLevel":1,"isCharging":false,"timestamp":1661769040375},{"batteryLevel":0.9900000095367432,"isCharging":false,"timestamp":1661769070378},{"batteryLevel":0.9900000095367432,"isCharging":false,"timestamp":1661769100382},{"batteryLevel":0.9900000095367432,"isCharging":false,"timestamp":1661769130363},{"batteryLevel":0.9900000095367432,"isCharging":false,"timestamp":1661769160400},{"batteryLevel":0.9900000095367432,"isCharging":false,"timestamp":1661769190380},{"batteryLevel":0.9800000190734863,"isCharging":false,"timestamp":1661769221054},{"batteryLevel":0.9800000190734863,"isCharging":false,"timestamp":1661769251014},{"batteryLevel":0.9399999976158142,"isCharging":false,"timestamp":1661780665362},{"batteryLevel":0.9399999976158142,"isCharging":false,"timestamp":1661780695275},{"batteryLevel":0.9399999976158142,"isCharging":false,"timestamp":1661780695371},{"batteryLevel":0.9399999976158142,"isCharging":false,"timestamp":1661780725285},{"batteryLevel":0.9399999976158142,"isCharging":false,"timestamp":1661780725387},{"batteryLevel":0.9300000071525574,"isCharging":false,"timestamp":1661780755292},{"batteryLevel":0.9300000071525574,"isCharging":false,"timestamp":1661780755399},{"batteryLevel":0.9300000071525574,"isCharging":false,"timestamp":1661780785306},{"batteryLevel":0.9300000071525574,"isCharging":false,"timestamp":1661780785415},{"batteryLevel":0.9300000071525574,"isCharging":false,"timestamp":1661780815326},{"batteryLevel":0.9300000071525574,"isCharging":false,"timestamp":1661780815433},{"batteryLevel":0.9300000071525574,"isCharging":false,"timestamp":1661780845957},{"batteryLevel":0.9300000071525574,"isCharging":false,"timestamp":1661780845961},{"batteryLevel":0.9300000071525574,"isCharging":false,"timestamp":1661780876509},{"batteryLevel":0.9300000071525574,"isCharging":false,"timestamp":1661780876523},{"batteryLevel":0.9300000071525574,"isCharging":false,"timestamp":1661780905937},{"batteryLevel":0.9300000071525574,"isCharging":false,"timestamp":1661780905941}]},"wifi":{"enabled":false},"bluetooth":{"enabled":false},"gps":{"enabled":true},"location":{"alt":-21.80000114440918,"lat":13.6164523,"long":100.6536014},"uptime":4251}	'''
    send_msg(activation_id, user_id, phone_model, bi, photo)