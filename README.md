# Easycron

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

easycron "0 0 * * *" # to evaluate an expression
```

## License

MIT
