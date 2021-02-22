#!/usr/bin/env bash

set -euo pipefail

me="$( basename "${BASH_SOURCE[0]}" )"
repo_root="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"

args="get pods"
api_version=client.authentication.k8s.io/v1beta1
exec_plugin=/tmp/sample-exec-plugin
kubectl=kubectl
skip_user=false
quiet=false
user=sample-exec-plugin

function usage() {
  echo "usage: $me [-ahknoquv]"
  echo "  -a      set kubectl args (default: $args)"
  echo "  -h      print this usage"
  echo "  -k      set kubectl to use (default: $kubectl)"
  echo "  -n      skip setting user (default: $skip_user)"
  echo "  -o      set output build file (default: $exec_plugin)"
  echo "  -q      set exec plugin to be quiet (default: $quiet)"
  echo "  -u      set kubeconfig exec plugin user (default: $user)"
  echo "  -v      set exec api version (default: $api_version)"
  exit 1
}

while getopts a:hk:no:qu:v: o
do case "$o" in
     a) args="$OPTARG" ;;
     h) usage ;;
     k) kubectl="$OPTARG" ;;
     n) skip_user="true" ;;
     o) exec_plugin="$OPTARG" ;;
     q) quiet="true" ;;
     u) user="$OPTARG" ;;
     v) api_version="$OPTARG" ;;
     [?]) usage ;;
esac
done

cd "$repo_root"

go build -o "$exec_plugin"

if [[ "$skip_user" != "true" ]]; then
  "$kubectl" config set-credentials "$user" \
    --exec-command="$exec_plugin" \
    --exec-api-version="$api_version" \
    --exec-env=QUIET="$quiet"
fi

"$kubectl" --user "$user" ${args}
