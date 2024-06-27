import os
from functools import lru_cache

from pydantic_settings import BaseSettings


class Settings(BaseSettings):
    POSTGRES_USER: str = os.environ.get('POSTGRES_USER', 'user')
    POSTGRES_PASSWORD: str = os.environ.get('POSTGRES_PASSWORD', 'admin')
    POSTGRES_DB: str = os.environ.get('POSTGRES_DB', 'microservices_db')
    POSTGRES_HOST: str = os.environ.get('POSTGRES_HOST', 'localhost')
    POSTGRES_PORT: str = os.environ.get('POSTGRES_PORT', '5432')
    POSTGRES_URL: str = f"postgresql://{POSTGRES_USER}:{POSTGRES_PASSWORD}@{POSTGRES_HOST}:{POSTGRES_PORT}/{POSTGRES_DB}"

    SECRET_KEY: str = os.environ.get('SECRET_KEY', "your_secret_key")
    ALGORITHM: str = os.environ.get('ALGORITHM', "HS256")


@lru_cache()
def get_settings():
    return Settings()


settings = get_settings()
