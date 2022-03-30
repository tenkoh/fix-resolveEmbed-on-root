FROM golang:1.18

WORKDIR /build
RUN git clone https://go.googlesource.com/go

# build after fixing pkg.go
COPY ./src/pkg.go ./go/src/cmd/go/internal/load
WORKDIR /build/go/src
RUN ./make.bash

WORKDIR /
COPY ./ ./

# CMD ["/build/go/bin/go", "list", "-f", "'pattern: {{.EmbedPatterns}}, files: {{.EmbedFiles}}'"]
CMD ["bash"]
