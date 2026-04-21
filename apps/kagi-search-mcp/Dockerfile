# syntax=docker/dockerfile:1

FROM docker.io/library/python:3.12-alpine3.22
ARG VERSION

ENV \
    PIP_BREAK_SYSTEM_PACKAGES=1 \
    PIP_DISABLE_PIP_VERSION_CHECK=1 \
    PIP_NO_CACHE_DIR=1 \
    PIP_ROOT_USER_ACTION=ignore \
    PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1 \
    UV_NO_CACHE=true \
    UV_SYSTEM_PYTHON=true

USER root
WORKDIR /app

RUN \
    apk add --no-cache \
        catatonit \
    && pip install --no-cache-dir uv \
    && uv pip install kagimcp=="${VERSION}" \
    && pip uninstall --yes uv \
    && rm -rf /root/.cache /tmp/*

USER nobody:nogroup

ENTRYPOINT ["/usr/bin/catatonit", "--", "kagimcp"]
