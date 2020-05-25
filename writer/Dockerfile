
# --- buider stage

FROM golang:1.14-buster AS builder

# Set working directory
WORKDIR /src

# Prepare Golang environment
ENV GO111MODULE on

# Copy go-modules definition for the application
COPY go.mod go.mod

# Download required go-modules
RUN go mod download

# Copy application code
COPY . .

# Compile application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bin/app


# --- final stage

#FROM scratch
FROM alpine

# Install additionals
RUN apk add --no-cache bash

# Copy application executable
WORKDIR /bin/
COPY --from=builder /bin/app .

## Metrics
EXPOSE 9090
## KubeProbes
EXPOSE 9091

# Run application
ENTRYPOINT "/bin/app"
