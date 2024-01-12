from typing import Any, Dict
from models.stdBusniess import StdBusniessMdl


class StdBusniess:
    _std_busniess: StdBusniessMdl = None

    def __init__(
        self,
        business_type: str = "",
        content: Dict[str, Any] = {},
        message_id: str = "",
    ) -> None:
        self._std_busniess = StdBusniessMdl(
            message_id=message_id,
            business_type=business_type,
            content=content,
        )

    def to_dict(self) -> dict:
        return self._std_busniess.__dict__

    def to_json_str(self) -> str:
        return self._std_busniess.json()
