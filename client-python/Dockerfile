# pull official base image
FROM python:3.8-slim-buster

#set environment variables
# I don't want to generate pcy files
ENV PYTHONDONTWRITEBYTECODE 1
# ignore buffering
ENV PYTHONUNBUFFERED 1
# set encoding
ENV PPYTHONENCODING utf-8

#set work directory
WORKDIR /root/workdir/src

COPY ../src /root/workdir/src

#
RUN apt-get update

RUN pip install --upgrade pip
RUN pip install -r ./requirements.txt