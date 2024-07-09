package encoder

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

const (
	ModeGob = iota
	ModeJson
	ModeYaml
	ModeToml
)

var encoder = ModeJson

func SetEncoderGob() {
	encoder = ModeGob
}

func SetEncoderJson() {
	encoder = ModeJson
}

func SetEncoderYaml() {
	encoder = ModeYaml
}

func SetEncoderToml() {
	encoder = ModeToml
}

func EncodeMessage(v any) ([]byte, error) {
	switch encoder {
	case ModeJson:
		return jsonEncode(v)
	case ModeGob:
		return gobEncode(v)
	case ModeYaml:
		return yamlEncode(v)
	case ModeToml:
		return tomlEncode(v)
	default:
		return nil, fmt.Errorf("unknown encoder mode")
	}
}

func DecodeMessage(data []byte, v any) error {
	switch encoder {
	case ModeJson:
		return jsonDecode(data, v)
	case ModeGob:
		return gobDecode(data, v)
	case ModeYaml:
		return yamlDecode(data, v)
	case ModeToml:
		return tomlDecode(data, v)
	default:
		return fmt.Errorf("unknown encoder mode")
	}
}

func jsonEncode(v any) ([]byte, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func jsonDecode(data []byte, v any) error {
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	return nil
}

func gobEncode(v any) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func gobDecode(data []byte, v any) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(v); err != nil {
		return err
	}
	return nil
}

func yamlEncode(v any) ([]byte, error) {
	b, err := yaml.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func yamlDecode(data []byte, v any) error {
	if err := yaml.Unmarshal(data, &v); err != nil {
		return err
	}
	return nil
}

func tomlEncode(v any) ([]byte, error) {
	var buf bytes.Buffer
	enc := toml.NewEncoder(&buf)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func tomlDecode(data []byte, v any) error {
	if _, err := toml.Decode(string(data), v); err != nil {
		return err
	}
	return nil
}
