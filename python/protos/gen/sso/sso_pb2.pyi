"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""

import builtins
import google.protobuf.descriptor
import google.protobuf.message
import typing

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

@typing.final
class RegisterRequest(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    EMAIL_FIELD_NUMBER: builtins.int
    PASSWORD_FIELD_NUMBER: builtins.int
    email: builtins.str
    password: builtins.str
    def __init__(
        self,
        *,
        email: builtins.str = ...,
        password: builtins.str = ...,
    ) -> None: ...
    def ClearField(
        self, field_name: typing.Literal["email", b"email", "password", b"password"]
    ) -> None: ...

global___RegisterRequest = RegisterRequest

@typing.final
class RegisterResponse(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    USER_ID_FIELD_NUMBER: builtins.int
    user_id: builtins.int
    def __init__(
        self,
        *,
        user_id: builtins.int = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing.Literal["user_id", b"user_id"]) -> None: ...

global___RegisterResponse = RegisterResponse

@typing.final
class LoginRequest(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    EMAIL_FIELD_NUMBER: builtins.int
    PASSWORD_FIELD_NUMBER: builtins.int
    APP_ID_FIELD_NUMBER: builtins.int
    email: builtins.str
    password: builtins.str
    app_id: builtins.int
    def __init__(
        self,
        *,
        email: builtins.str = ...,
        password: builtins.str = ...,
        app_id: builtins.int = ...,
    ) -> None: ...
    def ClearField(
        self,
        field_name: typing.Literal[
            "app_id", b"app_id", "email", b"email", "password", b"password"
        ],
    ) -> None: ...

global___LoginRequest = LoginRequest

@typing.final
class LoginResponse(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    TOKEN_FIELD_NUMBER: builtins.int
    token: builtins.str
    def __init__(
        self,
        *,
        token: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing.Literal["token", b"token"]) -> None: ...

global___LoginResponse = LoginResponse

@typing.final
class IsAdminRequest(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    USER_ID_FIELD_NUMBER: builtins.int
    user_id: builtins.int
    def __init__(
        self,
        *,
        user_id: builtins.int = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing.Literal["user_id", b"user_id"]) -> None: ...

global___IsAdminRequest = IsAdminRequest

@typing.final
class IsAdminResponse(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    IS_ADMIN_FIELD_NUMBER: builtins.int
    is_admin: builtins.bool
    def __init__(
        self,
        *,
        is_admin: builtins.bool = ...,
    ) -> None: ...
    def ClearField(
        self, field_name: typing.Literal["is_admin", b"is_admin"]
    ) -> None: ...

global___IsAdminResponse = IsAdminResponse
