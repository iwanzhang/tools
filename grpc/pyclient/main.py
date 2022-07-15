import grpc

import helloworld_pb2
import helloworld_pb2_grpc


def run():
    # 配置grpc服务端地址
    channel = grpc.insecure_channel('localhost:50051')
    stub = helloworld_pb2_grpc.GreeterStub(channel)
    # 请求服务端的SayHello方法
    response = stub.SayHello(helloworld_pb2.HelloRequest(name='you'))
    print("Greeter client received: " + response.message)
    # 请求服务端的SayHelloAgain方法
    # response = stub.SayHelloAgain(helloworld_pb2.HelloRequest(name='you'))
    # print("Greeter client received: " + response.message)


if __name__ == '__main__':
    run()
