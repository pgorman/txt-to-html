#!/bin/sh
set -euf

GOOS=openbsd GOARCH=386 go build
if [ $? -a -e "txt-to-html" ]; then
	scp -q txt-to-html paulgorman.org:bin/
else
	echo "Something went wrong!"
fi
