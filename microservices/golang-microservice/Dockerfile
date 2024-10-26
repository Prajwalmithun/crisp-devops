# Start with the base image
FROM golang:1.23.2-alpine3.20

# Set a working directory
WORKDIR /app-dir

# Create a non-root user and group
RUN adduser -D nonrootuser

# Give non-root user ownership of the app directory
RUN chown -R nonrootuser:nonrootuser /app-dir

# Copy go.mod file (contains infor about go modules ie., go dependancies) 
COPY go.mod .

# Downloading the go modules (dependancies)
RUN go mod download

# Copy the source code 
COPY *.go .

# Build the go app 
RUN go build -o /golang-ms

# Expose the port
EXPOSE 8888

# Run
CMD ["/golang-ms"]


