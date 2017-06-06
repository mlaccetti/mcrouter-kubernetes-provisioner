#!/bin/bash

while [ ! -e /mcrouter/mcrouter-config.json ]; do
  echo "waiting for config file";
  sleep 1;
done

/usr/local/bin/mcrouter -p 5000 -f /mcrouter/mcrouter-config.json
