package gcpurl_test

import (
	"reflect"
	"testing"

	"github.com/goccy/go-gcpurl"
)

func TestParseURL(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		url, err := gcpurl.ParseURL(`https://us-central1-awesome-project-id.cloudfunctions.net`)
		if err != nil {
			t.Fatal(err)
		}
		if url.Host != "us-central1-awesome-project-id.cloudfunctions.net" {
			t.Fatalf("failed to get original url")
		}
		if url.Region != "us-central1" {
			t.Fatalf("failed to parse region")
		}
		if url.ProjectID != "awesome-project-id" {
			t.Fatalf("failed to parse project id")
		}
		if url.ServiceDomain != "cloudfunctions.net" {
			t.Fatalf("failed to parse service domain")
		}
	})
	t.Run("failure", func(t *testing.T) {
		if _, err := gcpurl.ParseURL(`https://example.com`); err == nil {
			t.Fatalf("expected error")
		}
	})
}

func TestZone(t *testing.T) {
	if !reflect.DeepEqual(gcpurl.Region("us-central1").Zones(), []string{"a", "b", "c", "f"}) {
		t.Fatalf("failed to get zones from region")
	}
}
