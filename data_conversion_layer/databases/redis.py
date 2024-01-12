import redis
from redis.commands.core import ResponseT

class Redis:
    def __init__(self, database: int = 0) -> None:
        self._rds = redis.Redis(
            host="localhost", port=6379, password="", db=database
        )
        self._database = database
        

    def set_value(
        self,
        key: str,
        value: str,
        ex=None,
    ) -> "Redis":
        self._rds.set(name=key, value=value, ex=ex)
        return self

    def get_value(self, key: str) -> ResponseT:
        return self._rds.get(name=key)
