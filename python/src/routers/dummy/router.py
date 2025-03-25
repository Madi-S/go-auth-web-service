from fastapi import APIRouter
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


@router.get("/redis")
def dummy_redis_set(payload: DummyRedisPayload):
    redis_client.set(key=payload.key, value=payload.value)
    return {"message": "saved successfuly"}


@router.get("/grpc")
def dummy_grpc(payload: DummyGrpcPayload):
    user_id = grpc_client.register(email=payload.email, password=payload.password)
    token = grpc_client.login(
        email=payload.email, password=payload.password, app_id=config.get("app.id")
    )
    is_admin = grpc_client.is_admin(user_id=user_id)
    return {"user_id": user_id, "token": token, "is_admin": is_admin}
