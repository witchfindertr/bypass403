# bypass403

![made with go](https://img.shields.io/badge/made%20with-Go-0040ff.svg) ![maintenance](https://img.shields.io/badge/maintained%3F-yes-0040ff.svg) [![open issues](https://img.shields.io/github/issues-raw/drsigned/bypass403.svg?style=flat&color=0040ff)](https://github.com/drsigned/bypass403/issues?q=is:issue+is:open) [![closed issues](https://img.shields.io/github/issues-closed-raw/drsigned/bypass403.svg?style=flat&color=0040ff)](https://github.com/drsigned/bypass403/issues?q=is:issue+is:closed) [![license](https://img.shields.io/badge/license-MIT-gray.svg?colorB=0040FF)](https://github.com/drsigned/bypass403/blob/master/LICENSE) [![twitter](https://img.shields.io/badge/twitter-@drsigned-0040ff.svg)](https://twitter.com/drsigned)

bypass403 is tool to bypass `403 Forbidden` responses.

## Resources

* [Installation](#installation)
    * [From Binary](#from-binary)
    * [From source](#from-source)
    * [From github](#from-github)
* [Usage](#usage)
* [Contribution](#contribution)

## Installation

#### From Binary

You can download the pre-built binary for your platform from this repository's [releases](https://github.com/drsigned/bypass403/releases/) page, extract, then move it to your `$PATH`and you're ready to go.

#### From Source

bypass403 requires **go1.14+** to install successfully. Run the following command to get the repo

```bash
$ GO111MODULE=on go get -u -v github.com/drsigned/bypass403/cmd/bypass403
```

#### From Github

```bash
$ git clone https://github.com/drsigned/bypass403.git; cd bypass403/cmd/bypass403/; go build; mv bypass403 /usr/local/bin/; bypass403 -h
```


## Usage

To display help message for bypass403 use the `-h` flag:

```
$ bypass403 -h

 _                               _  _    ___ _____
| |__  _   _ _ __   __ _ ___ ___| || |  / _ \___ /
| '_ \| | | | '_ \ / _` / __/ __| || |_| | | ||_ \
| |_) | |_| | |_) | (_| \__ \__ \__   _| |_| |__) |
|_.__/ \__, | .__/ \__,_|___/___/  |_|  \___/____/ v1.2.0
       |___/|_|

USAGE:
  bypass403 [OPTIONS]

OPTIONS:
  -c               concurrency level (default: 20)
  -delay           delay between requests (ms) (default: 100)
  -urls            urls with 403 to bypass (use `-` to read stdin)
```

## Contribution

[Issues](https://github.com/drsigned/bypass403/issues) and [Pull Requests](https://github.com/drsigned/bypass403/pulls) are welcome!