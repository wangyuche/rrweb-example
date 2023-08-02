



```shell
docker run --name mongo -p 27017:27017 \
-v /$(PWD)/mongo:/data/db \
-e MONGO_INITDB_ROOT_USERNAME=root \
-e MONGO_INITDB_ROOT_PASSWORD=1234 \
-d mongo:6.0
```

