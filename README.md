# Easycron

[![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/elliot40404/easycron/release.yml)](https://github.com/elliot40404/easycron/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/elliot40404/easycron)](https://goreportcard.com/report/github.com/elliot40404/easycron)
[![Go Reference](https://pkg.go.dev/badge/github.com/elliot40404/easycron.svg)](https://pkg.go.dev/github.com/elliot40404/easycron)

```sh
 ______
|  ____|
| |__     __ _  ___  _   _   ___  _ __   ___   _ __  
|  __|   / _` |/ __|| | | | / __|| '__| / _ \ | '_ \ 
| |____ | (_| |\__ \| |_| || (__ | |   | (_) || | | |
|______| \__,_||___/ \__, | \___||_|    \___/ |_| |_|
                      __/ |
                     |___/
                            - elliot40404
```

![demo](./images/demo.gif)

Easycron is a simple cross platform cli app that helps you with cron

## Why

I do a lot of system admin work and that includes setting up and managing a lot of cron jobs. Usually I use [crontab.guru](https://crontab.guru) to validate the my expressions but I wanted to do this without leaving the terminal, this cli lets me do that.

## Installation

```bash
go install github.com/elliot40404/easycron/cmd/easycron@latest
```

## Usage 

```bash
easycron # to start in interactive mode

easycron <options> <expression>
```

## Features

- [x] Interactive mode
- [x] Human readable cron expression
- [x] Next 3 iterations
- [x] Non interactive mode
- [x] Configurable cli options
- [ ] Schedule cron jobs directly from easycron
- [ ] Cron job manager

## Build From Source with alternative engines

```bash
git clone https://github.com/elliot40404/easycron.git
cd easycron
go build -o easycron cmd/easycron/ # new charm engine
go build -tags tview -o easycron cmd/easycron/ # old tview engine
```

## License

MIT

## Support My Work

<a href="https://ko-fi.com/elliot40404">
<img src="https://storage.ko-fi.com/cdn/brandasset/v2/support_me_on_kofi_red.png" alt="Support Me on Ko-fi" width="200">
</a>
