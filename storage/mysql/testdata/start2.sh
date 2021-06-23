#!/bin/bash

set -eux

docker run -d \
	-p 13307:3306 \
	--name schemaless_mariadb2 \
	--env MYSQL_DATABASE=trips \
	--env MYSQL_USER=sltest \
	--env MYSQL_PASSWORD=`cat /home/elysium/.mysql_password` \
	--env MYSQL_ROOT_PASSWORD=`cat /home/elysium/.mysql_root_password` \
	--rm \
	mariadb:10.5

