# Go Error Handling Package (Grr)

Grr is an experimental error handling package for Go that utilizes generics, similar to error handling in Rust. This package provides a convenient and concise way to handle errors in your Go code.

## Installation

To install the Grr package, simply run the following command:

```
go get github.com/levinion/grr
```

Make sure to replace `levinion` with your actual GitHub username.

## Usage

To use the Grr package, import it into your Go code as follows:

```go
import "github.com/levinion/grr"
```

### Example

Here's a simple example to demonstrate the usage of the Grr package:

```go
package main

import (
	"fmt"
	"github.com/levinion/grr"
	"os"
)

func main() {
	Readfile("hello.md").Expect(func(v []byte) {
		fmt.Println("print result:", string(v))
	}).Else(func(err error) {
		fmt.Println("print err:", err)
	})
}

func Readfile(file string) *grr.Result[[]byte] {
	return grr.Try[[]byte](func(h *grr.Handler[[]byte]) {
		file, err := os.Open(file)
		h.Err(err)
		buffer := make([]byte, 4096)
		n, err := file.Read(buffer)
		h.Err(err)
		content := buffer[:n]
		h.OK(content)
	})
}
```

In this example, we use the `Readfile` function to read the contents of a file. The `Readfile` function returns a `Result` object, which allows us to handle both successful outcomes and errors.

The `Expect` method is called when the file is successfully read and contains the file content. It takes a callback function that receives the file content as a parameter.

The `Else` method is called when an error occurs during the file reading process. It takes a callback function that receives the error as a parameter.

By using the Grr package, we can handle errors and successful outcomes in a concise and expressive manner, reducing boilerplate code and improving code readability.

## Contributions

Contributions are always welcome! If you have any suggestions, bug reports, or feature requests, please open an issue on the GitHub repository.

## License

The Grr package is licensed under the [MIT License](https://opensource.org/licenses/MIT). See the LICENSE file for more details.

## Credits

The Grr package was created by [YourName](https://github.com/levinion).
