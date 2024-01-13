from pydantic import BaseModel


class StdBusniessMdl(BaseModel):
    message_id: str|None
    trace_id: str|None
    business_type: str
    content: dict | None
