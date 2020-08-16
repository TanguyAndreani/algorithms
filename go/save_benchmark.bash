#!/usr/bin/env bash
set -e

function check_args () {
    if [ $# -eq 0 ]
    then
        echo "No arguments supplied"
        exit 1
    fi
}

function script_directory () {
    echo "$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
}

check_args $#

# we want to be in the directory where the script is!
pushd $(script_directory)

make ${1}_benchmark.bin
./${1}_benchmark.bin > ./benchmarks_results/${1}.txt

echo ============
cat ./benchmarks_results/${1}.txt
echo ============

popd
