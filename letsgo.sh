#!/bin/sh

function install
{
    if hash go 2>/dev/null; then
    echo "Go already installed. Nice!"
    else
        if hash brew 2>/dev/null; then
            install_go
        else
            echo "Please install brew"
            exit
        fi
    fi
}

function install_go
{
    brew update
    brew install go
}

gopath="$HOME/go"
function make_gopath
{
    echo "Where would you like your $GOPATH? (Just hit enter for $HOME/go)"
    read gopath
    if [ -z "${gopath// }" ]; then
        gopath="$HOME/go"
    fi
    mkdir -p ${gopath}/{bin,src,pkg}

    echo "\nPlease add these to your bash profile for convenience (trust me):\n"
    echo "export GOPATH=$gopath"
    echo "export PATH=\$PATH:\$GOPATH/bin\n\n"
}

function make_userrepo
{
    echo "What is your github handle?"
    read username
    mkdir -p "$gopath/src/github.com/$username"
    echo "$gopath/src/github.com/$username is all set\n"
}

function install_tools
{
    echo "Installing go tools..."
    go get -u github.com/nsf/gocode
    go get -u github.com/rogpeppe/godef
    go get -u github.com/golang/lint/golint
    go get -u github.com/lukehoban/go-outline
    go get -u sourcegraph.com/sqs/goreturns
    go get -u golang.org/x/tools/cmd/gorename
    go get -u github.com/tpng/gopkgs
    go get -u github.com/newhook/go-symbols
    go get -u golang.org/x/tools/cmd/guru
}

install
make_gopath
make_userrepo
install_tools

echo "\n\nYou, my friend, are ready to GO!\n\n\n"
