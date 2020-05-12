#!/bin/bash

docker run -p 8080:8080 dgraph/standalone:v2.0.0-rc1 &
DGRAPH_PID=$!

echo "DGRAPH_PID: " $DGRAPH_PID "\n"

sleep 10 # note: this should be a while loop

curl -X POST localhost:8080/admin/schema -d '@schema.graphql'
RESPONSE_STATUS=$?

echo "RESPONSE_STATUS: " $RESPONSE_STATUS "\n"

# notes:
# [ ] create json strings: https://stackoverflow.com/questions/48470049/build-a-json-string-with-bash-variables
# - [ ] possibly write json string as stdout via echo
# [ ] kill background process by pid:
# - [ ] `pidof <program>` (e.g. `docker`) → PID number
# - [ ] `kill -9 <pid>`
