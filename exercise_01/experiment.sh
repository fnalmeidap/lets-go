#! /bin/bash

# for (( i=0; i<30; i++ ))
# do
#     go run server_concurrent.go &
#     go run client.go &
#     echo $i
#     wait
# done

for (( i=0; i<30; i++ ))
do
    go run server_sequential.go &
    go run client.go &

    wait
done