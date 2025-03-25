import grpc
import typing as t
from protos.gen.sso.sso_pb2_grpc import AuthStub
from protos.gen.sso.sso_pb2 import RegisterRequest, LoginRequest, IsAdminRequest
from src.config import Config


class GrpcClient:
    def __init__(self, config: Config) -> None:
        grpc_host = config.get("grpc.host")
        grpc_port = config.get("grpc.port")

        channel = grpc.insecure_channel(f"{grpc_host}:{grpc_port}")
        self.stub = AuthStub(channel)

    def register(self, email: str, password: str) -> int:
        request = RegisterRequest(email=email, password=password)
        response = self.stub.Register(request)
        return t.cast(int, response.user_id)

    def login(self, email: str, password: str, app_id: int) -> str:
        request = LoginRequest(email=email, password=password, app_id=app_id)
        response = self.stub.Login(request)
        return t.cast(str, response.token)

    def is_admin(self, user_id: int) -> bool:
        request = IsAdminRequest(user_id=user_id)
        response = self.stub.IsAdmin(request)
        return t.cast(bool, response.is_admin)
