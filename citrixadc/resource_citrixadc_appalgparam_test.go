/*
Copyright 2016 Citrix Systems, Inc

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
package citrixadc

import (
	"fmt"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

const testAccAppalgparam_add = `

	resource "citrixadc_appalgparam" "tf_appalgparam" {
		pptpgreidletimeout = 8000
	}
`
const testAccAppalgparam_update = `

	resource "citrixadc_appalgparam" "tf_appalgparam" {
		pptpgreidletimeout = 9000
	}
`

func TestAccAppalgparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAppalgparam_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppalgparamExist("citrixadc_appalgparam.tf_appalgparam", nil),
					resource.TestCheckResourceAttr("citrixadc_appalgparam.tf_appalgparam", "pptpgreidletimeout", "8000"),
				),
			},
			resource.TestStep{
				Config: testAccAppalgparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppalgparamExist("citrixadc_appalgparam.tf_appalgparam", nil),
					resource.TestCheckResourceAttr("citrixadc_appalgparam.tf_appalgparam", "pptpgreidletimeout", "9000"),
				),
			},
		},
	})
}

func testAccCheckAppalgparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appalgparam name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Appalgparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appalgparam %s not found", n)
		}

		return nil
	}
}
