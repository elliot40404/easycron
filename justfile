set windows-shell := ["pwsh.exe", "-NoLogo", "-Command"]

default: build

build_cmd := if os() == "windows" { "go build -o ./bin/easycron.exe ./cmd/easycron/" } else { "go build -o ./bin/easycron ./cmd/easycron/" }

build: clean lint
    {{build_cmd}}

run iter='' cron='':
    go run ./cmd/easycron/ {{iter}} "{{cron}}"

exec iter='' cron='':
    ./bin/easycron {{iter}} "{{cron}}"


install:
    go install ./cmd/easycron/

build-run: build exec

rmcmd := if os() == "windows" { "mkdir ./bin -Force; Remove-Item -Recurse -Force ./bin" } else { "rm -rf ./bin" }

clean:
    {{rmcmd}}

lint:
    golangci-lint run

lint-fix:
    golangci-lint run --fix

vendor:
    go mod tidy
    go mod vendor
    go mod tidy

release:
    goreleaser release --snapshot --clean
