#/usr/bin/env bash
_netbox_completions()
{
ARGS=`$GOPATH/bin/netbox-autocompletion -netbox-api-token <netbox-api-token> -netbox-device ${COMP_WORDS[1]}`
for i in $ARGS; do
    COMPREPLY+=("$i")
done;
}

complete -F  _netbox_completions ssh
