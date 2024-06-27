import datetime

from pydantic import BaseModel, field_validator, Field


class UserUpdate(BaseModel):
    username: str = Field(pattern=r'^[a-zA-Z]+$')

    @field_validator("username", mode="before")
    def convert_list_to_str(cls, v):
        return v.lower()


class UserUpdatePassword(BaseModel):
    password: str


class UserCreate(UserUpdate, UserUpdatePassword):
    pass


class UserOut(UserUpdate):
    id: int
    created_at: datetime.datetime
    last_password_change: datetime.datetime
