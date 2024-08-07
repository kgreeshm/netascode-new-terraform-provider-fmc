// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccFmcExtendedACL(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_extended_acl.test", "name", "extended_acl_1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_extended_acl.test", "description", "My Extended Access Control List"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_extended_acl.test", "entries.0.action", "DENY"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_extended_acl.test", "entries.0.log_level", "WARNING"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_extended_acl.test", "entries.0.logging", "PER_ACCESS_LIST_ENTRY"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_extended_acl.test", "entries.0.log_interval_seconds", "120"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_extended_acl.test", "entries.0.source_network_literals.0.value", "10.1.1.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_extended_acl.test", "entries.0.source_network_literals.0.type", "Network"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_extended_acl.test", "entries.0.destination_network_literals.0.value", "10.2.2.2"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_extended_acl.test", "entries.0.destination_network_literals.0.type", "Host"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcExtendedACLPrerequisitesConfig + testAccFmcExtendedACLConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcExtendedACLPrerequisitesConfig + testAccFmcExtendedACLConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_extended_acl.test",
		ImportState:  true,
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccFmcExtendedACLPrerequisitesConfig = `
resource "fmc_network" "test" {
  name   = "mynetwork2"
  prefix = "10.0.0.0/24"
}

resource "fmc_host" "test" {
  name = "myhost2"
  ip   = "10.1.1.1"
}

resource "fmc_port" "test" {
  name = "myport2"
  protocol = "UDP"
  port = "65000"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcExtendedACLConfig_minimum() string {
	config := `resource "fmc_extended_acl" "test" {` + "\n"
	config += `	name = "extended_acl_1"` + "\n"
	config += `	entries = [{` + "\n"
	config += `		action = "PERMIT"` + "\n"
	config += `		logging = "DEFAULT"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcExtendedACLConfig_all() string {
	config := `resource "fmc_extended_acl" "test" {` + "\n"
	config += `	name = "extended_acl_1"` + "\n"
	config += `	description = "My Extended Access Control List"` + "\n"
	config += `	entries = [{` + "\n"
	config += `		action = "DENY"` + "\n"
	config += `		log_level = "WARNING"` + "\n"
	config += `		logging = "PER_ACCESS_LIST_ENTRY"` + "\n"
	config += `		log_interval_seconds = 120` + "\n"
	config += `		source_network_literals = [{` + "\n"
	config += `			value = "10.1.1.0/24"` + "\n"
	config += `			type = "Network"` + "\n"
	config += `		}]` + "\n"
	config += `		destination_network_literals = [{` + "\n"
	config += `			value = "10.2.2.2"` + "\n"
	config += `			type = "Host"` + "\n"
	config += `		}]` + "\n"
	config += `		source_network_objects = [{` + "\n"
	config += `			id = fmc_network.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		destination_network_objects = [{` + "\n"
	config += `			id = fmc_host.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		source_port_objects = [{` + "\n"
	config += `			id = fmc_port.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		destination_port_objects = [{` + "\n"
	config += `			id = fmc_port.test.id` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
