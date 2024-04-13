# go-tutorials
Tutorials sandbox in golang

## Resources

https://go.dev/doc/

## Topics covered

- [ ] [Tutorial: Getting started with multi-module workspaces](https://go.dev/doc/tutorial/workspaces.html)
- [x] [Tutorial: Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin.html)
- [x] [Tutorial: Getting started with fuzzing](https://go.dev/doc/tutorial/fuzz.html)
- [x] [Writing Web Applications](https://go.dev/doc/articles/wiki/)
- [ ] [Standard library](https://pkg.go.dev/std) understanding
  - [ ] archive/
  - [ ] bufio/
  - [ ] builtin/
  - [ ] bytes/
  - [ ] cmp/
  - [ ] compress/
  - [ ] container/
  - [ ] context/
  - [ ] crypto/
  - [ ] database/
  - [ ] debug/
  - [ ] encoding/
  - [ ] errors/
  - [ ] expvar/
  - [ ] flag/
  - [ ] fmt/
  - [ ] go/
  - [ ] hash/
  - [ ] html/
  - [ ] image/
  - [ ] index/
  - [ ] io/
  - [ ] log/
  - [ ] maps/
  - [ ] math/
  - [ ] mime/
  - [ ] net/
  - [ ] os/
  - [ ] path/
  - [ ] plugin/
  - [ ] reflect/
  - [ ] regexp/
  - [ ] runtime/
  - [ ] slices/
  - [ ] sort/
  - [ ] strconv/
  - [ ] strings/
  - [ ] sync/
  - [ ] syscall/
  - [ ] testing/
  - [ ] text/
  - [ ] time/
  - [ ] unicode/
  - [ ] unsafe/

## Programing themes to understand std library

| Theme                                              | Program Description                                                  | Libraries Used                          |
|----------------------------------------------------|----------------------------------------------------------------------|-----------------------------------------|
| File Operation System                              | Create, read, write, and delete files and directories.               | `os`, `io`, `bufio`, `path`, `errors`   |
| Network Communication                              | Create an HTTP server and client to perform basic communication.     | `net`, `http`, `html`, `log`, `context` |
| Data Compression and Decompression                 | Compress and decompress files in gzip format.                        | `compress`, `archive`, `bytes`, `io`    |
| Data Handling                                      | Encode and decode data in JSON and XML formats.                      | `encoding`, `strings`, `bytes`          |
| Data Validation and Regular Expressions            | Perform pattern matching and validation on input strings.            | `regexp`, `strings`, `strconv`          |
| Performance and Debugging                          | Collect application metrics and display debug information.           | `runtime`, `debug`, `expvar`, `log`     |
| Security and Encryption                            | Encrypt and decrypt data, and generate hashes.                       | `crypto`, `hash`                        |
| Synchronization and Concurrency                    | Process data using multiple goroutines and manage synchronization.   | `sync`, `context`, `runtime`, `io`      |
| Command Line Tools                                 | Parse command line arguments and provide a simple interface.         | `flag`, `os`, `fmt`, `log`              |
| Text Processing and Time Management                | Log time to a log file and manipulate text data in specific formats. | `time`, `text`, `strings`, `strconv`    |
| File and Database Integration                      | Manage data combining database operations with file system.          | `database/sql`, `os`, `bufio`, `fmt`    |
| Image Processing                                   | Load, process, and save image files.                                 | `image`, `image/color`, `io`            |
| Internationalization and Character Code Processing | Handle Unicode strings and support multiple languages.               | `unicode`, `strings`, `strconv`         |


