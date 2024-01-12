from typing import Any, Dict
from pydantic import BaseModel


class StdBusniessMdl(BaseModel):
    message_id: str
    trace_id: str
    business_type: str
    content: Dict[str, Any]
