from typing import List, Union

from .baseMdl import BaseMdl


class RbacPermissionMdl(BaseMdl):
    name: str
    method: str
    uri: str
    rbac_role: Union["RbacRoleMdl", None] = None


class RbacRoleMdl(BaseMdl):
    name: str
    rbac_permissions: List[RbacPermissionMdl] = []
