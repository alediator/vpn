# VPN connection helper

This is a command build just for learn a bit of golang that depends on `openconnect` and any post script you want to execute it before. 

## Requirements

* Openconnect

## How to run it

1. Download the binary from
2. Copy it to the place you want to use it (f.i. `/usr/local/bin`)
3. Execute it with: `vpn connect -h`

And follow the instructions.

If you want to save part of the configuration, just get the `.env.dist` file, edit it and save it at `/home/$USER/.vpn`
