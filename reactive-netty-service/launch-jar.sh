#!/bin/bash
set -e -x

opts="-Djava.security.egd=file:/dev/./urandom"

[[ ! -z $JAVA_OPTS ]] && opts="${opts} $JAVA_OPTS"

exec java $opts -jar $@