from pydantic import BaseModel


class Token(BaseModel):
    access_token: str
    token_type: str = 'bearer'


class JWT(BaseModel):
    username: str
    id: int
