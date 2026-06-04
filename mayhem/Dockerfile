FROM golang:1.21 AS builder

ENV GOTOOLCHAIN=auto

ARG project="levenshtein"
ARG fuzzTargetPath="fuzz"
ARG fuzzTarget="computedist-fuzz"

ADD . ${project}
WORKDIR ${project}
RUN go test -c -fuzz=FuzzComputeDistance \
    -o /${fuzzTarget}.libfuzzer \
    ./${fuzzTargetPath}

FROM ubuntu:20.04
COPY --from=builder /computedist-fuzz.libfuzzer /
