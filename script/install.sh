#!/bin/bash

remote_repo=https://github.com/ielepro/pandora.git
repo_path=$HOME/.pandora-repo
install_path=$( cd `dirname $0`; pwd )/pandora-deploy

checkCommand() {
    type $1 > /dev/null 2>&1
    if [[ $? -ne 0 ]]; then
        echo "error: $1 must be installed"
        echo "install exit"
        exit 1
    fi
}

checkCommand "go"
checkCommand "git"
checkCommand "make"

if [[ -d ${install_path} ]];then
    install_path=${install_path}-$( date +%Y%m%d%H%M%S )
fi

rm -fr ${repo_path}
git clone ${remote_repo} ${repo_path}

cd ${repo_path}
make

rm -fr ${install_path}
cp -r ${repo_path}/output ${install_path}

rm -fr ${repo_path}

cat << EOF

Installing Pandora Path:  ${install_path}
Install complete.

EOF