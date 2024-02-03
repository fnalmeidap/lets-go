#! /bin/bash

for (( i=0; i<30; i++ ))
do
    go run exercise_02/server.go &
    sleep 0.1
    go run exercise_02/client.go &
    wait
done