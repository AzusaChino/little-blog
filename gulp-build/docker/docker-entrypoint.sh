#!/usr/bin/env sh

JAR_PARAM = ''
if [$REMOTE_DEBUG -eq 'true']; then
  JAR_PARAM = agentlib:jdwp=transport=dt_socket,server=y,suspend=n,address=5005
else
  JAR_PARAM = ''
java -Dspring.profiles.active=docker -jar $JAR_PARAM app.jar