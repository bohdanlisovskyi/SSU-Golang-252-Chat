# Base this docker container off the official golang docker image.
# Docker containers inherit everything from their base.
FROM golang:1.8.3
# Create a directory inside the container to store all our application and then make it the working directory.
RUN mkdir -p /go/src/github.com/8tomat8/SSU-Golang-252-Chat
WORKDIR /go/src/github.com/8tomat8/SSU-Golang-252-Chat
# Copy the example-app directory (where the Dockerfile lives) into the container.
COPY . /go/src/github.com/8tomat8/SSU-Golang-252-Chat
# Download and install any required third party dependencies into the container.
RUN go get ./...
# Set the PORT environment variable inside the container
ENV PORT 5000
# Expose port 8080 to the host so we can access our application
EXPOSE 5000
# Now tell Docker what command to run when the container starts
#CMD ["go-wrapper","run"]
CMD go run /go/src/github.com/8tomat8/SSU-Golang-252-Chat/server/main.go
