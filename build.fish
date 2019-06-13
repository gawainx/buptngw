#!/usr/local/bin/fish
function build_linux
    set CGO_ENABLE 0
    set GOOS linux
    set GOARCH amd64
    echo "Building for linux"
    go build -o linux main.go 
end

function build_darwin
    echo "Building for Darwin"
    go build -o main_darwin main.go 
end

switch $argv
    case linux
        build_linux
    case '*'
        build_darwin
end