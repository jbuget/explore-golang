#!/bin/bash

status_code=$(curl --write-out %{http_code} --silent --output /dev/null -X POST http://localhost/accounts -d "name=titi&email=tutuf@example.org&password=Abcd1234")

if [[ "$status_code" -ne 200 ]] ; then
  echo "Failure" >&2
  exit 1
else
  echo "Success"
  exit 0
fi
