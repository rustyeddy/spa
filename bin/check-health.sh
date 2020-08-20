#!/bin/sh

host=$1

curl http://{$host}:1233/api/health
