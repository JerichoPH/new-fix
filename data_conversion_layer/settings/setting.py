import configparser
import os

def dir(filename: str):
    return os.path.join(os.path.dirname(os.path.abspath(__file__)), filename)


def get(filename: str, encoding: str = "utf-8", is_absolute: bool = False):
    if is_absolute is False:
        filename = os.path.join(os.path.dirname(os.path.abspath(__file__)), filename)
    cf = configparser.ConfigParser()
    cf.read(filename, encoding=encoding)
    return cf
