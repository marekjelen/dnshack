# DnsHack

DnsHack hijacks some of your `DNS` resolutions so that your machine
does not have to rely on external network.

## Why?

Services like `nip.io` or `xip.io` can be sometimes blocked by your DNS
servers, or you may find yourself without connectivity.

## Other solutions

You can use `dnsmasq` or any other of the bazillion possible tools.

## Does it work everywhere?

Possibly, but the instructions are only for macOS.

## Future?

No idea, solved my problem today.

## How to?

Forward the domain names to the local server

```sh
$ sudo cat <<EOF > /etc/resolver/xip.io
port 5300
nameserver 127.0.0.1
EOF

$ sudo cp /etc/resolver/xip.io /etc/resolver/nip.io
$ sudo cp /etc/resolver/xip.io /etc/resolver/test
```

build the server

```sh
$ make
```

and start the server automatically (replace the quoted `[items]` with your values)

```sh
$ sudo cat <<EOF > /Library/LaunchDaemons/dns.plist
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
    <dict>
        <key>Label</key>
        <string>dnshack</string>
        <key>Program</key>
        <string>[path to the binary]</string>
        <key>RunAtLoad</key>
        <true/>
        <key>KeepAlive</key>
        <true/>
        <key>UserName</key>
        <string>[username]</string>
        <key>GroupName</key>
        <string>staff</string>
        <key>InitGroups</key>
        <true/>
    </dict>
</plist>
EOF

$ sudo launchctl load -w /Library/LaunchDaemons/dns.plist
```

## License

`dnshack` is released under the terms of the MIT License.