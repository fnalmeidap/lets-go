#! /bin/bash

for (( i=0; i<30; i++ ))
do
    go run server_sequential.go &
    sleep 0.1
    go run client.go &

    wait
done

for (( i=0; i<30; i++ ))
do
    go run server_concurrent.go &
    sleep 0.1
    go run client.go &
    wait
done