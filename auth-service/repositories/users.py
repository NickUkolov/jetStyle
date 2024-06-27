from typing import Sequence, Optional

from fastapi import status, HTTPException
from sqlalchemy import select
from sqlalchemy.orm import Session

from db.tables import User
from models.users import UserUpdate, UserCreate, UserUpdatePassword
from repositories.common import Repository
from services import auth_service


class UsersRepository(Repository):

    def __init__(self, db: Session) -> None:
        super().__init__(db)
        self.auth_service = auth_service

    def get_all(self) -> Sequence[User]:
        pass

    def get_one(self, username: str) -> User:
        result = self.session.execute(select(User).where(User.username == username).limit(1))
        return result.scalars().first()

    def create(self, user_data: UserCreate) -> User:
        if self.get_one(username=user_data.username):
            raise HTTPException(
                status_code=status.HTTP_400_BAD_REQUEST,
                detail="Username already exists")
        user_data.password = self.auth_service.get_password_hash(password=user_data.password)
        db_user = User(**user_data.model_dump())
        self.session.add(db_user)
        self.session.commit()
        self.session.refresh(db_user)
        return db_user

    def update(self, user_data: UserUpdate | UserUpdatePassword, user: User) -> User:
        if isinstance(user_data, UserUpdatePassword):
            user_data.password = self.auth_service.get_password_hash(password=user_data.password)
        elif isinstance(user_data, UserUpdate):
            if self.get_one(username=user_data.username):
                raise HTTPException(
                    status_code=status.HTTP_400_BAD_REQUEST,
                    detail="Username already exists")

        for key, value in user_data.model_dump().items():
            setattr(user, key, value)
        self.session.commit()
        self.session.refresh(user)
        return user

    def delete(self, user: User) -> None:
        pass

    def authenticate_user(self, username: str, password: str) -> Optional[User]:
        db_user = self.get_one(username)
        if not db_user:
            return None
        if not self.auth_service.verify_password(password=password, hashed_password=db_user.password):
            return None
        return db_user
