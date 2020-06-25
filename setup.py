import os
from setuptools import setup, find_packages
from egida import egida

setup(
    name = "egida",
    version = "1.0.0",
    author = "Antonio Paya Gonzalez",
    author_email = "antoniopaya@outlook.com",
    description = "Egida Project Main",
    license = "MIT",
    url = "https://antonioalfa22.github.io/egida/",
    packages=find_packages(include=['egida', 'egida.*']),
    install_requires=[
        'antlr4-python3-runtime==4.8',
        'certifi==2020.4.5.1',
        'cffi==1.14.0',
        'chardet==3.0.4',
        'console-menu==0.6.0',
        'cryptography==2.9.2',
        'idna==2.9',
        'Jinja2==2.11.2',
        'MarkupSafe==1.1.1',
        'prompt-toolkit==1.0.14',
        'pycparser==2.20',
        'Pygments==2.6.1',
        'PyInquirer==1.0.3',
        'PyYAML==5.3.1',
        'regex==2020.4.4',
        'requests==2.23.0',
        'six==1.14.0',
        'urllib3==1.25.9',
        'wcwidth==0.1.9',
    ],
    entry_points = {
        'console_scripts' : ['egida = egida.egida:main']
    },
    package_data={'egida': [
        'common/hosts',
        'common/playbook.yml.j2',
        'common/playbook-custom.yml',
        'custom/vars_template.yml'
    ]}
)