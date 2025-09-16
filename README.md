# semaver

`semaver` is a Go package for semantic versioning (SemVer) comparison. It allows you to easily parse and compare semantic versions in the format of `MAJOR.MINOR.PATCH[-PRERELEASE][+BUILD]`. The package supports all the core features of [Semantic Versioning 2.0](https://semver.org/), including version comparisons with pre-release versions, build metadata, and comparison operators.

## Features

- Parse and validate semantic version strings.
- Compare versions using operators like `=`, `<`, `>`, `<=`, `>=`.
- Handle versions with pre-release labels (e.g., `alpha`, `beta`) and build metadata (e.g., `+build123`).
- Fully compliant with [SemVer 2.0](https://semver.org/).

## Installation

To use `semaver`, simply import the package into your Go project:

```bash
go get github.com/nurettintopal/semaver
````

Then import it in your Go files:

```go
import "github.com/nurettintopal/semaver"
```

## Usage

Here is an example of how to use `semaver` to compare versions:

```go
package main

import (
	"fmt"
	"log"

	"github.com/nurettintopal/semaver"
)

func main() {
	result, err := semaver.CompareVersions("1.0.0-alpha.1", "<", "1.0.0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0-alpha.1", ">", "1.0.0-alpha")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0-alpha.1", "=", "1.0.0-alpha.1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0-beta.2", "<", "1.0.0-beta.3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0+build123", ">", "1.0.0+build456")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("false => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0", "<", "1.2.2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.2.1", ">", "1.2.0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("2.0.3", "=", "2.0.3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("2.0.2", ">=", "2.0.3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("false => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0-alpha.1", "<", "1.0.0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0-alpha.1", ">", "1.0.0-alpha")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0-alpha.1", "=", "1.0.0-alpha.1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0-beta.2", "<", "1.0.0-beta.3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.1.0-beta.2", "<", "1.0.0-beta.3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("false => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0-alpha", ">=", "1.0.0-alpha.1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("false => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0+build123", ">", "1.0.0+build456")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("false => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0-alpha+build789", "<", "1.0.0-alpha+build000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("false => %v\n", result)

	result, err = semaver.CompareVersions("1.2.0-alpha", "<", "1.2.0-beta")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.2.0-alpha.1", ">", "1.2.0-alpha")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0+build0001", "=", "1.0.0+build0001")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0", "<", "1.2.2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.2.1", ">", "1.2.0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("2.0.3", "=", "2.0.3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("2.0.2", ">=", "2.0.3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("false => %v\n", result)

	result, err = semaver.CompareVersions("1.4.3", ">=", "1.0.1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0", "<=", "1.0.0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("true => %v\n", result)

	result, err = semaver.CompareVersions("1.0.0", ">", "1.0.0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("false => %v\n", result)
}
```

### Available Operators

* `=`: Equal to
* `<`: Less than
* `>`: Greater than
* `<=`: Less than or equal to
* `>=`: Greater than or equal to

### Functionality

The main function in the `semaver` package is `CompareVersions`:

```go
func CompareVersions(v1, operator, v2 string) (bool, error)
```

* `v1`: First version string.
* `operator`: Comparison operator (`=`, `<`, `>`, `<=`, `>=`).
* `v2`: Second version string.

It returns a `bool` indicating whether the comparison is true and an `error` if the version strings are invalid.

### Version Parsing

You can also parse a version string into a `Version` struct using:

```go
func ParseVersion(v string) (Version, error)
```

This returns a `Version` struct that contains:

* `major`: The major version number.
* `minor`: The minor version number.
* `patch`: The patch version number.
* `prerelease`: The optional pre-release label (e.g., `alpha`, `beta`).
* `build`: The optional build metadata (e.g., `+build123`).

## Examples

Here are some more examples of how `semaver` can be used:

```go
result, err := semaver.CompareVersions("1.0.0-alpha", "<", "1.0.0-beta")
fmt.Println(result) // true (alpha < beta)

result, err = semaver.CompareVersions("2.0.0", ">", "1.9.9")
fmt.Println(result) // true (2.0.0 > 1.9.9)

result, err = semaver.CompareVersions("1.0.0+build123", "=", "1.0.0+build123")
fmt.Println(result) // true (build metadata is ignored)
```

## Example Comparisons

| Version A            | Version B            | Result |
| -------------------- | -------------------- | ------ |
| 1.0.0-alpha.1        | 1.0.0                | A < B  |
| 1.0.0-alpha.1        | 1.0.0-alpha          | A > B  |
| 1.0.0-alpha.1        | 1.0.0-alpha.1        | A == B |
| 1.0.0-beta.2         | 1.0.0-beta.3         | A < B  |
| 1.0.0+build123       | 1.0.0+build456       | A < B  |
| 1.0.0                | 1.2.2                | A < B  |
| 1.2.1                | 1.2.0                | A > B  |
| 2.0.3                | 2.0.3                | A == B |
| 2.0.2                | 2.0.3                | A < B  |
| 1.0.0-alpha.1        | 1.0.0                | A < B  |
| 1.0.0-alpha.1        | 1.0.0-alpha          | A > B  |
| 1.0.0-alpha.1        | 1.0.0-alpha.1        | A == B |
| 1.0.0-beta.2         | 1.0.0-beta.3         | A < B  |
| 1.1.0-beta.2         | 1.0.0-beta.3         | A > B  |
| 1.0.0-alpha          | 1.0.0-alpha.1        | A < B  |
| 1.0.0+build123       | 1.0.0+build456       | A < B  |
| 1.0.0-alpha+build789 | 1.0.0-alpha+build000 | A > B  |
| 1.2.0-alpha          | 1.2.0-beta           | A < B  |
| 1.2.0-alpha.1        | 1.2.0-alpha          | A > B  |
| 1.0.0+build0001      | 1.0.0+build0001      | A == B |
| 1.0.0                | 1.2.2                | A < B  |
| 1.2.1                | 1.2.0                | A > B  |
| 2.0.3                | 2.0.3                | A == B |
| 2.0.2                | 2.0.3                | A < B  |
| 1.4.3                | 1.0.1                | A > B  |
| 1.0.0                | 1.0.0                | A == B |
| 1.0.0                | 1.0.0                | A == B |
| 1.0.0                | 1.0.0                | A == B |

## Contributing

We welcome contributions to `semaver`! To get started:

1. Fork the repository.
2. Create a new branch.
3. Make your changes and add tests (if applicable).
4. Submit a pull request.

Please make sure your code adheres to the [Go Code Style](https://golang.org/doc/effective_go).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

* [Semantic Versioning 2.0](https://semver.org/) - The specification that powers this package.
