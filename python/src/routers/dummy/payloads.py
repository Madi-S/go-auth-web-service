from pydantic import BaseModel


class DummyRedisPayload(BaseModel):
    key: str
    value: str


class DummyGrpcPayload(BaseModel):
    email: str
    password: str
    app_id: int
