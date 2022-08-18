// Copyright 2022 Juan Pablo Tosso and the OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

//go:build tinygo
// +build tinygo

package operators

import (
	"github.com/corazawaf/coraza/v3"
)

type inspectFile struct{}

func (*inspectFile) Init(coraza.RuleOperatorOptions) error { return nil }

func (*inspectFile) Evaluate(*coraza.Transaction, string) bool { return true }
