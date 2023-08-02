from concurrent.futures import ThreadPoolExecutor
import grpc

from helloWorld_pb2 import HelloWorldResponse
from helloWorld_pb2_grpc import HelloWorldServicer, add_HelloWorldServicer_to_server


def print_hello_world():
    return "Hello Meat"


class HelloWorldServer(HelloWorldServicer):
    def Print(self, request, context):
        hello = request.request
        resp = HelloWorldResponse(response=hello)
        return resp


if __name__ == "__main__":
    server = grpc.server(ThreadPoolExecutor())
    add_HelloWorldServicer_to_server(HelloWorldServer(), server)
    port = 9999
    server.add_insecure_port(f"[::]:{port}")
    server.start()
    # logging.info('server ready on port %r', port)
    server.wait_for_termination()
