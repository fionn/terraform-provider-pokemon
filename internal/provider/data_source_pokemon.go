package provider

import (
	"context"

	"github.com/mtslzr/pokeapi-go"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type pokemonDataSource struct {
}

func NewPokemonDataSource() datasource.DataSource {
	return &pokemonDataSource{}
}

func (d *pokemonDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	// If we wanted to have other data or resource types, we'd add a postfix
	// here like "_pokemon", etc.
	resp.TypeName = req.ProviderTypeName
}

type pokemonDataSourceModel struct {
	ID   types.Int32  `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func (d *pokemonDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// TODO: expand attributes.
		Attributes: map[string]schema.Attribute{
			"id": schema.Int32Attribute{
				Description: "Pokédex number",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "Pokémon name",
				Computed:    true,
			},
		},
	}
}

func (d *pokemonDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state pokemonDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	pokemon, err := pokeapi.Pokemon(state.ID.String())
	if err != nil {
		resp.Diagnostics.AddError("Cannot find Pokémon", err.Error())
		return
	}

	state = pokemonDataSourceModel{
		ID:   types.Int32Value(int32(pokemon.ID)),
		Name: types.StringValue(pokemon.Name),
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
