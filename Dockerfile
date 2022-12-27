FROM golang:latest
LABEL maintainer="teguh-satriya"
ENV GATEWAY_PORT="8081" \
    GRPC_PORT="8080" \
    HOME="/opt/app-root"
RUN set -xe \
 && apt-get update -qy \
 && apt-get install -qy software-properties-common gettext-base git \
 && mkdir -p ${HOME}
WORKDIR ${HOME}
COPY . .
RUN make install
RUN set -xe \
 && chown -R 1001 . \
 && chgrp -R 0 . \
 && chmod -R g=u .
USER 1001
EXPOSE ${GATEWAY_PORT}
EXPOSE ${GRPC_PORT}
CMD ["sh", "-c",  "privy-cake"]