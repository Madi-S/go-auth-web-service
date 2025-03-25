from fastapi import APIRouter
import grpc
import grpc._channel
from src.routers.dummy.payloads import DummyGrpcPayload, DummyRedisPayload
from src.config import config
from src.services import RedisClient, GrpcClient

router = APIRouter(prefix="/example", tags=["Example"])
redis_client = RedisClient(config)
grpc_client = GrpcClient(config)


@router.get("/redis/{key}")
def dummy_redis_get(key: str):
    value = redis_client.get(key=key)
    return {"message": str(value)}


@router.post("/redis")
def dummy_redis_set(payload: DummyRedisPayload):
    redis_client.set(key=payload.key, value=payload.value)
    return {"message": "saved successfuly"}


@router.post("/grpc")
def dummy_grpc(payload: DummyGrpcPayload):
    try:
        user_id = grpc_client.register(email=payload.email, password=payload.password)
        token = grpc_client.login(
            email=payload.email, password=payload.password, app_id=config.get("app.id")
        )
        is_admin = grpc_client.is_admin(user_id=user_id)
        return {"user_id": user_id, "token": token, "is_admin": is_admin}
    except grpc._channel._InactiveRpcError as e:
        return {"message": f"grpc error {e.code()}", "details": e.details()}
