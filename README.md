# Encoder Package

A Go package for encoding and decoding messages in various formats: JSON, GOB, YAML, and TOML.

## Installation

To install the package, use `go get`:

```sh
go get github.com/dyammarcano/encoder
```

## Usage

To use the package, import it in your code:

```go
import "github.com/dyammarcano/encoder"

type ExampleStruct struct {
    Name string
    Age  int
}

func main() {
    example := ExampleStruct{Name: "Alice", Age: 30}

    encoder.SetEncoderJson()
    encodedData, err := encoder.EncodeMessage(example)
    if err != nil {
        fmt.Println("Error encoding message:", err)
        return
    }

    var decodedExample ExampleStruct
    err = encoder.DecodeMessage(encodedData, &decodedExample)
    if err != nil {
        fmt.Println("Error decoding message:", err)
        return
    }

    fmt.Printf("Decoded Message: %+v\n", decodedExample)
}
```

# Contributing
Feel free to open issues or submit pull requests with improvements.

# License
This project is licensed under the MIT License. See the LICENSE file for details.

# Acknowledgements
This project uses the following libraries:

- [gopkg.in/yaml.v3](gopkg.in/yaml.v3)
- [github.com/BurntSushi/toml](github.com/BurntSushi/toml)
