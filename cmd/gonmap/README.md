# nmap playground

```sh
OTEL_SERVICE_NAME="gonmap" OTEL_RESOURCE_ATTRIBUTES="deployment.environment=hack-1743" OTEL_EXPORTER_OTLP_ENDPOINT="http://localhost:4317" go run . -sV --script default,vuln,brute localhost
```

```sh
sudo apt install nmap
sudo nmap --script-updatedb
```

<https://nmap.org/book/man-briefoptions.html>

<https://nmap.org/book/vscan-examples.html>

```sh
nmap -A -T4 localhost 
```

<https://nmap.org/book/man-version-detection.html>

```sh
nmap -sV localhost
```

<https://nmap.org/book/nse-usage.html#nse-categories>

```sh
nmap -sV --script vuln localhost
```

```sh
nmap -sV --script default,vuln,brute localhost
```

```sh
ls /usr/share/nmap/scripts
```

