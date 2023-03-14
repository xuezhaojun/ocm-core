FROM golang:1.19 AS builder
ARG OS=linux
ARG ARCH=amd64
WORKDIR /go/src/open-cluster-management.io/ocm-core
COPY . .
ENV GO_PACKAGE open-cluster-management.io/ocm-core

RUN GOOS=${OS} \
    GOARCH=${ARCH} \
    make build --warn-undefined-variables

FROM registry.access.redhat.com/ubi8/ubi-minimal:latest
ENV USER_UID=10001

COPY --from=builder /go/src/open-cluster-management.io/ocm-core/registration /
COPY --from=builder /go/src/open-cluster-management.io/ocm-core/work /

USER ${USER_UID}