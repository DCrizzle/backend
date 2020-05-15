#!/bin/bash

function dgraph_cleanup()
{
	echo && echo "starting cleanup"
	docker kill $DGRAPH_ID > /dev/null
	exit
}

trap dgraph_cleanup SIGINT ERR

DGRAPH_ID=$(docker run -d -p 8080:8080 dgraph/standalone:v2.0.0-rc1)

echo "dgraph container id: " $DGRAPH_ID # temp

sleep 3 # note: maybe not needed

curl -X POST localhost:8080/admin/schema -d '@database/schema.graphql' > /dev/null
while [ $? -ne 0 ]
do
	echo "retrying schema upload"
	sleep 3
	!!
done
echo "schema upload status: " $?

go build -o backend
echo "compiled backend binary"
./backend > /dev/null
