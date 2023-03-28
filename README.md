GPT-Go: OpenAI ChatGPT/GPT-4/GPT-3 SDK Go Client to Interact with the GPT-4/GPT-3 APIs.
========================

<p align="center">
    <br> English | <a href="README-CN.md">中文</a>
</p>

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/csuzhang/gpt-go/main/LICENSE) ![Go](https://github.com/hanyuancheung/gpt-go/workflows/Go/badge.svg)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/hanyuancheung/gpt-go)](https://pkg.go.dev/github.com/hanyuancheung/gpt-go)
[![Go Report Card](https://goreportcard.com/badge/hanyuancheung/gpt-go)](https://goreportcard.com/report/hanyuancheung/gpt-go)
[![codecov](https://codecov.io/gh/hanyuancheung/gpt-go/branch/main/graph/badge.svg)](https://codecov.io/gh/hanyuancheung/gpt-go)

OpenAI Docs API Reference: https://platform.openai.com/docs/api-reference/introduction

> **Note**: Already support GPT-4 API, please use Chat Completions API.

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

## How To Get `API_KEY`

https://platform.openai.com/account/api-keys

## Documentation

Check out the go docs for more detailed documentation on the types and methods provided: https://pkg.go.dev/github.com/hanyuancheung/gpt-go

## Support

- [x] List Engines API
- [x] Get Engine API
- [x] Completion API (this is the main gpt-3 API)
- [x] Streaming support for the Completion API
- [x] Document Search API
- [x] Image generation API
- [x] Overriding default url, user-agent, timeout, and other options

## Contributor

<a href="https://github.com/hanyuancheung/gpt-go/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=hanyuancheung/gpt-go" />
</a>

## Contribute

Please open up an issue on GitHub before you put a lot of efforts on pull request.
The code submitting to PR must be filtered with `gofmt`.

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=hanyuancheung/gpt-go&type=Date)](https://star-history.com/#hanyuancheung/gpt-go&Date)

## License

This package is licensed under MIT license. See LICENSE for details.

## Show your support

Give a ⭐️ if this project helped you!
