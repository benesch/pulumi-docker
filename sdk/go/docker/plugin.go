// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package docker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages the lifecycle of a Docker plugin.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-docker/sdk/v2/go/docker"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := docker.NewPlugin(ctx, "sample_volume_plugin", nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-docker/sdk/v2/go/docker"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := docker.NewPlugin(ctx, "sample_volume_plugin", &docker.PluginArgs{
// 			Alias:         pulumi.String("sample-volume-plugin"),
// 			EnableTimeout: pulumi.Int(60),
// 			Enabled:       pulumi.Bool(false),
// 			Envs: pulumi.StringArray{
// 				pulumi.String("DEBUG=1"),
// 			},
// 			ForceDestroy:        pulumi.Bool(true),
// 			ForceDisable:        pulumi.Bool(true),
// 			GrantAllPermissions: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Docker plugins can be imported using the long id, e.g. for a plugin `tiborvass/sample-volume-plugin:latest`
//
// ```sh
//  $ pulumi import docker:index/plugin:Plugin sample-volume-plugin $(docker plugin inspect -f "{{.ID}}" tiborvass/sample-volume-plugin:latest)
// ```
type Plugin struct {
	pulumi.CustomResourceState

	// The alias of the Docker plugin. If the tag is omitted, `:latest` is complemented to the attribute value.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// HTTP client timeout to enable the plugin.
	EnableTimeout pulumi.IntPtrOutput `pulumi:"enableTimeout"`
	// If true, the plugin is enabled. The default value is `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// . The environment variables.
	Envs pulumi.StringArrayOutput `pulumi:"envs"`
	// If true, the plugin is removed forcibly when the plugin is removed.
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// If true, then the plugin is disabled forcibly when the plugin is disabled.
	ForceDisable pulumi.BoolPtrOutput `pulumi:"forceDisable"`
	// If true, grant all permissions necessary to run the plugin. This attribute conflicts with `grantPermissions`.
	GrantAllPermissions pulumi.BoolPtrOutput `pulumi:"grantAllPermissions"`
	// grant permissions necessary to run the plugin. This attribute conflicts with `grantAllPermissions`. See grantPermissions below for details.
	GrantPermissions PluginGrantPermissionArrayOutput `pulumi:"grantPermissions"`
	// Docker Plugin name
	Name pulumi.StringOutput `pulumi:"name"`
	// (string) The plugin reference.
	PluginReference pulumi.StringOutput `pulumi:"pluginReference"`
}

// NewPlugin registers a new resource with the given unique name, arguments, and options.
func NewPlugin(ctx *pulumi.Context,
	name string, args *PluginArgs, opts ...pulumi.ResourceOption) (*Plugin, error) {
	if args == nil {
		args = &PluginArgs{}
	}

	var resource Plugin
	err := ctx.RegisterResource("docker:index/plugin:Plugin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPlugin gets an existing Plugin resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPlugin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PluginState, opts ...pulumi.ResourceOption) (*Plugin, error) {
	var resource Plugin
	err := ctx.ReadResource("docker:index/plugin:Plugin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Plugin resources.
type pluginState struct {
	// The alias of the Docker plugin. If the tag is omitted, `:latest` is complemented to the attribute value.
	Alias *string `pulumi:"alias"`
	// HTTP client timeout to enable the plugin.
	EnableTimeout *int `pulumi:"enableTimeout"`
	// If true, the plugin is enabled. The default value is `true`.
	Enabled *bool `pulumi:"enabled"`
	// . The environment variables.
	Envs []string `pulumi:"envs"`
	// If true, the plugin is removed forcibly when the plugin is removed.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// If true, then the plugin is disabled forcibly when the plugin is disabled.
	ForceDisable *bool `pulumi:"forceDisable"`
	// If true, grant all permissions necessary to run the plugin. This attribute conflicts with `grantPermissions`.
	GrantAllPermissions *bool `pulumi:"grantAllPermissions"`
	// grant permissions necessary to run the plugin. This attribute conflicts with `grantAllPermissions`. See grantPermissions below for details.
	GrantPermissions []PluginGrantPermission `pulumi:"grantPermissions"`
	// Docker Plugin name
	Name *string `pulumi:"name"`
	// (string) The plugin reference.
	PluginReference *string `pulumi:"pluginReference"`
}

type PluginState struct {
	// The alias of the Docker plugin. If the tag is omitted, `:latest` is complemented to the attribute value.
	Alias pulumi.StringPtrInput
	// HTTP client timeout to enable the plugin.
	EnableTimeout pulumi.IntPtrInput
	// If true, the plugin is enabled. The default value is `true`.
	Enabled pulumi.BoolPtrInput
	// . The environment variables.
	Envs pulumi.StringArrayInput
	// If true, the plugin is removed forcibly when the plugin is removed.
	ForceDestroy pulumi.BoolPtrInput
	// If true, then the plugin is disabled forcibly when the plugin is disabled.
	ForceDisable pulumi.BoolPtrInput
	// If true, grant all permissions necessary to run the plugin. This attribute conflicts with `grantPermissions`.
	GrantAllPermissions pulumi.BoolPtrInput
	// grant permissions necessary to run the plugin. This attribute conflicts with `grantAllPermissions`. See grantPermissions below for details.
	GrantPermissions PluginGrantPermissionArrayInput
	// Docker Plugin name
	Name pulumi.StringPtrInput
	// (string) The plugin reference.
	PluginReference pulumi.StringPtrInput
}

func (PluginState) ElementType() reflect.Type {
	return reflect.TypeOf((*pluginState)(nil)).Elem()
}

type pluginArgs struct {
	// The alias of the Docker plugin. If the tag is omitted, `:latest` is complemented to the attribute value.
	Alias *string `pulumi:"alias"`
	// HTTP client timeout to enable the plugin.
	EnableTimeout *int `pulumi:"enableTimeout"`
	// If true, the plugin is enabled. The default value is `true`.
	Enabled *bool `pulumi:"enabled"`
	// . The environment variables.
	Envs []string `pulumi:"envs"`
	// If true, the plugin is removed forcibly when the plugin is removed.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// If true, then the plugin is disabled forcibly when the plugin is disabled.
	ForceDisable *bool `pulumi:"forceDisable"`
	// If true, grant all permissions necessary to run the plugin. This attribute conflicts with `grantPermissions`.
	GrantAllPermissions *bool `pulumi:"grantAllPermissions"`
	// grant permissions necessary to run the plugin. This attribute conflicts with `grantAllPermissions`. See grantPermissions below for details.
	GrantPermissions []PluginGrantPermission `pulumi:"grantPermissions"`
	// Docker Plugin name
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Plugin resource.
type PluginArgs struct {
	// The alias of the Docker plugin. If the tag is omitted, `:latest` is complemented to the attribute value.
	Alias pulumi.StringPtrInput
	// HTTP client timeout to enable the plugin.
	EnableTimeout pulumi.IntPtrInput
	// If true, the plugin is enabled. The default value is `true`.
	Enabled pulumi.BoolPtrInput
	// . The environment variables.
	Envs pulumi.StringArrayInput
	// If true, the plugin is removed forcibly when the plugin is removed.
	ForceDestroy pulumi.BoolPtrInput
	// If true, then the plugin is disabled forcibly when the plugin is disabled.
	ForceDisable pulumi.BoolPtrInput
	// If true, grant all permissions necessary to run the plugin. This attribute conflicts with `grantPermissions`.
	GrantAllPermissions pulumi.BoolPtrInput
	// grant permissions necessary to run the plugin. This attribute conflicts with `grantAllPermissions`. See grantPermissions below for details.
	GrantPermissions PluginGrantPermissionArrayInput
	// Docker Plugin name
	Name pulumi.StringPtrInput
}

func (PluginArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pluginArgs)(nil)).Elem()
}

type PluginInput interface {
	pulumi.Input

	ToPluginOutput() PluginOutput
	ToPluginOutputWithContext(ctx context.Context) PluginOutput
}

func (Plugin) ElementType() reflect.Type {
	return reflect.TypeOf((*Plugin)(nil)).Elem()
}

func (i Plugin) ToPluginOutput() PluginOutput {
	return i.ToPluginOutputWithContext(context.Background())
}

func (i Plugin) ToPluginOutputWithContext(ctx context.Context) PluginOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PluginOutput)
}

type PluginOutput struct {
	*pulumi.OutputState
}

func (PluginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PluginOutput)(nil)).Elem()
}

func (o PluginOutput) ToPluginOutput() PluginOutput {
	return o
}

func (o PluginOutput) ToPluginOutputWithContext(ctx context.Context) PluginOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PluginOutput{})
}
