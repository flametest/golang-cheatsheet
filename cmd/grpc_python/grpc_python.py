import json
import time
import grpc
import customer_account_v2_pb2
import customer_account_v2_pb2_grpc

customer_id = "a3845c5c-5a03-4551-be6d-49e07d63a201"

channel = grpc.insecure_channel('customer-account.prod.svc.cluster.local:9000')
stub = customer_account_v2_pb2_grpc.CustomerAccountStub(channel)
# r = stub.DeleteUser(customer_account_v2_pb2.DeleteUserRequest(
#     user_id=customer_id
# ))
# print(r)
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
