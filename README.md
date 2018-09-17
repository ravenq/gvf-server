# gvf-server

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