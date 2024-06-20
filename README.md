# Helium

---

![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/BROngineer/helium/tests.yaml?branch=main&logo=github&label=tests)
![Codecov (with branch)](https://img.shields.io/codecov/c/github/BROngineer/helium/main?logo=codecov)
[![Go Report Card](https://goreportcard.com/badge/github.com/brongineer/helium)](https://goreportcard.com/report/github.com/brongineer/helium)
![GitHub License](https://img.shields.io/github/license/BROngineer/helium)
![GitHub go.mod Go version (branch)](https://img.shields.io/github/go-mod/go-version/BROngineer/helium/main?logo=go&label=Go)

Lightweight library to build command-line application (single-command app yet).

### Features

- Supported flag value types:

  - `string`
  - `int`, `int[8,16,32,64]`
  - `uint`, `uint[8,16,32,64]`
  - `float[32,64]`
  - `time.Duration`
  - `bool`
  - `counter` (which has value of type `int` under the hood)
  - slices of all types in above, except `counter`

- Allows to create custom-typed (generic) flags with user-defined input parser (see [example](./examples/custom/example.go)).
- Allows to override default parser for built-in flag types.

### Future plans

- Add help/usage rendering for flag set
- Add `Command` type to allow multi-command applications
