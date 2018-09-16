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

### run docker container

run a mysql container as data container.

```sh
docker run -d --name db -e MYSQL_ROOT_PASSWORD=my-secret-pw mysql
```

```sh
docker run -d --name gvf-server --link db:db \
  -p 8080:8080 \
  -e DB_MYSQL: db \
  -e QINIU_AK: <your qiniu ak> \
  -e QINIU_SK: <your qiniu sk> \
  -e QINIU_BUCKET: <your qiniu bucket name> \
  -e ADMIN_NAME: <admin name> \
  -e ADMIN_NICK: <admin nick> \
  -e ADMIN_PASSWORD: <admin password> \
  -e ADMIN_EMAIL: <admin email> \
  -e RUNMODE: prod \
  ravenq/gvf-server
```

### run with docker-compose

```sh
docker-compose up -d
```