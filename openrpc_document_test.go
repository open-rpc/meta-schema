package openrpc_document

import (
	"encoding/json"
	"github.com/go-test/deep"
	"io/ioutil"
	"testing"
)

// TestOpenRPCDocument_MarshalUnmarshal tests that marshaling and unmarshaling
// yield equivalent JSON information.
func TestOpenRPCDocument_MarshalUnmarshal(t *testing.T) {
	referenceData, err := ioutil.ReadFile("testdata/calculator_discovery.json")
	if err != nil {
		t.Fatal(err)
	}

	reference := &OpenrpcDocument{}
	err = json.Unmarshal(referenceData, reference)
	if err != nil {
		t.Fatal(err)
	}

	marshaledReferenceData, err := json.Marshal(reference)
	if err != nil {
		t.Fatal(err)
	}

	target := &OpenrpcDocument{}
	err = json.Unmarshal(marshaledReferenceData, &target)
	if err != nil {
		t.Fatal(err)
	}

	if diff := deep.Equal(reference, target); diff != nil {
		for _, d := range diff {
			t.Log(d)
		}
		t.Error("reference != target")

		t.Log("---------")
		t.Log(string(referenceData))

		marshaledTargetData, err := json.MarshalIndent(target, "", "    ")
		if err != nil {
			t.Fatal(err)
		}
		t.Log("---------")
		t.Log(string(marshaledTargetData))
	}
}
