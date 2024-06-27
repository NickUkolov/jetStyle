from datetime import datetime

from sqlalchemy import Column, Integer, String, DateTime

from db.base import Base


class User(Base):
    __tablename__ = "users"

    id = Column(Integer, primary_key=True, index=True)
    username = Column(String, unique=True, index=True)
    password = Column(String)
    created_at = Column(DateTime, default=datetime.utcnow)
    last_password_change = Column(DateTime, default=datetime.utcnow)
