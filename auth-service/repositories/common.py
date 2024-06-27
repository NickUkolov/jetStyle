from abc import ABCMeta, abstractmethod

from sqlalchemy.orm import Session


class AbstractRepository(metaclass=ABCMeta):

    @abstractmethod
    def get_all(self, *args, **kwargs):
        ...

    @abstractmethod
    def get_one(self, *args, **kwargs):
        ...

    @abstractmethod
    def create(self, *args, **kwargs):
        ...

    @abstractmethod
    def update(self, *args, **kwargs):
        ...

    @abstractmethod
    def delete(self, *args, **kwargs):
        ...


class Repository(AbstractRepository, metaclass=ABCMeta):

    def __init__(self, session: Session):
        self.session = session
