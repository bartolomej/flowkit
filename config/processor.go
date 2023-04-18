/*
 * Flow CLI
 *
 * Copyright 2019 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"encoding/json"
)

// processorRun all pre-processors.
func processorRun(raw []byte) []byte {
	type config struct {
		Accounts    map[string]map[string]any `json:"accounts,omitempty"`
		Contracts   any                       `json:"contracts,omitempty"`
		Networks    any                       `json:"networks,omitempty"`
		Deployments any                       `json:"deployments,omitempty"`
		Emulators   any                       `json:"emulators,omitempty"`
	}

	var conf config
	_ = json.Unmarshal(raw, &conf)

	raw, _ = json.Marshal(conf)
	return raw
}
