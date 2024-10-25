#!/usr/bin/env bash
set -e

worker_path=${HOME}/.ssl-certificate

context="export PATH="\${HOME}/.ssl-certificate/bin:\$PATH""

rm -rf ${worker_path}

envfile=""
if [ -x "$(command -v bash)" ]; then
    envfile=${HOME}/.bashrc
fi

if [ -x "$(command -v zsh)" ]; then
    envfile=${HOME}/.zshrc
fi

echo  "Manually remove code \nOn File: \"${envfile}\" \n${context}"
