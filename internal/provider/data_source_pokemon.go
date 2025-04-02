package provider

import (
	"context"

	"github.com/mtslzr/pokeapi-go"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
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
	ID     types.Int32   `tfsdk:"id"`
	Name   types.String  `tfsdk:"name"`
	Types  types.List    `tfsdk:"types"`
	Height types.Float64 `tfsdk:"height"`
	Weight types.Float64 `tfsdk:"weight"`
}

func (d *pokemonDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A pokémon",
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
			"types": schema.ListAttribute{
				Description: "Pokémon types",
				ElementType: types.StringType,
				Computed:    true,
			},
			"height": schema.Float64Attribute{
				Description: "Height in meters",
				Computed:    true,
			},
			"weight": schema.Float64Attribute{
				Description: "Weight in kilograms",
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

	pokemonTypes := make([]attr.Value, len(pokemon.Types))
	for i, t := range pokemon.Types {
		pokemonTypes[i] = basetypes.NewStringValue(t.Type.Name)
	}

	pokemonTypesListValue, _ := types.ListValue(types.StringType, pokemonTypes)

	state = pokemonDataSourceModel{
		ID:     types.Int32Value(int32(pokemon.ID)),
		Name:   types.StringValue(pokemon.Name),
		Types:  pokemonTypesListValue,
		Height: types.Float64Value(float64(pokemon.Height) / 10),
		Weight: types.Float64Value(float64(pokemon.Weight) / 10),
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
