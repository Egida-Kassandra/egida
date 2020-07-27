from setuptools import setup, find_packages

from os import path
this_directory = path.abspath(path.dirname(__file__))
with open(path.join(this_directory, 'README.md'), encoding='utf-8') as f:
    long_description = f.read()

setup(
    name = "egida",
    version = "1.0.1",
    author = "Antonio Paya Gonzalez",
    author_email = "antoniopaya@outlook.com",
    description = "Egida Project Main",
    long_description=long_description,
    long_description_content_type='text/markdown',
    license = "MIT",
    url = "https://antonioalfa22.github.io/egida/",
    download_url = "https://github.com/antonioalfa22/egida/releases/tag/1.0.1",
    packages=find_packages('egida'),
    entry_points = {
        'console_scripts' : ['egida = egida.egida:main']
    }
)