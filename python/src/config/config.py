import yaml


class Config:
    def __init__(self, filename: str):
        with open(filename, "r") as file:
            self.config = yaml.safe_load(file)

    def get(self, key: str, default=None):
        keys = key.split(".")
        value = self.config
        for k in keys:
            value = value.get(k, {})
        return value or default


def get_config() -> Config:
    # TODO: add singleton
    return Config(filename="src/config/dev.yml")


config = get_config()
