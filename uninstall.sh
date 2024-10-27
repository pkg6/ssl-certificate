#!/usr/bin/env bash
set -e

worker_path=${HOME}/.ssl-certificate
echo "[3/1] Delete working directory"
rm -rf ${worker_path}

echo "[3/2] Delete environment variables"
if [ -x "$(command -v bash)" ]; then
    grep -v '^export PATH="${HOME}/\.ssl-certificate/bin:$PATH"$' "$HOME/.bashrc" > "$HOME/.bashrc.tmp" && mv "$HOME/.bashrc.tmp" "$HOME/.bashrc"
fi
if [ -x "$(command -v zsh)" ]; then
    grep -v '^export PATH="${HOME}/\.ssl-certificate/bin:$PATH"$' "${HOME}/.zshrc" > "${HOME}/.zshrc.tmp" && mv "${HOME}/.zshrc.tmp" "${HOME}/.zshrc"
fi

echo  "[3/3] Uninstalling completed"
