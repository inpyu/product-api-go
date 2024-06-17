FROM alpine:latest as base

# Create group and user
RUN addgroup api && \
  adduser -D -G api api

# Create application directory
RUN mkdir /app

# Install necessary libraries for the binary (if needed)
RUN apk add --no-cache libc6-compat

# Copy the product-api binary
COPY main /app/product-api
RUN chmod +x /app/product-api

# Copy the configuration file
COPY conf.json /conf.json

# Set entrypoint
ENTRYPOINT [ "/app/product-api" ]
