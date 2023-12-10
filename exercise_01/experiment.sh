#! /bin/bash

CONNECTION_TYPE="multiple" # single | multiple

echo "Experiment with $CONNECTION_TYPE" connections.
for (( i=0; i<30; i++ ))
do
    go run ${CONNECTION_TYPE}_connection/server_sequential.go &
    sleep 0.1
    go run ${CONNECTION_TYPE}_connection/client.go &

    wait
done

echo "================="

for (( i=0; i<30; i++ ))
do
    go run ${CONNECTION_TYPE}_connection/server_concurrent.go &
    sleep 0.1
    go run ${CONNECTION_TYPE}_connection/client.go &
    wait
done