// Copyright 2016 Claudemiro Alves Feitosa Neto. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package banner

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"testing"
	"time"
)

func Test_printBanner(t *testing.T) {
	bannerContent := `Hello, {{ .GoVersion }}, {{ .GOOS }}, {{ .GOARCH }}, {{ .NumCPU }}, {{ .GOPATH }}, {{ .GOROOT }}, {{ .Compiler }}`
	var buffer bytes.Buffer

	version := runtime.Version()
	goos := runtime.GOOS
	goarch := runtime.GOARCH
	numCPU := runtime.NumCPU()
	gopath := os.Getenv("GOPATH")
	goroot := runtime.GOROOT()
	compiler := runtime.Compiler

	expected := fmt.Sprintf(`Hello, %s, %s, %s, %d, %s, %s, %s`, version, goos, goarch, numCPU, gopath, goroot, compiler)

	show(&buffer, bannerContent)

	result, err := ioutil.ReadAll(&buffer)

	if err != nil {
		t.Error(err)
	}

	if string(result) != expected {
		t.Errorf("result != expected, got %s", string(result))
	}
}

func Test_printBanner_invalid(t *testing.T) {
	var buffer bytes.Buffer

	expected := ""

	show(&buffer, "{{}")

	result, err := ioutil.ReadAll(&buffer)

	if err != nil {
		t.Error(err)
	}

	if string(result) != expected {
		t.Errorf("result != expected, got %s", string(result))
	}
}

func Test_printBanner_flags(t *testing.T) {
	var buffer bytes.Buffer

	Init(&buffer, true, "test-banner.txt")

	expected := "Test Banner"

	result, err := ioutil.ReadAll(&buffer)

	if err != nil {
		t.Error(err)
	}

	if string(result) != expected {
		t.Errorf("result != expected, got %s", string(result))
	}
}

func Test_printBanner_invalid_file(t *testing.T) {
	var buffer bytes.Buffer

	Init(&buffer, true, "invalid.txt")

	expected := ""

	result, err := ioutil.ReadAll(&buffer)

	if err != nil {
		t.Error(err)
	}

	if string(result) != expected {
		t.Errorf("result != expected, got %s", string(result))
	}
}

func Test_printBanner_banner_disabled(t *testing.T) {
	var buffer bytes.Buffer

	Init(&buffer, false, "test-banner.txt")

	expected := ""

	result, err := ioutil.ReadAll(&buffer)

	if err != nil {
		t.Error(err)
	}

	if string(result) != expected {
		t.Errorf("result != expected, got %s", string(result))
	}
}

func Test_vars_Env(t *testing.T) {
	v := vars{}
	gopath := v.Env("GOPATH")

	expected := os.Getenv("GOPATH")

	if gopath != expected {
		t.Errorf("gopath != expected, got %s", gopath)
	}
}

func Test_vars_Now(t *testing.T) {
	v := vars{}
	gopath := v.Now("Monday, 2 Jan 2006")

	expected := time.Now().Format("Monday, 2 Jan 2006")

	if gopath != expected {
		t.Errorf("gopath != expected, got %s", gopath)
	}
}

func Test_SetLog(t *testing.T) {
	oldLogger := logger
	SetLog(log.New(os.Stderr, "", log.LstdFlags))

	if oldLogger == logger {
		t.Errorf("logger was changed, must not be equal")
	}
}

func Test_SetLog_nil(t *testing.T) {
	oldLogger := logger
	SetLog(nil)

	if oldLogger != logger {
		t.Errorf("logger was not changed, must be equal")
	}
}
