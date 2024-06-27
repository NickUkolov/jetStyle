from typing import Callable, Type

from fastapi import Depends
from sqlalchemy.orm import Session

from db.base import get_db
from repositories.common import Repository


def get_repository(repo_type: Type[Repository]) -> Callable:
    def get_repo(db: Session = Depends(get_db)) -> Repository:
        return repo_type(db)
    return get_repo
