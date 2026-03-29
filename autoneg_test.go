package goautoneg

import (
	"testing"
)

var chrome = "application/xml,application/xhtml+xml,text/html;q=0.9,text/plain;q=0.8,image/png,*/*;q=0.5"

func TestParseAccept(t *testing.T) {
	alternatives := []string{"text/html", "image/png"}
	content_type := Negotiate(chrome, alternatives)
	if content_type != "image/png" {
		t.Errorf("got %s expected image/png", content_type)
	}

	alternatives = []string{"text/html", "text/plain", "text/n3"}
	content_type = Negotiate(chrome, alternatives)
	if content_type != "text/html" {
		t.Errorf("got %s expected text/html", content_type)
	}

	alternatives = []string{"text/n3", "text/plain"}
	content_type = Negotiate(chrome, alternatives)
	if content_type != "text/plain" {
		t.Errorf("got %s expected text/plain", content_type)
	}

	alternatives = []string{"text/n3", "application/rdf+xml"}
	content_type = Negotiate(chrome, alternatives)
	if content_type != "text/n3" {
		t.Errorf("got %s expected text/n3", content_type)
	}
}

func TestParseAcceptEmptyHeader(t *testing.T) {
	accept := ParseAccept("")

	if len(accept) != 1 {
		t.Fatalf("expected 1 accept clause, got %d", len(accept))
	}

	if accept[0].Type != "*" || accept[0].SubType != "*" {
		t.Fatalf("expected */*, got %s/%s", accept[0].Type, accept[0].SubType)
	}

	if accept[0].Q != 1.0 {
		t.Fatalf("expected q=1.0, got %f", accept[0].Q)
	}
}

func TestParseAcceptWhitespaceHeader(t *testing.T) {
	accept := ParseAccept("   \t  ")

	if len(accept) != 1 {
		t.Fatalf("expected 1 accept clause, got %d", len(accept))
	}

	if accept[0].Type != "*" || accept[0].SubType != "*" {
		t.Fatalf("expected */*, got %s/%s", accept[0].Type, accept[0].SubType)
	}

	if accept[0].Q != 1.0 {
		t.Fatalf("expected q=1.0, got %f", accept[0].Q)
	}
}

func TestNegotiateEmptyHeaderAllowsAnything(t *testing.T) {
	alternatives := []string{"application/json", "text/html"}

	contentType := Negotiate("", alternatives)
	if contentType == "" {
		t.Fatal("expected a match, got empty string")
	}
}

func TestNegotiateEmptyHeaderPrefersFirstAlternative(t *testing.T) {
	alternatives := []string{"application/json", "text/html"}

	contentType := Negotiate("", alternatives)
	if contentType != "application/json" {
		t.Errorf("got %s expected application/json", contentType)
	}
}

func TestNegotiateWhitespaceHeaderPrefersFirstAlternative(t *testing.T) {
	alternatives := []string{"application/json", "text/html"}

	contentType := Negotiate("   ", alternatives)
	if contentType != "application/json" {
		t.Errorf("got %s expected application/json", contentType)
	}
}
