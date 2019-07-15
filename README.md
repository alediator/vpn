# VPN connection helper

This is a command build just for learn a bit of golang that depends on `openconnect` and any post script you want to execute it before. 

## About

Bear in mind that this is a first language project just to play a bit with [golang](https://golang.org/) and [cobra](https://github.com/spf13/cobra) for the cli. Generating more than 4MB binary for something that you would do it with a simple bash script is too much, by the way it works! 

On the other hand, thanks to [CodelyTV](https://pro.codely.tv/library/introduccion-a-go-tu-primera-app/89042/path/step/57224893/) and [friendsofgo](https://blog.friendsofgo.tech/) for introducing me to golang.

## Requirements

* Openconnect. This is a simple wrapper to use openconnect, so you might have installed it as well as any additional dependency to connect your VPN
* In addition, you could install [vpn-slice](https://github.com/dlenski/vpn-slice) and use it as part of your setup with the post script information `VPN_POST_SCRIPT`

## How to run it

1. Download the [release binary](https://github.com/alediator/vpn/releases/download/v1.0.0/vpn)
2. Make it executable: `chmod +x vpn`
3. Copy it to the place you want to use it (f.i. `/usr/local/bin`)
4. Execute it with: `vpn connect -h`

And follow the instructions:

```
> vpn connect -h
Helper to connect to a VPN through openconnect. 
Note that this is able to read `.env` files (by default in `/home/$USER/.vpn`)

Usage:
  vpn-cli connect [flags]

Flags:
  -a, --flags string      additional flags for openconnect - VPN_ADDITIONAL_FLAGS at .env
  -h, --help              help for connect
  -v, --host string       host to use for the vpn - VPN_HOST at .env
  -p, --password string   static password - VPN_STATIC_PASSWORD at .env
  -s, --script string     script to execute - VPN_POST_SCRIPT at .env
  -t, --token string      dynamic token key
  -u, --user string       user to use - VPN_USER at .env
```

If you want to save part of the configuration, just get the [.env.dist](https://github.com/alediator/vpn/blob/master/.env.dist) file, edit it and save it at `/home/$USER/.vpn`