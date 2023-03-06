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

package corelogic

import (
	"fmt"
	"github.com/corelayer/go-corelogic-nitro/pkg/resource/controllers"
	"github.com/corelayer/go-netscaler-nitro/pkg/appconfig"
	"github.com/corelayer/go-netscaler-nitro/pkg/nitro"
	"time"
)

type HttpProvider struct {
	environment appconfig.Environment
}

// NewCoreLogicHttpProvider returns a HTTPProvider instance with a configured list of hosts
func NewCoreLogicHttpProvider(environment appconfig.Environment) (*HttpProvider, error) {
	c := &HttpProvider{
		environment: environment,
	}

	return c, nil
}

func (p *HttpProvider) Present(domain string, token string, keyAuth string) error {
	var err error
	var clients map[string]*nitro.Client

	clients, err = p.environment.GetAllNitroClients()
	if err != nil {
		return err
	}

	var primary string
	primary, err = p.environment.GetPrimaryNodeName()
	if err != nil {
		return err
	}

	csvc := controllers.NewCoreLogicController(clients[primary])

	fmt.Printf("PRESENTING TOKEN\n")
	fmt.Println("")
	fmt.Printf("Domain: %s\n", domain)
	fmt.Printf("Token: %s\n", token)
	fmt.Printf("Authorization code: %s\n", keyAuth)
	fmt.Println("")
	fmt.Printf("Getting primary node for environment %s: %s", p.environment.Name, primary)
	_, err = csvc.ListCsVserver()
	fmt.Printf("Getting csvserver for domain %s\n", domain)
	fmt.Printf("Writing token %s and key %s to SM_ACME\n", token, keyAuth)
	fmt.Println("")
	fmt.Printf("bind policy stringmap SM_ACME \"csvserver=csvserver;type=http;token=%s\" \"authorization=%s;\"\n", token, keyAuth)
	fmt.Println("")
	fmt.Printf("Saving config\n")

	time.Sleep(15 * time.Second)
	return nil
}

func (p *HttpProvider) CleanUp(domain string, token string, keyAuth string) error {
	//var err error

	fmt.Printf("CLEANUP TOKEN\n")
	fmt.Println("")
	fmt.Printf("Domain: %s\n", domain)
	fmt.Printf("Token: %s\n", token)
	fmt.Println("")
	fmt.Printf("Getting primary node for environment %s", p.environment.Name)
	fmt.Printf("Getting csvserver for domain %s\n", domain)
	fmt.Printf("Removing token %s and key %s from SM_ACME\n", token, keyAuth)
	fmt.Println("")
	fmt.Printf("unbind policy stringmap SM_ACME \"csvserver=csvserver;type=http;token=%s\"\n", token)
	fmt.Println("")
	fmt.Printf("Saving config\n")

	time.Sleep(15 * time.Second)
	return nil
}
