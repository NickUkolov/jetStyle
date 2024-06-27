from typing import Optional

from fastapi import Depends
from fastapi.security import OAuth2PasswordBearer

from db.tables import User
from dependencies.db import get_repository
from repositories.users import UsersRepository
from services import auth_service
from services.authentication import AuthException

oauth2_scheme = OAuth2PasswordBearer(tokenUrl="login")


def get_user_from_token(
        token: str = Depends(oauth2_scheme),
        user_repo: UsersRepository = Depends(get_repository(UsersRepository)),
) -> Optional[User]:
    try:
        username = auth_service.get_username_from_token(token=token)
        user = user_repo.get_one(username=username)
    except Exception as e:
        raise e

    return user


def get_current_active_user(current_user: User = Depends(get_user_from_token)) -> Optional[User]:
    if not current_user:
        raise AuthException
    return current_user
