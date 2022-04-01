#!/usr/bin/env sh

set -e

#
# set up host.docker.internal for a linux machine
#

IP=$(ip -4 route list match 0/0 | awk '{print $3}')
echo "Host ip is $IP"
echo "$IP   host.docker.internal" >> /etc/hosts

#
# the real entry point of the envoy
#

echo "ARGUMENTS $@"

sh /docker-entrypoint.sh "$@"
