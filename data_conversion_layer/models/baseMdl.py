from typing import Union
from pydantic import BaseModel


class BaseMdl(BaseModel):
    created_at: str
    updated_at: str
    deleted_at: Union[str, None] = None
    uuid: str
    sort: int
