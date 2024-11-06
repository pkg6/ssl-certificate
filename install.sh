#!/usr/bin/env bash
set -e

gh_user="pkg6"
gh_repos="ssl-certificate"
worker_path=${HOME}/.ssl-certificate
ssl_certificate_commands=("ssl-certificate" "ssl-certificate-local")
env_content='export PATH="${HOME}/.ssl-certificate/bin:$PATH"'

# Function: Get the current system architecture
function liunx_arch() {
    case $(uname -m) in
        "x86_64" | "amd64")
            echo "amd64" ;;  # 64-bit x86 architecture
        "i386" | "i486" | "i586")
            echo "386" ;;  # 32-bit x86 architecture
        "aarch64" | "arm64")
            echo "arm64" ;;  # 64-bit ARM architecture
        "armv6l" | "armv7l")
            echo "arm" ;;  # 32-bit ARM architecture
        "s390x")
            echo "s390x" ;;  # IBM Z series architecture
        *)
            echo "unsupported" ;;  # Unsupported architecture
    esac
}

# Function: Get the current operating system name
function liunx_os() {
    # Convert the OS name to lowercase and output it
    echo $(uname -s | awk '{print tolower($0)}')
}

function check_command(){
    commands=("curl" "jq")
    for cmd in "${commands[@]}"
    do
        if ! command -v ${cmd} &> /dev/null; then
            echo "Please install: ${cmd}"
            exit 1
        fi
    done
}

function gh_version(){
   echo $(curl --silent "https://api.github.com/repos/${gh_user}/${gh_repos}/tags" | jq -r '.[0].name')
}

function gh_url(){
    echo "https://github.com/${gh_user}/${gh_repos}/releases/download/$1/ssl-certificate_$1-$2-$3.tar.gz"
}

function worker_tar_file(){
    echo "${worker_path}/downloads/$1.$2-$3.tar.gz";
}
# Check the environment
check_command

# Define variables
arch=$(liunx_arch)
os=$(liunx_os)
version=$(gh_version)
download_url=$(gh_url ${version} ${os} ${arch})
tarfile=$(worker_tar_file ${version} ${os} ${arch})

rm -rf ${worker_path}

# download
echo "[1/3] Downloading ${download_url}"
curl -s -S -L --create-dirs -o "${tarfile}" "${download_url}"

# decompression
bin_path=${worker_path}/bin
echo "[2/3] Install g to the ${bin_path}"
rm -rf ${bin_path}
mkdir -p ${bin_path}
tar -xz -f  "${tarfile}" -C "${bin_path}"
for ssl_certificate_command in "${ssl_certificate_commands[@]}"
do
    chmod +x "${bin_path}/${ssl_certificate_command}"
done

# Set environment variables
echo "[3/3] Set environment variables"
# Check for bash
if [ -x "$(command -v bash)" ]; then
    if ! grep -qF "${env_content}" "${HOME}/.bashrc"; then
        echo "${env_content}" >> "${HOME}/.bashrc"
    fi
fi

# Check for zsh
if [ -x "$(command -v zsh)" ]; then
    if ! grep -qF "${env_content}" "${HOME}/.zshrc"; then
        echo "${env_content}" >> "${HOME}/.zshrc"
    fi
fi

echo "ssl-certificate Installation successful"

exit 0
