#!/bin/bash


if [ "${DEBUG}" != "" ]; then
    set -euxo pipefail
else
    set -euo pipefail
fi


export TMP_SECRET_DIR=/tmp/secrets
mkdir -p ${TMP_SECRET_DIR}

# if file is NOT empty, do...
# if [ ! -s ${SSH_KEY_PRIV_PATH} ]; then
if [[ ! $(cat ${SSH_KEY_PRIV_PATH}) = "" ]]; then
    # cp ${SSH_KEY_PUB_PATH} ${TMP_SECRET_DIR}/$(basename ${SSH_KEY_PUB_PATH})
    cp ${SSH_KEY_PRIV_PATH} ${TMP_SECRET_DIR}/$(basename ${SSH_KEY_PRIV_PATH})
    # chmod -R 644 ${SECRETS_VOLUMN_PATH}
    chmod -R 600 ${TMP_SECRET_DIR}/$(basename ${SSH_KEY_PRIV_PATH})

    sed -i "s|HashKnownHosts yes|HashKnownHosts no|g" /etc/ssh/ssh_config
    eval `ssh-agent`
    # ssh-add ${SECRETS_VOLUMN_PATH}/$(basename ${SSH_KEY_PRIV_PATH})
    ssh-add ${TMP_SECRET_DIR}/$(basename ${SSH_KEY_PRIV_PATH})
fi

