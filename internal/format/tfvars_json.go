/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package format

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/iancoleman/orderedmap"

	"github.com/terraform-docs/terraform-docs/internal/print"
	"github.com/terraform-docs/terraform-docs/internal/terraform"
)

// TfvarsJSON represents Terraform tfvars JSON format.
type TfvarsJSON struct{}

// NewTfvarsJSON returns new instance of TfvarsJSON.
func NewTfvarsJSON(settings *print.Settings) print.Engine {
	return &TfvarsJSON{}
}

// Print a Terraform module as Terraform tfvars JSON.
func (j *TfvarsJSON) Print(module *terraform.Module, settings *print.Settings) (string, error) {
	copy := orderedmap.New()
	copy.SetEscapeHTML(false)
	for _, i := range module.Inputs {
		copy.Set(i.Name, i.Default)
	}

	buffer := new(bytes.Buffer)

	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)

	err := encoder.Encode(copy)
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(buffer.String(), "\n"), nil
}

func init() {
	register(map[string]initializerFn{
		"tfvars json": NewTfvarsJSON,
	})
}
