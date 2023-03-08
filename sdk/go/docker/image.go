// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package docker

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Builds a Docker Image and pushes to a Docker registry.
//
// ## Example Usage
// ### A Docker image build
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			demoImage, err := docker.NewImage(ctx, "demo-image", &docker.ImageArgs{
//				Build: pulumi.Any{
//					Context:    pulumi.String("."),
//					Dockerfile: pulumi.String("Dockerfile"),
//				},
//				ImageName: pulumi.String("username/image:tag1"),
//				SkipPush:  pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("imageName", demoImage.ImageName)
//			return nil
//		})
//	}
//
// ```
//
// {{% /examples %}}
type Image struct {
	pulumi.CustomResourceState

	// The fully qualified image name that was pushed to the registry.
	BaseImageName pulumi.StringPtrOutput `pulumi:"baseImageName"`
	// The fully qualified image name
	ImageName pulumi.StringPtrOutput `pulumi:"imageName"`
	// The name of the registry server hosting the image.
	RegistryServer pulumi.StringPtrOutput `pulumi:"registryServer"`
}

// NewImage registers a new resource with the given unique name, arguments, and options.
func NewImage(ctx *pulumi.Context,
	name string, args *ImageArgs, opts ...pulumi.ResourceOption) (*Image, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ImageName == nil {
		return nil, errors.New("invalid value for required argument 'ImageName'")
	}
	if args.Build != nil {
		args.Build = args.Build.ToDockerBuildPtrOutput().ApplyT(func(v *DockerBuild) *DockerBuild { return v.Defaults() }).(DockerBuildPtrOutput)
	}
	if isZero(args.SkipPush) {
		args.SkipPush = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("docker:image:Image"),
		},
	})
	opts = append(opts, aliases)
	var resource Image
	err := ctx.RegisterResource("docker:index/image:Image", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImage gets an existing Image resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageState, opts ...pulumi.ResourceOption) (*Image, error) {
	var resource Image
	err := ctx.ReadResource("docker:index/image:Image", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Image resources.
type imageState struct {
}

type ImageState struct {
}

func (ImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageState)(nil)).Elem()
}

type imageArgs struct {
	// The Docker build context
	Build *DockerBuild `pulumi:"build"`
	// The image name
	ImageName string `pulumi:"imageName"`
	// The registry to push the image to
	Registry *Registry `pulumi:"registry"`
	// A flag to skip a registry push.
	SkipPush *bool `pulumi:"skipPush"`
}

// The set of arguments for constructing a Image resource.
type ImageArgs struct {
	// The Docker build context
	Build DockerBuildPtrInput
	// The image name
	ImageName pulumi.StringInput
	// The registry to push the image to
	Registry RegistryPtrInput
	// A flag to skip a registry push.
	SkipPush pulumi.BoolPtrInput
}

func (ImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageArgs)(nil)).Elem()
}

type ImageInput interface {
	pulumi.Input

	ToImageOutput() ImageOutput
	ToImageOutputWithContext(ctx context.Context) ImageOutput
}

func (*Image) ElementType() reflect.Type {
	return reflect.TypeOf((**Image)(nil)).Elem()
}

func (i *Image) ToImageOutput() ImageOutput {
	return i.ToImageOutputWithContext(context.Background())
}

func (i *Image) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageOutput)
}

// ImageArrayInput is an input type that accepts ImageArray and ImageArrayOutput values.
// You can construct a concrete instance of `ImageArrayInput` via:
//
//	ImageArray{ ImageArgs{...} }
type ImageArrayInput interface {
	pulumi.Input

	ToImageArrayOutput() ImageArrayOutput
	ToImageArrayOutputWithContext(context.Context) ImageArrayOutput
}

type ImageArray []ImageInput

func (ImageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Image)(nil)).Elem()
}

func (i ImageArray) ToImageArrayOutput() ImageArrayOutput {
	return i.ToImageArrayOutputWithContext(context.Background())
}

func (i ImageArray) ToImageArrayOutputWithContext(ctx context.Context) ImageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageArrayOutput)
}

// ImageMapInput is an input type that accepts ImageMap and ImageMapOutput values.
// You can construct a concrete instance of `ImageMapInput` via:
//
//	ImageMap{ "key": ImageArgs{...} }
type ImageMapInput interface {
	pulumi.Input

	ToImageMapOutput() ImageMapOutput
	ToImageMapOutputWithContext(context.Context) ImageMapOutput
}

type ImageMap map[string]ImageInput

func (ImageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Image)(nil)).Elem()
}

func (i ImageMap) ToImageMapOutput() ImageMapOutput {
	return i.ToImageMapOutputWithContext(context.Background())
}

func (i ImageMap) ToImageMapOutputWithContext(ctx context.Context) ImageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageMapOutput)
}

type ImageOutput struct{ *pulumi.OutputState }

func (ImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Image)(nil)).Elem()
}

func (o ImageOutput) ToImageOutput() ImageOutput {
	return o
}

func (o ImageOutput) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return o
}

// The fully qualified image name that was pushed to the registry.
func (o ImageOutput) BaseImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.BaseImageName }).(pulumi.StringPtrOutput)
}

// The fully qualified image name
func (o ImageOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.ImageName }).(pulumi.StringPtrOutput)
}

// The name of the registry server hosting the image.
func (o ImageOutput) RegistryServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.RegistryServer }).(pulumi.StringPtrOutput)
}

type ImageArrayOutput struct{ *pulumi.OutputState }

func (ImageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Image)(nil)).Elem()
}

func (o ImageArrayOutput) ToImageArrayOutput() ImageArrayOutput {
	return o
}

func (o ImageArrayOutput) ToImageArrayOutputWithContext(ctx context.Context) ImageArrayOutput {
	return o
}

func (o ImageArrayOutput) Index(i pulumi.IntInput) ImageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Image {
		return vs[0].([]*Image)[vs[1].(int)]
	}).(ImageOutput)
}

type ImageMapOutput struct{ *pulumi.OutputState }

func (ImageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Image)(nil)).Elem()
}

func (o ImageMapOutput) ToImageMapOutput() ImageMapOutput {
	return o
}

func (o ImageMapOutput) ToImageMapOutputWithContext(ctx context.Context) ImageMapOutput {
	return o
}

func (o ImageMapOutput) MapIndex(k pulumi.StringInput) ImageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Image {
		return vs[0].(map[string]*Image)[vs[1].(string)]
	}).(ImageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImageInput)(nil)).Elem(), &Image{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageArrayInput)(nil)).Elem(), ImageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageMapInput)(nil)).Elem(), ImageMap{})
	pulumi.RegisterOutputType(ImageOutput{})
	pulumi.RegisterOutputType(ImageArrayOutput{})
	pulumi.RegisterOutputType(ImageMapOutput{})
}
