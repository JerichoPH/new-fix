from models.stdBusniess import StdBusniessMdl


class StdBusniess:
    _std_busniess: StdBusniessMdl = None

    def __init__(
        self,
        business_type: str = "",
        content: dict = {},
        message_id: str = "",
        trace_id: str = "",
    ):
        self._std_busniess = StdBusniessMdl(
            message_id=message_id,
            business_type=business_type,
            content=content,
            trace_id=trace_id,
        )

    def to_dict(self) -> dict:
        return self._std_busniess.__dict__

    def to_json_str(self) -> str:
        return self._std_busniess.json()
