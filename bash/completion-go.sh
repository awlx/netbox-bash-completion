#/usr/bin/env bash
_netbox_completions()
{
ARGS=`$GOPATH/bin/netbox-bash-completion -netbox <netbox_url> -netbox-api-token <token> -netbox-device ${COMP_WORDS[1]} -tld "" 2> /dev/null`

if [ $? -eq 0 ]; then
COMPREPLY=($(compgen -W "$ARGS" -- ${COMP_WORDS[COMP_CWORD]} ))
fi
}

complete -F  _netbox_completions ssh
