from fastapi import APIRouter, Depends, status
from fastapi.security import OAuth2PasswordRequestForm

from db.tables import User
from dependencies.auth import get_current_active_user
from dependencies.db import get_repository
from models.token import Token
from models.users import UserCreate, UserUpdate, UserUpdatePassword, UserOut
from repositories.users import UsersRepository
from services.authentication import AuthException

router = APIRouter()


@router.post("/register", response_model=UserOut, status_code=status.HTTP_201_CREATED)
def register_new_user(user_data: UserCreate,
                      user_repo: UsersRepository = Depends(get_repository(UsersRepository))):
    created_user = user_repo.create(user_data=user_data)
    return created_user


@router.post("/login", response_model=Token, status_code=status.HTTP_200_OK)
def user_login_with_email_and_password(user_repo: UsersRepository = Depends(get_repository(UsersRepository)),
                                       form_data: OAuth2PasswordRequestForm = Depends(OAuth2PasswordRequestForm)):
    user = user_repo.authenticate_user(username=form_data.username, password=form_data.password)
    if not user:
        raise AuthException

    access_token = Token(access_token=user_repo.auth_service.create_access_token(user=user))

    return access_token


@router.get("/me", response_model=UserOut, status_code=status.HTTP_200_OK)
def get_currently_authenticated_user(current_user: User = Depends(get_current_active_user)):
    return current_user


@router.patch("/me", response_model=UserOut, status_code=status.HTTP_202_ACCEPTED)
def update_me(user_data: UserUpdate,
              current_user: User = Depends(get_current_active_user),
              user_repo: UsersRepository = Depends(get_repository(UsersRepository))):
    return user_repo.update(user_data, current_user)


@router.patch("/me/password", response_model=UserOut, status_code=status.HTTP_202_ACCEPTED)
def update_password(user_data: UserUpdatePassword,
                    current_user: User = Depends(get_current_active_user),
                    user_repo: UsersRepository = Depends(get_repository(UsersRepository))):
    return user_repo.update(user_data, current_user)
