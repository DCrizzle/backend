#!/bin/bash

function dgraph_cleanup()
{
	echo "STARTING CLEANUP"
	docker kill $DGRAPH_ID
	exit
}

trap dgraph_cleanup SIGINT

DGRAPH_ID=$(docker run -d -p 8080:8080 dgraph/standalone:v2.0.0-rc1)

sleep 3 # note: maybe not needed

curl -X POST localhost:8080/admin/schema -d '@database/schema.graphql' # note: possibly write out to null
while [ $? -ne 0 ]
do
	sleep 3
	!!
done

echo $DGRAPH_ID # temp

echo "STARTING SLEEP" # temp

sleep 100 # temp



# outline:
# [ ] build main binary
# [ ] run main binary
# - [ ] create "blocking" call to keep script running
