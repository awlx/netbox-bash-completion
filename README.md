# Netbox-Autocompletion
Looks up devices from Netbox and offers autocompletion for SSH

### Usage
```
Usage of ./netbox-autocompletion:
  -netbox string
        Netbox BaseURL (default "https://netbox.local")
  -netbox-api-token string
        Mandatory: Netbox API Token
  -netbox-device string
        Device String to search for
  -tld string
        Default TLD for devices (default "local")
```

### Add this to your .bashrc
```source $GOTPATH/src/github.com/awlx/netbox-autocompletion/bash/completion-go.sh```
