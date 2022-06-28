import json
import time
import grpc
import customer_account_v2_pb2
import customer_account_v2_pb2_grpc

customer_id = "4633ee97-1892-4a82-8883-618a2d9d7fad"

channel = grpc.insecure_channel('customer-account.prod.svc.cluster.local:9000')
stub = customer_account_v2_pb2_grpc.CustomerAccountStub(channel)
r = stub.GetUserByID(customer_account_v2_pb2.GetUserByIDRequest(
    user_id=customer_id,
))
exp = int(time.time()) + 24 * 60 * 60
user = r.user
res = stub.GenUserJWTToken(customer_account_v2_pb2.GenUserJWTTokenRequest(
    exp=exp,
    payload=json.dumps({
        "dialing_prefix": user.PhonePrefixNumber,
        "exp": exp,
        "phone_number": user.phoneNumber,
        "source": user.source,
        "full_name": user.fullName,
        "user_id": customer_id
    })
))
print(res.jwtToken)
