#!/usr/bin/env fish

function build_arm
    echo ">> Building Go-FlexGet for ARM…"
    set -lx GOARCH "arm"
    set -lx GOOS "linux"
    set -lx GOARM 5

    go build
    echo ">> Done"
end

function package
    echo ">> Packaging Go-FlexGet…"
    mkdir go-flexget.pkg/

    mv go-flexget go-flexget.pkg/
    cp application.properties logging.properties go-flexget.pkg/
    cp -R public/ go-flexget.pkg/public
    cp -R views/ go-flexget.pkg/views

    mv go-flexget.pkg go-flexget
    zip -r go-flexget.zip go-flexget/
    rm -r go-flexget/
    echo ">> Done"
end

switch $argv
    case "build"
        build_arm
    case "package"
        package
end
