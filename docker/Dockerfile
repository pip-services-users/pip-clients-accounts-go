FROM debian:stretch

# Set working directory
WORKDIR /app

# Copy entire project
COPY ./dist/run .

# Set default environment variables
ENV COUCHBASE_SERVICE_URI ""
ENV COUCHBASE_SERVICE_HOST couchbase
ENV COUCHBASE_SERVICE_PORT 8091
ENV COUCHBASE_DB app

EXPOSE 8080

ENTRYPOINT [ "/app/run" ]