#!/bin/bash

set -v

promtail-linux-amd64 -config.file /tmp/promtail-config.yml &

exec "$@"