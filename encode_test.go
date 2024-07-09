package encoder

import (
	"os"
	"testing"
)

type obj struct {
	Message string `json:"message" yaml:"message" toml:"message" gob:"message"`
}

func TestEncoderGob(t *testing.T) {
	t.Log("TestSetEncoderGob")

	SetEncoderGob()

	if encoder != ModeGob {
		t.Errorf("expected ModeGob, got %d", encoder)
	}

	n := obj{}
	n.Message = "Hello, Gob!"

	data, err := EncodeMessage(n)
	if err != nil {
		t.Error(err)
	}

	if len(data) == 0 {
		t.Error("expected data, got empty")
	}

	n2 := obj{}
	if err = DecodeMessage(data, &n2); err != nil {
		t.Error(err)
	}

	t.Log("TestSetEncoderGob passed")
}

func TestEncoderJson(t *testing.T) {
	t.Log("TestSetEncoderJson")

	SetEncoderJson()

	if encoder != ModeJson {
		t.Errorf("expected ModeJson, got %d", encoder)
	}

	n := obj{}
	n.Message = "Hello, Json!"

	data, err := EncodeMessage(n)
	if err != nil {
		t.Error(err)
	}

	if len(data) == 0 {
		t.Error("expected data, got empty")
	}

	n2 := obj{}
	if err = DecodeMessage(data, &n2); err != nil {
		t.Error(err)
	}

	t.Log("TestSetEncoderJson passed")
}

func TestEncoderYaml(t *testing.T) {
	t.Log("TestSetEncoderYaml")

	SetEncoderYaml()

	if encoder != ModeYaml {
		t.Errorf("expected ModeYaml, got %d", encoder)
	}

	n := obj{}
	n.Message = "Hello, Yaml!"

	data, err := EncodeMessage(n)
	if err != nil {
		t.Error(err)
	}

	if len(data) == 0 {
		t.Error("expected data, got empty")
	}

	n2 := obj{}
	if err = DecodeMessage(data, &n2); err != nil {
		t.Error(err)
	}

	t.Log("TestSetEncoderYaml passed")
}

func TestEncoderToml(t *testing.T) {
	t.Log("TestSetEncoderToml")

	SetEncoderToml()

	if encoder != ModeToml {
		t.Errorf("expected ModeToml, got %d", encoder)
	}

	n := obj{}
	n.Message = "Hello, Toml!"

	data, err := EncodeMessage(n)
	if err != nil {
		t.Error(err)
	}

	if len(data) == 0 {
		t.Error("expected data, got empty")
	}

	n2 := obj{}
	if err = DecodeMessage(data, &n2); err != nil {
		t.Error(err)
	}

	t.Log("TestSetEncoderToml passed")
}

func TestEncoderGobFile(t *testing.T) {
	t.Log("TestSetEncoderGobFile")

	SetEncoderGob()

	if encoder != ModeGob {
		t.Errorf("expected ModeGob, got %d", encoder)
	}

	data, err := os.ReadFile("testdata/hello.gob")
	if err != nil {
		t.Error(err)
	}

	n := obj{}

	if err = DecodeMessage(data, &n); err != nil {
		t.Error(err)
	}

	if n.Message != "Hello, Gob!" {
		t.Errorf("expected Hello, Gob!, got %s", n.Message)
	}

	t.Log("TestSetEncoderGobFile passed")
}

func TestEncoderJsonFile(t *testing.T) {
	t.Log("TestSetEncoderJsonFile")

	SetEncoderJson()

	if encoder != ModeJson {
		t.Errorf("expected ModeJson, got %d", encoder)
	}

	data, err := os.ReadFile("testdata/hello.json")
	if err != nil {
		t.Error(err)
	}

	n := obj{}

	if err = DecodeMessage(data, &n); err != nil {
		t.Error(err)
	}

	if n.Message != "Hello, Json!" {
		t.Errorf("expected Hello, Json!, got %s", n.Message)
	}

	t.Log("TestSetEncoderJsonFile passed")
}
