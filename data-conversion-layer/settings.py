# config.py
from pydantic import BaseSettings

class Settings(BaseSettings):
    base_url: str = "http://localhost:18081"
    api_prefix: str = "/api"

def get_url(path: str) -> str:
    return f"{Settings.base_url}/{Settings.api_prefix}{path}"