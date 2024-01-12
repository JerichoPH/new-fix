from typing import Any
from pydantic import BaseModel


class ResponseMdl(BaseModel):
    message_id: str
    trace_id: str
    msg: str
    content: Any
    status: int
    error_code: int
    pagination: Any
    business_type: str
