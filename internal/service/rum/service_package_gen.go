// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package rum

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	rum_sdkv2 "github.com/aws/aws-sdk-go-v2/service/rum"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceAppMonitor,
			TypeName: "aws_rum_app_monitor",
			Name:     "App Monitor",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceMetricsDestination,
			TypeName: "aws_rum_metrics_destination",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.RUM
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*rum_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return rum_sdkv2.NewFromConfig(cfg, func(o *rum_sdkv2.Options) {
		if endpoint := config[names.AttrEndpoint].(string); endpoint != "" {
			tflog.Debug(ctx, "setting endpoint", map[string]any{
				"tf_aws.endpoint": endpoint,
			})
			o.BaseEndpoint = aws_sdkv2.String(endpoint)

			if o.EndpointOptions.UseFIPSEndpoint == aws_sdkv2.FIPSEndpointStateEnabled {
				tflog.Debug(ctx, "endpoint set, ignoring UseFIPSEndpoint setting")
				o.EndpointOptions.UseFIPSEndpoint = aws_sdkv2.FIPSEndpointStateDisabled
			}
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
