from setuptools import setup, find_packages

setup(
    name="auth_proto",
    version="0.1.0",
    packages=find_packages(),
    install_requires=["grpcio", "protobuf"],
    include_package_data=True,
)
