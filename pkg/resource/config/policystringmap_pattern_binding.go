/*
 * Copyright 2022 CoreLayer BV
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package config

import (
	"github.com/corelayer/go-netscaler-nitro/pkg/resource/config"
	"strings"
)

func ParseStringmapEntry(e config.PolicyStringmapPatternBinding) (map[string]string, error) {
	var err error
	elements := make(map[string]string, 5)

	for _, e := range strings.Split(e.Key, ";") {
		if key, value, found := strings.Cut(e, "="); found {
			elements[key] = value
		}
	}

	for _, e := range strings.Split(e.Value, ";") {
		if key, value, found := strings.Cut(e, "="); found {
			elements[key] = value
		}
	}

	return elements, err
}
