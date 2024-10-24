# nmap playground

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

```sh
go run . -sV --script default,vuln,brute localhost
```
