#!/bin/bash


if [ "${DEBUG}" != "" ]; then
    set -euxo pipefail
else
    set -euo pipefail
fi


sleep infinity

