package config

import (
	"path/filepath"
	"testing"
)

// This is the directory where our test fixtures are.
const fixtureDir = "./test-fixtures"

func TestConfigValidate(t *testing.T) {
	c := testConfig(t, "validate-good")
	if err := c.Validate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestConfigValidate_unknownVar(t *testing.T) {
	c := testConfig(t, "validate-unknownvar")
	if err := c.Validate(); err == nil {
		t.Fatal("should not be valid")
	}
}

func TestNewResourceVariable(t *testing.T) {
	v, err := NewResourceVariable("foo.bar.baz")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if v.Type != "foo" {
		t.Fatalf("bad: %#v", v)
	}
	if v.Name != "bar" {
		t.Fatalf("bad: %#v", v)
	}
	if v.Field != "baz" {
		t.Fatalf("bad: %#v", v)
	}

	if v.FullKey() != "foo.bar.baz" {
		t.Fatalf("bad: %#v", v)
	}
}

func TestNewUserVariable(t *testing.T) {
	v, err := NewUserVariable("var.bar")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if v.Name != "bar" {
		t.Fatalf("bad: %#v", v.Name)
	}
	if v.FullKey() != "var.bar" {
		t.Fatalf("bad: %#v", v)
	}
}

func TestProviderConfigName(t *testing.T) {
	pcs := map[string]*ProviderConfig{
		"aw":   new(ProviderConfig),
		"aws":  new(ProviderConfig),
		"a":    new(ProviderConfig),
		"gce_": new(ProviderConfig),
	}

	n := ProviderConfigName("aws_instance", pcs)
	if n != "aws" {
		t.Fatalf("bad: %s", n)
	}
}

func testConfig(t *testing.T, name string) *Config {
	c, err := Load(filepath.Join(fixtureDir, name, "main.tf"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	return c
}
