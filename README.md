GPT-Go: OpenAI GPT-3 SDK Client Enables Go Programs to Interact with the GPT-3 APIs.
========================

<p align="center">
    <br> English | <a href="README-CN.md">中文</a>
</p>

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/csuzhang/gpt-go/main/LICENSE) ![Go](https://github.com/hanyuancheung/gpt-go/workflows/Go/badge.svg)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/hanyuancheung/gpt-go)](https://pkg.go.dev/github.com/hanyuancheung/gpt-go)

## Quick Start

```shell
# clone the project
git clone https://github.com/hanyuancheung/gpt-go.git

# go to the project directory
cd gpt-go

# set API_KEY as environment variable
export API_KEY={YOUR_API_KEY} chatgpt

# go build example binary
make chatgpt-example

# run example
./chatgpt
```

## Snapshot

![](img/chatgpt.gif)

## Documentation

Check out the go docs for more detailed documentation on the types and methods provided: https://pkg.go.dev/github.com/hanyuancheung/gpt-go

## Support

- [x] List Engines API
- [x] Get Engine API
- [x] Completion API (this is the main gpt-3 API)
- [x] Streaming support for the Completion API
- [x] Document Search API
- [x] Overriding default url, user-agent, timeout, and other options

## Contribute

Please open up an issue on GitHub before you put a lot of efforts on pull request.
The code submitting to PR must be filtered with `gofmt`.

## License

This package is licensed under MIT license. See LICENSE for details.
