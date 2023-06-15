#!/bin/sh

/src/kratos serve --config /src/kratos.yml &
caddy run --config /src/Caddyfile --adapter caddyfile
