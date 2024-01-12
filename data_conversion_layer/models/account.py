from typing import List, Union
from .baseMdl import BaseMdl
from .rbac import RbacRoleMdl


class AccountMdl(BaseMdl):
    be_enable: bool
    username: str
    nickname: str
    be_admin: bool
    avatar_uuid: str | None = None
    avatar: None
    rbac_roles: List[RbacRoleMdl] = []
    breakdown_logs_by_report_person: None
    warehouse_scan_items_by_processor: None
    warehouse_orders_by_processor: None
