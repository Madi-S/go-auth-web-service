import redis
from src.config import Config
import typing as t
from functools import wraps


def key_formatter(func) -> t.Callable:
    """Make sure that key is the first keyword or positional argument of the function"""

    @wraps(func)
    def inner(*args: t.Any, **kwargs: t.Any) -> str:
        key = kwargs.get("key") or args[0]
        kwargs["key"] = key
        return func(*args, **kwargs)

    return inner


class RedisClient:
    PREFIX = "blablabla-app"

    def __init__(self, config: Config) -> None:
        self.redis_host = config.get("redis.host")
        self.redis_port = config.get("redis.port")
        self.client = redis.Redis(
            host=self.redis_host, port=self.redis_port, decode_responses=True
        )

    @key_formatter
    def set(self, key: str, value: str) -> None:
        self.client.set(key, value)

    @key_formatter
    def get(self, key: str) -> t.Any:
        return self.client.get(key)
