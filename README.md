# microcart

[![tests](https://github.com/brokeyourbike/microcart/actions/workflows/tests.yml/badge.svg)](https://github.com/brokeyourbike/microcart/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/brokeyourbike/microcart)](https://goreportcard.com/report/github.com/brokeyourbike/microcart)
[![Maintainability](https://api.codeclimate.com/v1/badges/fbaad4c7799f157bde69/maintainability)](https://codeclimate.com/github/brokeyourbike/microcart/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/fbaad4c7799f157bde69/test_coverage)](https://codeclimate.com/github/brokeyourbike/microcart/test_coverage)

Compact e-commerce solution, minimal resources, ultra-fast performance.

## Local development

### Requirements

1. [Visual Studio Code](https://code.visualstudio.com/)
1. [Dev Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)
1. [Docker](https://www.docker.com/products/docker-desktop/)

### Container Setup

Run the command from the command palette `Dev Containers: Reopen in Container` to open the project in a container.

![Reopen in Container](https://i.imgur.com/eKiWJn3.png)

### Database Migration

Run `go run app/tooling/cli/main.go migrate` to migrate models to the database.

## Authors
- [Ivan Stasiuk](https://github.com/brokeyourbike) | [Twitter](https://twitter.com/brokeyourbike) | [LinkedIn](https://www.linkedin.com/in/brokeyourbike) | [stasi.uk](https://stasi.uk)

## License
[BSD-3-Clause License](https://github.com/glocurrency/microcart/blob/main/LICENSE)
