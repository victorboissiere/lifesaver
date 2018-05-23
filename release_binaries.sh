#!/usr/bin/env sh

export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

echo "Type github token"
read token
export GITHUB_TOKEN="$token"
echo "Type the version you want to release: [ENTER]"
read version
echo "Type the description of your release: [ENTER]"
read description

git tag -a "$version" -m "Release version $version" && git push --tags

mkdir -p bin

for plateform in "linux"
do
    GOOS=linux go build -o "bin/$plateform" -ldflags "-s -w"
done

# go get github.com/aktau/github-release
github-release release \
    --user victorboissiere \
    --repo lifesaver \
    --tag "$version" \
    --description "$description" \

github-release upload \
    --user victorboissiere \
    --repo lifesaver \
    --tag "$version" \
    --name "linux_lifesaver" \
    --file bin/linux

rm -r bin
