from fastapi import APIRouter
from typing import Union
from pydantic import BaseModel
import requests


class LoginForm(BaseModel):
    username: str
    password: str


router = APIRouter()


@router.post("/auth/login")
async def login(login_form: LoginForm):
    resp = requests.post(
        "http://127.0.0.1:18080/api/auth/login",
        json=login_form.__dict__,
        headers={"Content-Type": "application/json"},
    )
    return resp.json()
