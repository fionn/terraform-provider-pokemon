package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type pokemonProvider struct {
	version string
}

type pokemonProviderModel struct {
}

func (p *pokemonProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "pokemon"
	resp.Version = p.version
}

func (p *pokemonProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (p *pokemonProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	tflog.Debug(ctx, "Configuring Pokémon client")
	var config pokemonProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (p *pokemonProvider) Resources(ctx context.Context) []func() resource.Resource {
	return nil
}

func (p *pokemonProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewPokemonDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &pokemonProvider{
			version: version,
		}
	}
}
