from fastapi import APIRouter
from typing import Union
from pydantic import BaseModel
import requests
from .settings import get_url


class LoginForm(BaseModel):
    username: str
    password: str


router = APIRouter()


@router.post("/auth/login")
async def login(login_form: LoginForm):
    resp = requests.post(
        get_url("/auth/login"),
        json=login_form.__dict__,
        headers={"Content-Type": "application/json"},
    )
    return resp.json()
