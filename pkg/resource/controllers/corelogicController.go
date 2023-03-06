/*
 * Copyright 2023 CoreLayer BV
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

package controllers

import (
	"fmt"
	"github.com/corelayer/go-netscaler-nitro/pkg/nitro"
	"github.com/corelayer/go-netscaler-nitro/pkg/resource/config"
	"github.com/corelayer/go-netscaler-nitro/pkg/resource/controllers"
)

type CoreLogicController struct {
	client *nitro.Client
}

func NewCoreLogicController(c *nitro.Client) CoreLogicController {
	return CoreLogicController{
		client: c,
	}
}

func (c *CoreLogicController) ListStringmapEntries(name string) ([]config.PolicyStringmapPatternBinding, error) {
	var err error
	var output []config.PolicyStringmapPatternBinding
	n := controllers.NewPolicyStringmapController(c.client)

	r, err := n.GetBindings(name, nil)
	if err != nil {
		return nil, err
	}

	if r.Data == nil {
		return nil, fmt.Errorf("no stringmap entries found in CS_CONTROL")
	}

	for _, v := range r.Data {
		output = append(output, v)
	}
	return output, nil
}

func (c *CoreLogicController) ListCsVserver() ([]config.CsVserver, error) {
	var err error
	var output []config.CsVserver
	n := controllers.NewCsVserverController(c.client)

	r, err := n.List(nil, []string{"name", "servicetype"})
	if err != nil {
		return nil, err
	}

	if r.Data == nil {
		return nil, fmt.Errorf("no csvservers found in response")
	}
	for _, v := range r.Data {
		output = append(output, v)
	}
	return output, nil
}

func (c *CoreLogicController) ListLbVserver() ([]config.LbVserver, error) {
	var err error
	var output []config.LbVserver
	n := controllers.NewLbVserverController(c.client)

	r, err := n.List(nil, []string{"name", "servicetype"})
	if err != nil {
		return nil, err
	}

	if r.Data == nil {
		return nil, fmt.Errorf("no lbvservers found in response")
	}
	for _, v := range r.Data {
		output = append(output, v)
	}
	return output, nil
}
