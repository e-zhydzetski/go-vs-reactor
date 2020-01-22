#!/bin/bash
set -e -x

opts="-Djava.security.egd=file:/dev/./urandom"

[[ ! -z $JAVA_OPTS ]] && opts="${opts} $JAVA_OPTS"

jmxHostPort="${JMX_HOST_PORT// }"
if [ ! -z $jmxHostPort ]; then
  jmxHostPortParts=(${jmxHostPort//:/ })
  if [ ${#jmxHostPortParts[@]} == 2 ]; then
    jmxHost=${jmxHostPortParts[0]}
    jmxPort=${jmxHostPortParts[1]}
    opts="${opts} -Dcom.sun.management.jmxremote=true -Dcom.sun.management.jmxremote.authenticate=false -Dcom.sun.management.jmxremote.ssl=false -Dcom.sun.management.jmxremote.port=${jmxPort} -Dcom.sun.management.jmxremote.rmi.port=${jmxPort} -Djava.rmi.server.hostname=${jmxHost}"
  else
    echo "Invalid JMX_HOST_PORT: ${jmxHostPort}. Should be in format of HOST:PORT"
    exit 1
  fi
fi

exec java $opts -jar $@