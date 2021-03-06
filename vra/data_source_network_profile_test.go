package vra

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"regexp"
	"testing"
)

func TestAccDataSourceVRANetworkProfile(t *testing.T) {
	rInt := acctest.RandInt()
	resourceName1 := "vra_network_profile.this"
	dataSourceName1 := "data.vra_network_profile.this"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckAWS(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVRANetworkProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceVRANetworkProfileNotFound(rInt),
				ExpectError: regexp.MustCompile("vra_network_profile filter did not match any network profile"),
			},
			{
				Config: testAccDataSourceVRANetworkProfileNameFilter(rInt),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(resourceName1, "id", dataSourceName1, "id"),
					resource.TestCheckResourceAttrPair(resourceName1, "description", dataSourceName1, "description"),
					resource.TestCheckResourceAttrPair(resourceName1, "name", dataSourceName1, "name"),
					resource.TestCheckResourceAttrPair(resourceName1, "isolation_type", dataSourceName1, "isolation_type"),
					resource.TestCheckResourceAttrPair(resourceName1, "region_id", dataSourceName1, "region_id"),
				),
			},
			{
				Config: testAccDataSourceVRANetworkProfileRegionIdFilter(rInt),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(resourceName1, "id", dataSourceName1, "id"),
					resource.TestCheckResourceAttrPair(resourceName1, "description", dataSourceName1, "description"),
					resource.TestCheckResourceAttrPair(resourceName1, "name", dataSourceName1, "name"),
					resource.TestCheckResourceAttrPair(resourceName1, "isolation_type", dataSourceName1, "isolation_type"),
					resource.TestCheckResourceAttrPair(resourceName1, "region_id", dataSourceName1, "region_id"),
				),
			},
			{
				Config: testAccDataSourceVRANetworkProfileById(rInt),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(resourceName1, "id", dataSourceName1, "id"),
					resource.TestCheckResourceAttrPair(resourceName1, "description", dataSourceName1, "description"),
					resource.TestCheckResourceAttrPair(resourceName1, "name", dataSourceName1, "name"),
					resource.TestCheckResourceAttrPair(resourceName1, "isolation_type", dataSourceName1, "isolation_type"),
					resource.TestCheckResourceAttrPair(resourceName1, "region_id", dataSourceName1, "region_id"),
				),
			},
		},
	})
}

func testAccDataSourceVRANetworkProfileNotFound(rInt int) string {
	return testAccCheckVRANetworkProfileConfig(rInt) + fmt.Sprintf(`
	data "vra_network_profile" "this" {
		filter = "name eq 'foobar'"
	}`)
}

func testAccDataSourceVRANetworkProfileNameFilter(rInt int) string {
	return testAccCheckVRANetworkProfileConfig(rInt) + fmt.Sprintf(`
	data "vra_network_profile" "this" {
		filter = "name eq '${vra_network_profile.this.name}'"
	}`)
}

func testAccDataSourceVRANetworkProfileRegionIdFilter(rInt int) string {
	return testAccCheckVRANetworkProfileConfig(rInt) + fmt.Sprintf(`
	data "vra_network_profile" "this" {
		filter = "regionId eq '${data.vra_region.this.id}'"
	}`)
}

func testAccDataSourceVRANetworkProfileById(rInt int) string {
	return testAccCheckVRANetworkProfileConfig(rInt) + fmt.Sprintf(`
	data "vra_network_profile" "this" {
		id = vra_network_profile.this.id
	}`)
}
