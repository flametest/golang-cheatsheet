import json

import grpc
import customer_account_v2_pb2
import customer_account_v2_pb2_grpc

channel = grpc.insecure_channel('customer-account.prod.svc.cluster.local:9000')
stub = customer_account_v2_pb2_grpc.CustomerAccountStub(channel)
res = stub.GenUserJWTToken(customer_account_v2_pb2.GenUserJWTTokenRequest(
    exp=1646711154,
    payload=json.dumps({
        "dialing_prefix": "60",
        "exp": 1646711154,
        "phone_number": "192886070",
        "source": "Shopee-Malaysia",
        "user_id": "3107b925-03cd-46b8-b3ee-90e8a17d116b"
    })
))
print(res)
