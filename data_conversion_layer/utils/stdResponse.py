from fastapi import HTTPException
from urllib3 import HTTPResponse
from models.stdResponse import StdResponseMdl
from models.stdBusniess import StdBusniessMdl
from typing import Any, Dict


class StdResponse:
    _message_id = ""
    _trace_id = ""
    _msg = ""
    _content = None
    _status = 200
    _error_code = 0
    _pagination = None
    _business_type = ""

    def success(self, content=None, pagination=None) -> dict:
        self._msg = "success"
        self._content = content
        self._page = pagination
        return self.dict()

    def empty(self, msg="not found") -> dict:
        self._msg = msg
        self._status = 404
        self._error_code = 2
        return self.dict()

    def forbidden(self, msg="forbidden") -> dict:
        self._msg = msg
        self._status = 403
        self._error_code = 1
        return self.dict()

    def dict(self) -> dict:
        return StdResponseMdl(
            message_id=self._message_id,
            trace_id=self._trace_id,
            msg=self._msg,
            content=self._content,
            status=self._status,
            error_code=self._error_code,
            pagination=self._pagination,
            business_type=self._business_type,
        ).__dict__
