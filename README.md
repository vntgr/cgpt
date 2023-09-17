# kily

Mighty ChatGPT CLI Client. A CLI tool written in Golang that uses OpenAI's ChatGPT API to interact just from your terminal.

## Go version

`1.21.0`

## Setup

```console
$ go build
```

Then initialize configuration with `kily init`:

```console
$ kily init

██   ██ ██ ██      ██    ██
██  ██  ██ ██       ██  ██
█████   ██ ██        ████
██  ██  ██ ██         ██
██   ██ ██ ███████    ██

 INFO  initializing kily configuration

Enter your OpenAI API key: ***************************************************

 SUCCESS  kily configuration initialized

```

## Usage

```console
$ kily messenger Tell me one famous quotation from Rumi

"Let yourself be silently drawn by the strange pull of what you really love. It will not lead you as
tray." - Rumi
```

```console
$ kily messenger When was Bruce Lee born?

Bruce Lee was born on November 27, 1940.
```
