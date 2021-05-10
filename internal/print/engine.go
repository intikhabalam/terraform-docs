/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package print

import (
	"github.com/terraform-docs/terraform-docs/internal/terraform"
)

// Engine represents a format engine (e.g. json, table, yaml, ...)
type Engine interface {
	Generate(*terraform.Module) (*Generator, error)
}
