import bcrypt
import jwt
from fastapi import HTTPException, status
from pydantic import ValidationError

from db.tables import User
from models.token import JWT
from settings import settings


class AuthException(HTTPException):
    def __init__(self):
        super().__init__(status_code=status.HTTP_401_UNAUTHORIZED,
                         detail="Could not validate credentials",
                         headers={"WWW-Authenticate": "Bearer"})


class Authentication:

    @staticmethod
    def get_password_hash(password: str) -> str:
        return bcrypt.hashpw(password.encode('utf-8'), bcrypt.gensalt()).decode('utf-8')

    @staticmethod
    def verify_password(password: str, hashed_password: str) -> bool:
        return bcrypt.checkpw(password.encode('utf-8'), hashed_password.encode('utf-8'))

    @staticmethod
    def create_access_token(user: User) -> str:
        return jwt.encode(JWT(username=user.username, id=user.id).model_dump(), settings.SECRET_KEY, algorithm=settings.ALGORITHM)

    @staticmethod
    def get_username_from_token(token: str) -> str:
        try:
            decoded_token = jwt.decode(token, settings.SECRET_KEY, algorithms=[settings.ALGORITHM])
            payload = JWT(**decoded_token)
        except (jwt.PyJWTError, ValidationError):
            raise AuthException
        return payload.username
