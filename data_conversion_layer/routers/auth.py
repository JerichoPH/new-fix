import json
from fastapi import APIRouter
from pydantic import BaseModel
from databases.redis import Redis
from type.constants import constants
from models.account import AccountMdl
from utils.response import Response


class LoginForm(BaseModel):
    username: str
    password: str


class SaveAuthForm(BaseModel):
    token: str
    account: AccountMdl


router = APIRouter()


@router.post("/auth/login")
async def login(login_form: LoginForm):
    rds = Redis(constants["redis_database_auth"])
    rds.set_value(key=login_form.username, value=login_form.password, ex=60)
    return {"message": rds.get_value(key=login_form.username)}


@router.post("/auth/saveAuth")  # 保存权鉴信息
async def save_auth(save_auth_form: SaveAuthForm):
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
    return Response().success(tokens)
