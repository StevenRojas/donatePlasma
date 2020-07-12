FROM golang:alpine AS register

RUN mkdir -p /code/certs
COPY certs/donatuplasma.org.cer /code/certs/donatuplasma.org.cer
COPY certs/donatuplasma.org.key /code/certs/donatuplasma.org.key
# Working directory for build
WORKDIR /build

# Download dependecies from go.mod
COPY go.mod .
#COPY go.sum .
RUN go mod download
RUN mkdir -p /frontend/static/images

# Build the code to get the binary (register)
COPY . .
RUN go build -o register

# Working directory for place the binary
WORKDIR /bin
RUN cp /build/register .
EXPOSE 8000
ENTRYPOINT ["/bin/register"]