#!/bin/bash

# NOTE!! Never ever use this if youn want to check the error codes by yourself
#    set -e
# as we want to do here:

if ! command -v lichen &> /dev/null
then
	go install github.com/uw-labs/lichen@latest
fi

case "$2" in
	c)
		lichen \
			-c ./scripts/lichen-cfg.yaml \
			--template="{{range .Modules}}{{range .Module.Licenses}}{{.Name | printf \"%s\n\"}}{{end}}{{end}}" \
			"$1" | sort | uniq -c | sort -nr
		;;
	v)
		lichen -c ./scripts/lichen-cfg.yaml "$1"
		;;
	*)
		lichen -c ./scripts/lichen-cfg.yaml --template="" $1
		;;
esac

result=$?

rm "$1"

if [ $result -gt 0 ]
then
	echo "Licensing error"
	exit 1
fi
echo "Licensing OK"

