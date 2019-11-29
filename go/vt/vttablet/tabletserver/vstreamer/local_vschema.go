/*
Copyright 2019 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vstreamer

import (
	"fmt"

	"vitess.io/vitess/go/vt/vtgate/vindexes"
)

// localVSchema provides vschema behavior specific to vstreamer.
// Tables are searched within keyspace, but vindexes can be referenced
// outside the current keyspace.
type localVSchema struct {
	keyspace string
	vschema  *vindexes.VSchema
}

func (lvs *localVSchema) FindTable(tablename string) (*vindexes.Table, error) {
	ks, ok := lvs.vschema.Keyspaces[lvs.keyspace]
	if !ok {
		return nil, fmt.Errorf("keyspace %s not found in vschema", lvs.keyspace)
	}
	table := ks.Tables[tablename]
	if table == nil {
		return nil, fmt.Errorf("table %s not found", tablename)
	}
	return table, nil
}
