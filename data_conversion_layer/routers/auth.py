import json
from fastapi import APIRouter
from pydantic import BaseModel
from databases.redis import Redis
from type.constants import constants
from models.account import AccountMdl
from utils.stdResponse import StdResponse


class LoginForm(BaseModel):
    username: str
    password: str


class SaveAuthForm(BaseModel):
    token: str
    account: AccountMdl


class UpdateAuthForm(BaseModel):
    account: AccountMdl


router = APIRouter()


@router.get("/auth/{token}")  # 获取权鉴信息
async def index(token: str):
    # 读取redis中的权鉴信息
    rds = Redis(constants["redis_database_auth"])
    account = rds.get_value(token)
    if not account:
        return StdResponse().empty("token不存在")
    else:
        account = json.loads(account.decode())
        return StdResponse().success({"account": account})


@router.post("/auth")  # 保存权鉴信息
async def store(save_auth_form: SaveAuthForm):
    # 保存权鉴信息到redis
    rds = Redis(constants["redis_database_auth"])
    rds.set_value(
        save_auth_form.token, save_auth_form.account.json(), ex=constants["ex"]
    )

    # 读取当前用户对应token列表
    tokens = rds.get_value(save_auth_form.account.uuid)
    if not tokens:
        tokens = [save_auth_form.token]
    else:
        tokens = json.loads(tokens.decode())
        if save_auth_form.token not in tokens:
            tokens.append(save_auth_form.token)

    # 回存当前用户对应token列表
    rds.set_value(
        save_auth_form.account.uuid,
        json.dumps(tokens, ensure_ascii=True),
        constants["ex"],
    )
    return StdResponse().success(tokens)


@router.put("/auth/{account_uuid}")  # 更新权鉴数据
async def update(account_uuid: str, update_auth_form: UpdateAuthForm):
    # 读取用户对应token列表
    rds = Redis(constants["redis_database_auth"])
    tokens = rds.get_value(account_uuid)
    if not tokens:
        return StdResponse().empty("用户数据不存在")
    tokens = json.loads(tokens.decode())

    # 修改所有对应token的用户信息
    for token in tokens:
        rds.set_value(token, update_auth_form.account.json(), constants["ex"])

    return StdResponse().success()


@router.delete("/auth/{account_uuid}")  # 删除权鉴数据
async def destroy(account_uuid: str):
    # 读取用户对应token列表
    rds = Redis(constants["redis_database_auth"])
    tokens = rds.get_value(account_uuid)
    if not tokens:
        return StdResponse().empty("用户数据不存在")
    tokens = json.loads(tokens.decode())

    # 删除所有对应token的用户信息
    rds.delete_values(tokens)

    # 删除用户权鉴数据
    rds.delete_value(account_uuid)
