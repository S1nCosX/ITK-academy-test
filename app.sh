#!/bin/bash

if [ -n "$1" ]
then
    if $1 == "--start"
    then
        if $2 == "-b" 
        then
        docker compose --env-file ./config.env up --build
        else
            docker compose --env-file ./config.env up
    else if $1 == "--restart"
    then
    docker compose down
        if $2 == "-b" 
        then
        docker compose --env-file ./config.env up --build
        else
            docker compose --env-file ./config.env up
    else if $1 == "--stop"
    then
    docker compose down
else