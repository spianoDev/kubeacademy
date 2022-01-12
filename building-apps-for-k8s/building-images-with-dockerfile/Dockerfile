# FROM sets base image
FROM golang
LABEL maintainer="spianodev@gmail.com"
# WORKDIR creates a directory and CD into it
WORKDIR /workspace
COPY . .
RUN go install
# ENTRYPOINT launches a process and starts that process
ENTRYPOINT ["hello"]
# same as entrypoint or provides arguments for the entrypoint
CMD ["world"]