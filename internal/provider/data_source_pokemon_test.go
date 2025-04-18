package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccPokemonsDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `data "pokemon" "test" {id = 1}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.pokemon.test", "id", "1"),
					resource.TestCheckResourceAttr("data.pokemon.test", "name", "bulbasaur"),
					resource.TestCheckResourceAttrSet("data.pokemon.test", "types.#"),
					resource.TestCheckResourceAttr("data.pokemon.test", "height", "0.7"),
					resource.TestCheckResourceAttr("data.pokemon.test", "weight", "6.9"),
				),
			},
		},
	})
}
