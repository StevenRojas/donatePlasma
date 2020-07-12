FROM golang:alpine AS matcher

RUN mkdir -p /code/certs
COPY certs/donatuplasma.org.cer /code/certs/donatuplasma.org.cer
COPY certs/donatuplasma.org.key /code/certs/donatuplasma.org.key
# Working directory for build
WORKDIR /build

# Download dependecies from go.mod
COPY go.mod .
#COPY go.sum .
RUN go mod download

# Build the code to get the binary (matcher)
COPY . .
RUN go build -o matcher

# Working directory for place the binary
WORKDIR /bin
RUN cp /build/matcher .
EXPOSE 8000
ENTRYPOINT ["/bin/matcher"]