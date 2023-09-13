FROM alpine:3.18
ENV TZ Asia/Shanghai
RUN apk add alpine-conf tzdata dumb-init && \
    /sbin/setup-timezone -z Asia/Shanghai && \
    apk del alpine-conf

ENV WORKDIR /app
VOLUME $WORKDIR/data
ADD config.example.toml $WORKDIR/data/
ADD Moe $WORKDIR
WORKDIR $WORKDIR

ENTRYPOINT ["dumb-init", "--"]
CMD ["./Moe"]


