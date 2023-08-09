FROM python:3.6.8-alpine

RUN apk add net-snmp net-snmp-dev musl-dev gcc
RUN pip install easysnmp

WORKDIR /snmpworkdir
COPY ./ /snmpworkdir