# gvf-server

> The server for the GVF project.

## database

install mysql, and create a database named 'myblog'.

```sh
mysql> create database myblog;
```

## dev

install bee.

```sh
go get github.com/beego/bee
```

set env on windows.

```sh
set QINIU_AK <your qiniu ak>
set QINIU_SK <your qiniu sk>
set QINIU_BUCKET <your qiniu bucket name>
set ADMIN_NAME <admin name>
set ADMIN_NICK <admin nick>
set ADMIN_PASSWORD <admin password>
set ADMIN_EMAIL <admin email>
```

and you can add it in you system environment variables.

set env on linux.

```sh
export QINIU_AK=<your qiniu ak>
export QINIU_SK=<your qiniu sk>
export QINIU_BUCKET=<your qiniu bucket name>
export ADMIN_NAME=<admin name>
export ADMIN_NICK=<admin nick>
export ADMIN_PASSWORD=<admin password>
export ADMIN_EMAIL=<admin email>
```

and you can add it in ~/.bash_profile

run

```sh
bee run
```

debug in vscode.config you launch.json.

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}",
            "env": {},
            "args": []
        }
    ]
}
```

## docker

### build the image

```sh
docker build -t ravenq/gvf-server .
```

or pull the image from DockerHub.

```sh
docker pull ravenq/gvf-server
```

### run with docker-compose

modify the docker-compose.yml for you config:

```sh
QINIU_AK: <your qiniu ak>
QINIU_SK: <your qiniu sk>
QINIU_BUCKET: <your qiniu bucket name>
ADMIN_NAME: <admin name>
ADMIN_NICK: <admin nick>
ADMIN_PASSWORD: <admin password>
ADMIN_EMAIL: <admin email>
```

and then run.

```sh
docker-compose up -d
```