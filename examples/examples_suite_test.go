package main_test

import (
	"github.com/ejdem86/api2go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

var api *api2go.API

func TestExamples(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Examples Suite")
}
