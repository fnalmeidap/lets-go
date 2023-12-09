#! /bin/bash

for (( i=0; i<30; i++ ))
do
    go run server_sequential_v1.go &
    sleep 0.1
    go run client.go &

    wait
done

echo "================="

for (( i=0; i<30; i++ ))
do
    go run server_concurrent_v1.go &
    sleep 0.1
    go run client.go &
    wait
done