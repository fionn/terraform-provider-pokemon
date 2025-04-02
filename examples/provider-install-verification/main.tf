terraform {
  required_providers {
    pokemon = {
      source  = "fionn/pokemon"
      version = "x.y.z"
    }
  }
}

provider "pokemon" {}

data "pokemon" "squirtle" {
  id = 7
}

data "pokemon" "eevee" {
  id = 133
}

output "squirtle" {
  value = data.pokemon.squirtle
}

output "eevee" {
  value = data.pokemon.eevee
}
