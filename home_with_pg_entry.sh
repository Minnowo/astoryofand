#!/usr/bin/env bash

# run postgres
/usr/local/bin/docker-entrypoint.sh postgres &

# wait for postgres to be ready
pg_isready > /dev/null

while [ $? -ne 0 ]; do

    sleep 1

    pg_isready > /dev/null

done

# run our app, cwd is important since it uses relative paths
cd /app

exec ./main

