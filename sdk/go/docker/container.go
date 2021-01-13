// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package docker

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages the lifecycle of a Docker container.
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
// 		ubuntuRemoteImage, err := docker.NewRemoteImage(ctx, "ubuntuRemoteImage", &docker.RemoteImageArgs{
// 			Name: pulumi.String("ubuntu:precise"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = docker.NewContainer(ctx, "ubuntuContainer", &docker.ContainerArgs{
// 			Image: ubuntuRemoteImage.Latest,
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
// Docker containers can be imported using the long id, e.g. for a container named `foo`
//
// ```sh
//  $ pulumi import docker:index/container:Container foo $(docker inspect -f {{.ID}} foo)
// ```
type Container struct {
	pulumi.CustomResourceState

	// If true attach to the container after its creation and waits the end of his execution.
	Attach pulumi.BoolPtrOutput `pulumi:"attach"`
	// The network bridge of the container as read from its NetworkSettings.
	Bridge pulumi.StringOutput `pulumi:"bridge"`
	// See Capabilities below for details.
	Capabilities ContainerCapabilitiesPtrOutput `pulumi:"capabilities"`
	// The command to use to start the
	// container. For example, to run `/usr/bin/myprogram -f baz.conf` set the
	// command to be `["/usr/bin/myprogram", "-f", "baz.conf"]`.
	Command pulumi.StringArrayOutput `pulumi:"command"`
	// The logs of the container if its execution is done (`attach` must be disabled).
	ContainerLogs pulumi.StringOutput `pulumi:"containerLogs"`
	// A comma-separated list or hyphen-separated range of CPUs a container can use, e.g. `0-1`.
	CpuSet pulumi.StringPtrOutput `pulumi:"cpuSet"`
	// CPU shares (relative weight) for the container.
	CpuShares pulumi.IntPtrOutput `pulumi:"cpuShares"`
	// If defined will attempt to stop the container before destroying. Container will be destroyed after `n` seconds or on successful stop.
	DestroyGraceSeconds pulumi.IntPtrOutput `pulumi:"destroyGraceSeconds"`
	// See Devices below for details.
	Devices ContainerDeviceArrayOutput `pulumi:"devices"`
	// Set of DNS servers.
	Dns pulumi.StringArrayOutput `pulumi:"dns"`
	// Set of DNS options used by the DNS provider(s), see `resolv.conf` documentation for valid list of options.
	DnsOpts pulumi.StringArrayOutput `pulumi:"dnsOpts"`
	// Set of DNS search domains that are used when bare unqualified hostnames are used inside of the container.
	DnsSearches pulumi.StringArrayOutput `pulumi:"dnsSearches"`
	// Domain name of the container.
	Domainname pulumi.StringPtrOutput `pulumi:"domainname"`
	// The command to use as the
	// Entrypoint for the container. The Entrypoint allows you to configure a
	// container to run as an executable. For example, to run `/usr/bin/myprogram`
	// when starting a container, set the entrypoint to be
	// `["/usr/bin/myprogram"]`.
	Entrypoints pulumi.StringArrayOutput `pulumi:"entrypoints"`
	// Environment variables to set.
	Envs pulumi.StringArrayOutput `pulumi:"envs"`
	// The exit code of the container if its execution is done (`mustRun` must be disabled).
	ExitCode pulumi.IntOutput `pulumi:"exitCode"`
	// *Deprecated:* Use `networkData` instead. The network gateway of the container as read from its
	// NetworkSettings.
	//
	// Deprecated: Use gateway from ip_adresses_data instead. This field exposes the data of the container's first network.
	Gateway pulumi.StringOutput `pulumi:"gateway"`
	// Add additional groups to run as.
	GroupAdds pulumi.StringArrayOutput `pulumi:"groupAdds"`
	// See Healthcheck below for details.
	Healthcheck ContainerHealthcheckOutput `pulumi:"healthcheck"`
	// Hostname of the container.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// Hostname to add.
	Hosts ContainerHostArrayOutput `pulumi:"hosts"`
	// The ID of the image to back this container.
	// The easiest way to get this value is to use the `RemoteImage` resource
	// as is shown in the example above.
	Image pulumi.StringOutput `pulumi:"image"`
	// Configured whether an init process should be injected for this container. If unset this will default to the `dockerd` defaults.
	Init pulumi.BoolOutput `pulumi:"init"`
	// *Deprecated:* Use `networkData` instead. The IP address of the container's first network it.
	//
	// Deprecated: Use ip_adresses_data instead. This field exposes the data of the container's first network.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// *Deprecated:* Use `networkData` instead. The IP prefix length of the container as read from its
	// NetworkSettings.
	//
	// Deprecated: Use ip_prefix_length from ip_adresses_data instead. This field exposes the data of the container's first network.
	IpPrefixLength pulumi.IntOutput `pulumi:"ipPrefixLength"`
	// IPC sharing mode for the container. Possible values are: `none`, `private`, `shareable`, `container:<name|id>` or `host`.
	IpcMode pulumi.StringOutput `pulumi:"ipcMode"`
	// Adding labels.
	Labels ContainerLabelArrayOutput `pulumi:"labels"`
	// Set of links for link based
	// connectivity between containers that are running on the same host.
	//
	// Deprecated: The --link flag is a legacy feature of Docker. It may eventually be removed.
	Links pulumi.StringArrayOutput `pulumi:"links"`
	// The logging driver to use for the container.
	// Defaults to "json-file".
	LogDriver pulumi.StringPtrOutput `pulumi:"logDriver"`
	// Key/value pairs to use as options for
	// the logging driver.
	LogOpts pulumi.MapOutput `pulumi:"logOpts"`
	// Save the container logs (`attach` must be enabled).
	Logs pulumi.BoolPtrOutput `pulumi:"logs"`
	// The maximum amount of times to an attempt
	// a restart when `restart` is set to "on-failure"
	MaxRetryCount pulumi.IntPtrOutput `pulumi:"maxRetryCount"`
	// The memory limit for the container in MBs.
	Memory     pulumi.IntPtrOutput `pulumi:"memory"`
	MemorySwap pulumi.IntPtrOutput `pulumi:"memorySwap"`
	// See Mounts below for details.
	Mounts  ContainerMountArrayOutput `pulumi:"mounts"`
	MustRun pulumi.BoolPtrOutput      `pulumi:"mustRun"`
	Name    pulumi.StringOutput       `pulumi:"name"`
	// Network aliases of the container for user-defined networks only. *Deprecated:* use `networksAdvanced` instead.
	//
	// Deprecated: Use networks_advanced instead. Will be removed in v2.0.0
	NetworkAliases pulumi.StringArrayOutput `pulumi:"networkAliases"`
	// (Map of a block) The IP addresses of the container on each
	// network. Key are the network names, values are the IP addresses.
	NetworkDatas ContainerNetworkDataArrayOutput `pulumi:"networkDatas"`
	// Network mode of the container.
	NetworkMode pulumi.StringPtrOutput `pulumi:"networkMode"`
	// Id of the networks in which the
	// container is. *Deprecated:* use `networksAdvanced` instead.
	//
	// Deprecated: Use networks_advanced instead. Will be removed in v2.0.0
	Networks pulumi.StringArrayOutput `pulumi:"networks"`
	// See Networks Advanced below for details. If this block has priority to the deprecated `networkAlias` and `network` properties.
	NetworksAdvanced ContainerNetworksAdvancedArrayOutput `pulumi:"networksAdvanced"`
	// The PID (Process) Namespace mode for the container. Either `container:<name|id>` or `host`.
	PidMode pulumi.StringPtrOutput `pulumi:"pidMode"`
	// See Ports below for details.
	Ports ContainerPortArrayOutput `pulumi:"ports"`
	// Run container in privileged mode.
	Privileged pulumi.BoolPtrOutput `pulumi:"privileged"`
	// Publish all ports of the container.
	PublishAllPorts pulumi.BoolPtrOutput `pulumi:"publishAllPorts"`
	// If true, this volume will be readonly.
	// Defaults to false.
	ReadOnly      pulumi.BoolPtrOutput `pulumi:"readOnly"`
	RemoveVolumes pulumi.BoolPtrOutput `pulumi:"removeVolumes"`
	// The restart policy for the container. Must be
	// one of "no", "on-failure", "always", "unless-stopped".
	Restart pulumi.StringPtrOutput `pulumi:"restart"`
	Rm      pulumi.BoolPtrOutput   `pulumi:"rm"`
	// Set of string values to customize labels for MLS systems, such as SELinux. See https://docs.docker.com/engine/reference/run/#security-configuration.
	SecurityOpts pulumi.StringArrayOutput `pulumi:"securityOpts"`
	// Size of `/dev/shm` in MBs.
	ShmSize pulumi.IntOutput `pulumi:"shmSize"`
	// If true, then the Docker container will be
	// started after creation. If false, then the container is only created.
	Start pulumi.BoolPtrOutput `pulumi:"start"`
	// A map of kernel parameters (sysctls) to set in the container.
	Sysctls pulumi.MapOutput `pulumi:"sysctls"`
	// A map of container directories which should be replaced by `tmpfs mounts`, and their corresponding mount options.
	Tmpfs pulumi.MapOutput `pulumi:"tmpfs"`
	// See Ulimits below for
	// details.
	Ulimits ContainerUlimitArrayOutput `pulumi:"ulimits"`
	// See File Upload below for details.
	Uploads ContainerUploadArrayOutput `pulumi:"uploads"`
	// User used for run the first process. Format is
	// `user` or `user:group` which user and group can be passed literraly or
	// by name.
	User pulumi.StringPtrOutput `pulumi:"user"`
	// Sets the usernamespace mode for the container when usernamespace remapping option is enabled.
	UsernsMode pulumi.StringPtrOutput `pulumi:"usernsMode"`
	// See Volumes below for details.
	Volumes ContainerVolumeArrayOutput `pulumi:"volumes"`
	// The working directory for commands to run in
	WorkingDir pulumi.StringPtrOutput `pulumi:"workingDir"`
}

// NewContainer registers a new resource with the given unique name, arguments, and options.
func NewContainer(ctx *pulumi.Context,
	name string, args *ContainerArgs, opts ...pulumi.ResourceOption) (*Container, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Image == nil {
		return nil, errors.New("invalid value for required argument 'Image'")
	}
	var resource Container
	err := ctx.RegisterResource("docker:index/container:Container", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainer gets an existing Container resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerState, opts ...pulumi.ResourceOption) (*Container, error) {
	var resource Container
	err := ctx.ReadResource("docker:index/container:Container", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Container resources.
type containerState struct {
	// If true attach to the container after its creation and waits the end of his execution.
	Attach *bool `pulumi:"attach"`
	// The network bridge of the container as read from its NetworkSettings.
	Bridge *string `pulumi:"bridge"`
	// See Capabilities below for details.
	Capabilities *ContainerCapabilities `pulumi:"capabilities"`
	// The command to use to start the
	// container. For example, to run `/usr/bin/myprogram -f baz.conf` set the
	// command to be `["/usr/bin/myprogram", "-f", "baz.conf"]`.
	Command []string `pulumi:"command"`
	// The logs of the container if its execution is done (`attach` must be disabled).
	ContainerLogs *string `pulumi:"containerLogs"`
	// A comma-separated list or hyphen-separated range of CPUs a container can use, e.g. `0-1`.
	CpuSet *string `pulumi:"cpuSet"`
	// CPU shares (relative weight) for the container.
	CpuShares *int `pulumi:"cpuShares"`
	// If defined will attempt to stop the container before destroying. Container will be destroyed after `n` seconds or on successful stop.
	DestroyGraceSeconds *int `pulumi:"destroyGraceSeconds"`
	// See Devices below for details.
	Devices []ContainerDevice `pulumi:"devices"`
	// Set of DNS servers.
	Dns []string `pulumi:"dns"`
	// Set of DNS options used by the DNS provider(s), see `resolv.conf` documentation for valid list of options.
	DnsOpts []string `pulumi:"dnsOpts"`
	// Set of DNS search domains that are used when bare unqualified hostnames are used inside of the container.
	DnsSearches []string `pulumi:"dnsSearches"`
	// Domain name of the container.
	Domainname *string `pulumi:"domainname"`
	// The command to use as the
	// Entrypoint for the container. The Entrypoint allows you to configure a
	// container to run as an executable. For example, to run `/usr/bin/myprogram`
	// when starting a container, set the entrypoint to be
	// `["/usr/bin/myprogram"]`.
	Entrypoints []string `pulumi:"entrypoints"`
	// Environment variables to set.
	Envs []string `pulumi:"envs"`
	// The exit code of the container if its execution is done (`mustRun` must be disabled).
	ExitCode *int `pulumi:"exitCode"`
	// *Deprecated:* Use `networkData` instead. The network gateway of the container as read from its
	// NetworkSettings.
	//
	// Deprecated: Use gateway from ip_adresses_data instead. This field exposes the data of the container's first network.
	Gateway *string `pulumi:"gateway"`
	// Add additional groups to run as.
	GroupAdds []string `pulumi:"groupAdds"`
	// See Healthcheck below for details.
	Healthcheck *ContainerHealthcheck `pulumi:"healthcheck"`
	// Hostname of the container.
	Hostname *string `pulumi:"hostname"`
	// Hostname to add.
	Hosts []ContainerHost `pulumi:"hosts"`
	// The ID of the image to back this container.
	// The easiest way to get this value is to use the `RemoteImage` resource
	// as is shown in the example above.
	Image *string `pulumi:"image"`
	// Configured whether an init process should be injected for this container. If unset this will default to the `dockerd` defaults.
	Init *bool `pulumi:"init"`
	// *Deprecated:* Use `networkData` instead. The IP address of the container's first network it.
	//
	// Deprecated: Use ip_adresses_data instead. This field exposes the data of the container's first network.
	IpAddress *string `pulumi:"ipAddress"`
	// *Deprecated:* Use `networkData` instead. The IP prefix length of the container as read from its
	// NetworkSettings.
	//
	// Deprecated: Use ip_prefix_length from ip_adresses_data instead. This field exposes the data of the container's first network.
	IpPrefixLength *int `pulumi:"ipPrefixLength"`
	// IPC sharing mode for the container. Possible values are: `none`, `private`, `shareable`, `container:<name|id>` or `host`.
	IpcMode *string `pulumi:"ipcMode"`
	// Adding labels.
	Labels []ContainerLabel `pulumi:"labels"`
	// Set of links for link based
	// connectivity between containers that are running on the same host.
	//
	// Deprecated: The --link flag is a legacy feature of Docker. It may eventually be removed.
	Links []string `pulumi:"links"`
	// The logging driver to use for the container.
	// Defaults to "json-file".
	LogDriver *string `pulumi:"logDriver"`
	// Key/value pairs to use as options for
	// the logging driver.
	LogOpts map[string]interface{} `pulumi:"logOpts"`
	// Save the container logs (`attach` must be enabled).
	Logs *bool `pulumi:"logs"`
	// The maximum amount of times to an attempt
	// a restart when `restart` is set to "on-failure"
	MaxRetryCount *int `pulumi:"maxRetryCount"`
	// The memory limit for the container in MBs.
	Memory     *int `pulumi:"memory"`
	MemorySwap *int `pulumi:"memorySwap"`
	// See Mounts below for details.
	Mounts  []ContainerMount `pulumi:"mounts"`
	MustRun *bool            `pulumi:"mustRun"`
	Name    *string          `pulumi:"name"`
	// Network aliases of the container for user-defined networks only. *Deprecated:* use `networksAdvanced` instead.
	//
	// Deprecated: Use networks_advanced instead. Will be removed in v2.0.0
	NetworkAliases []string `pulumi:"networkAliases"`
	// (Map of a block) The IP addresses of the container on each
	// network. Key are the network names, values are the IP addresses.
	NetworkDatas []ContainerNetworkData `pulumi:"networkDatas"`
	// Network mode of the container.
	NetworkMode *string `pulumi:"networkMode"`
	// Id of the networks in which the
	// container is. *Deprecated:* use `networksAdvanced` instead.
	//
	// Deprecated: Use networks_advanced instead. Will be removed in v2.0.0
	Networks []string `pulumi:"networks"`
	// See Networks Advanced below for details. If this block has priority to the deprecated `networkAlias` and `network` properties.
	NetworksAdvanced []ContainerNetworksAdvanced `pulumi:"networksAdvanced"`
	// The PID (Process) Namespace mode for the container. Either `container:<name|id>` or `host`.
	PidMode *string `pulumi:"pidMode"`
	// See Ports below for details.
	Ports []ContainerPort `pulumi:"ports"`
	// Run container in privileged mode.
	Privileged *bool `pulumi:"privileged"`
	// Publish all ports of the container.
	PublishAllPorts *bool `pulumi:"publishAllPorts"`
	// If true, this volume will be readonly.
	// Defaults to false.
	ReadOnly      *bool `pulumi:"readOnly"`
	RemoveVolumes *bool `pulumi:"removeVolumes"`
	// The restart policy for the container. Must be
	// one of "no", "on-failure", "always", "unless-stopped".
	Restart *string `pulumi:"restart"`
	Rm      *bool   `pulumi:"rm"`
	// Set of string values to customize labels for MLS systems, such as SELinux. See https://docs.docker.com/engine/reference/run/#security-configuration.
	SecurityOpts []string `pulumi:"securityOpts"`
	// Size of `/dev/shm` in MBs.
	ShmSize *int `pulumi:"shmSize"`
	// If true, then the Docker container will be
	// started after creation. If false, then the container is only created.
	Start *bool `pulumi:"start"`
	// A map of kernel parameters (sysctls) to set in the container.
	Sysctls map[string]interface{} `pulumi:"sysctls"`
	// A map of container directories which should be replaced by `tmpfs mounts`, and their corresponding mount options.
	Tmpfs map[string]interface{} `pulumi:"tmpfs"`
	// See Ulimits below for
	// details.
	Ulimits []ContainerUlimit `pulumi:"ulimits"`
	// See File Upload below for details.
	Uploads []ContainerUpload `pulumi:"uploads"`
	// User used for run the first process. Format is
	// `user` or `user:group` which user and group can be passed literraly or
	// by name.
	User *string `pulumi:"user"`
	// Sets the usernamespace mode for the container when usernamespace remapping option is enabled.
	UsernsMode *string `pulumi:"usernsMode"`
	// See Volumes below for details.
	Volumes []ContainerVolume `pulumi:"volumes"`
	// The working directory for commands to run in
	WorkingDir *string `pulumi:"workingDir"`
}

type ContainerState struct {
	// If true attach to the container after its creation and waits the end of his execution.
	Attach pulumi.BoolPtrInput
	// The network bridge of the container as read from its NetworkSettings.
	Bridge pulumi.StringPtrInput
	// See Capabilities below for details.
	Capabilities ContainerCapabilitiesPtrInput
	// The command to use to start the
	// container. For example, to run `/usr/bin/myprogram -f baz.conf` set the
	// command to be `["/usr/bin/myprogram", "-f", "baz.conf"]`.
	Command pulumi.StringArrayInput
	// The logs of the container if its execution is done (`attach` must be disabled).
	ContainerLogs pulumi.StringPtrInput
	// A comma-separated list or hyphen-separated range of CPUs a container can use, e.g. `0-1`.
	CpuSet pulumi.StringPtrInput
	// CPU shares (relative weight) for the container.
	CpuShares pulumi.IntPtrInput
	// If defined will attempt to stop the container before destroying. Container will be destroyed after `n` seconds or on successful stop.
	DestroyGraceSeconds pulumi.IntPtrInput
	// See Devices below for details.
	Devices ContainerDeviceArrayInput
	// Set of DNS servers.
	Dns pulumi.StringArrayInput
	// Set of DNS options used by the DNS provider(s), see `resolv.conf` documentation for valid list of options.
	DnsOpts pulumi.StringArrayInput
	// Set of DNS search domains that are used when bare unqualified hostnames are used inside of the container.
	DnsSearches pulumi.StringArrayInput
	// Domain name of the container.
	Domainname pulumi.StringPtrInput
	// The command to use as the
	// Entrypoint for the container. The Entrypoint allows you to configure a
	// container to run as an executable. For example, to run `/usr/bin/myprogram`
	// when starting a container, set the entrypoint to be
	// `["/usr/bin/myprogram"]`.
	Entrypoints pulumi.StringArrayInput
	// Environment variables to set.
	Envs pulumi.StringArrayInput
	// The exit code of the container if its execution is done (`mustRun` must be disabled).
	ExitCode pulumi.IntPtrInput
	// *Deprecated:* Use `networkData` instead. The network gateway of the container as read from its
	// NetworkSettings.
	//
	// Deprecated: Use gateway from ip_adresses_data instead. This field exposes the data of the container's first network.
	Gateway pulumi.StringPtrInput
	// Add additional groups to run as.
	GroupAdds pulumi.StringArrayInput
	// See Healthcheck below for details.
	Healthcheck ContainerHealthcheckPtrInput
	// Hostname of the container.
	Hostname pulumi.StringPtrInput
	// Hostname to add.
	Hosts ContainerHostArrayInput
	// The ID of the image to back this container.
	// The easiest way to get this value is to use the `RemoteImage` resource
	// as is shown in the example above.
	Image pulumi.StringPtrInput
	// Configured whether an init process should be injected for this container. If unset this will default to the `dockerd` defaults.
	Init pulumi.BoolPtrInput
	// *Deprecated:* Use `networkData` instead. The IP address of the container's first network it.
	//
	// Deprecated: Use ip_adresses_data instead. This field exposes the data of the container's first network.
	IpAddress pulumi.StringPtrInput
	// *Deprecated:* Use `networkData` instead. The IP prefix length of the container as read from its
	// NetworkSettings.
	//
	// Deprecated: Use ip_prefix_length from ip_adresses_data instead. This field exposes the data of the container's first network.
	IpPrefixLength pulumi.IntPtrInput
	// IPC sharing mode for the container. Possible values are: `none`, `private`, `shareable`, `container:<name|id>` or `host`.
	IpcMode pulumi.StringPtrInput
	// Adding labels.
	Labels ContainerLabelArrayInput
	// Set of links for link based
	// connectivity between containers that are running on the same host.
	//
	// Deprecated: The --link flag is a legacy feature of Docker. It may eventually be removed.
	Links pulumi.StringArrayInput
	// The logging driver to use for the container.
	// Defaults to "json-file".
	LogDriver pulumi.StringPtrInput
	// Key/value pairs to use as options for
	// the logging driver.
	LogOpts pulumi.MapInput
	// Save the container logs (`attach` must be enabled).
	Logs pulumi.BoolPtrInput
	// The maximum amount of times to an attempt
	// a restart when `restart` is set to "on-failure"
	MaxRetryCount pulumi.IntPtrInput
	// The memory limit for the container in MBs.
	Memory     pulumi.IntPtrInput
	MemorySwap pulumi.IntPtrInput
	// See Mounts below for details.
	Mounts  ContainerMountArrayInput
	MustRun pulumi.BoolPtrInput
	Name    pulumi.StringPtrInput
	// Network aliases of the container for user-defined networks only. *Deprecated:* use `networksAdvanced` instead.
	//
	// Deprecated: Use networks_advanced instead. Will be removed in v2.0.0
	NetworkAliases pulumi.StringArrayInput
	// (Map of a block) The IP addresses of the container on each
	// network. Key are the network names, values are the IP addresses.
	NetworkDatas ContainerNetworkDataArrayInput
	// Network mode of the container.
	NetworkMode pulumi.StringPtrInput
	// Id of the networks in which the
	// container is. *Deprecated:* use `networksAdvanced` instead.
	//
	// Deprecated: Use networks_advanced instead. Will be removed in v2.0.0
	Networks pulumi.StringArrayInput
	// See Networks Advanced below for details. If this block has priority to the deprecated `networkAlias` and `network` properties.
	NetworksAdvanced ContainerNetworksAdvancedArrayInput
	// The PID (Process) Namespace mode for the container. Either `container:<name|id>` or `host`.
	PidMode pulumi.StringPtrInput
	// See Ports below for details.
	Ports ContainerPortArrayInput
	// Run container in privileged mode.
	Privileged pulumi.BoolPtrInput
	// Publish all ports of the container.
	PublishAllPorts pulumi.BoolPtrInput
	// If true, this volume will be readonly.
	// Defaults to false.
	ReadOnly      pulumi.BoolPtrInput
	RemoveVolumes pulumi.BoolPtrInput
	// The restart policy for the container. Must be
	// one of "no", "on-failure", "always", "unless-stopped".
	Restart pulumi.StringPtrInput
	Rm      pulumi.BoolPtrInput
	// Set of string values to customize labels for MLS systems, such as SELinux. See https://docs.docker.com/engine/reference/run/#security-configuration.
	SecurityOpts pulumi.StringArrayInput
	// Size of `/dev/shm` in MBs.
	ShmSize pulumi.IntPtrInput
	// If true, then the Docker container will be
	// started after creation. If false, then the container is only created.
	Start pulumi.BoolPtrInput
	// A map of kernel parameters (sysctls) to set in the container.
	Sysctls pulumi.MapInput
	// A map of container directories which should be replaced by `tmpfs mounts`, and their corresponding mount options.
	Tmpfs pulumi.MapInput
	// See Ulimits below for
	// details.
	Ulimits ContainerUlimitArrayInput
	// See File Upload below for details.
	Uploads ContainerUploadArrayInput
	// User used for run the first process. Format is
	// `user` or `user:group` which user and group can be passed literraly or
	// by name.
	User pulumi.StringPtrInput
	// Sets the usernamespace mode for the container when usernamespace remapping option is enabled.
	UsernsMode pulumi.StringPtrInput
	// See Volumes below for details.
	Volumes ContainerVolumeArrayInput
	// The working directory for commands to run in
	WorkingDir pulumi.StringPtrInput
}

func (ContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerState)(nil)).Elem()
}

type containerArgs struct {
	// If true attach to the container after its creation and waits the end of his execution.
	Attach *bool `pulumi:"attach"`
	// See Capabilities below for details.
	Capabilities *ContainerCapabilities `pulumi:"capabilities"`
	// The command to use to start the
	// container. For example, to run `/usr/bin/myprogram -f baz.conf` set the
	// command to be `["/usr/bin/myprogram", "-f", "baz.conf"]`.
	Command []string `pulumi:"command"`
	// A comma-separated list or hyphen-separated range of CPUs a container can use, e.g. `0-1`.
	CpuSet *string `pulumi:"cpuSet"`
	// CPU shares (relative weight) for the container.
	CpuShares *int `pulumi:"cpuShares"`
	// If defined will attempt to stop the container before destroying. Container will be destroyed after `n` seconds or on successful stop.
	DestroyGraceSeconds *int `pulumi:"destroyGraceSeconds"`
	// See Devices below for details.
	Devices []ContainerDevice `pulumi:"devices"`
	// Set of DNS servers.
	Dns []string `pulumi:"dns"`
	// Set of DNS options used by the DNS provider(s), see `resolv.conf` documentation for valid list of options.
	DnsOpts []string `pulumi:"dnsOpts"`
	// Set of DNS search domains that are used when bare unqualified hostnames are used inside of the container.
	DnsSearches []string `pulumi:"dnsSearches"`
	// Domain name of the container.
	Domainname *string `pulumi:"domainname"`
	// The command to use as the
	// Entrypoint for the container. The Entrypoint allows you to configure a
	// container to run as an executable. For example, to run `/usr/bin/myprogram`
	// when starting a container, set the entrypoint to be
	// `["/usr/bin/myprogram"]`.
	Entrypoints []string `pulumi:"entrypoints"`
	// Environment variables to set.
	Envs []string `pulumi:"envs"`
	// Add additional groups to run as.
	GroupAdds []string `pulumi:"groupAdds"`
	// See Healthcheck below for details.
	Healthcheck *ContainerHealthcheck `pulumi:"healthcheck"`
	// Hostname of the container.
	Hostname *string `pulumi:"hostname"`
	// Hostname to add.
	Hosts []ContainerHost `pulumi:"hosts"`
	// The ID of the image to back this container.
	// The easiest way to get this value is to use the `RemoteImage` resource
	// as is shown in the example above.
	Image string `pulumi:"image"`
	// Configured whether an init process should be injected for this container. If unset this will default to the `dockerd` defaults.
	Init *bool `pulumi:"init"`
	// IPC sharing mode for the container. Possible values are: `none`, `private`, `shareable`, `container:<name|id>` or `host`.
	IpcMode *string `pulumi:"ipcMode"`
	// Adding labels.
	Labels []ContainerLabel `pulumi:"labels"`
	// Set of links for link based
	// connectivity between containers that are running on the same host.
	//
	// Deprecated: The --link flag is a legacy feature of Docker. It may eventually be removed.
	Links []string `pulumi:"links"`
	// The logging driver to use for the container.
	// Defaults to "json-file".
	LogDriver *string `pulumi:"logDriver"`
	// Key/value pairs to use as options for
	// the logging driver.
	LogOpts map[string]interface{} `pulumi:"logOpts"`
	// Save the container logs (`attach` must be enabled).
	Logs *bool `pulumi:"logs"`
	// The maximum amount of times to an attempt
	// a restart when `restart` is set to "on-failure"
	MaxRetryCount *int `pulumi:"maxRetryCount"`
	// The memory limit for the container in MBs.
	Memory     *int `pulumi:"memory"`
	MemorySwap *int `pulumi:"memorySwap"`
	// See Mounts below for details.
	Mounts  []ContainerMount `pulumi:"mounts"`
	MustRun *bool            `pulumi:"mustRun"`
	Name    *string          `pulumi:"name"`
	// Network aliases of the container for user-defined networks only. *Deprecated:* use `networksAdvanced` instead.
	//
	// Deprecated: Use networks_advanced instead. Will be removed in v2.0.0
	NetworkAliases []string `pulumi:"networkAliases"`
	// Network mode of the container.
	NetworkMode *string `pulumi:"networkMode"`
	// Id of the networks in which the
	// container is. *Deprecated:* use `networksAdvanced` instead.
	//
	// Deprecated: Use networks_advanced instead. Will be removed in v2.0.0
	Networks []string `pulumi:"networks"`
	// See Networks Advanced below for details. If this block has priority to the deprecated `networkAlias` and `network` properties.
	NetworksAdvanced []ContainerNetworksAdvanced `pulumi:"networksAdvanced"`
	// The PID (Process) Namespace mode for the container. Either `container:<name|id>` or `host`.
	PidMode *string `pulumi:"pidMode"`
	// See Ports below for details.
	Ports []ContainerPort `pulumi:"ports"`
	// Run container in privileged mode.
	Privileged *bool `pulumi:"privileged"`
	// Publish all ports of the container.
	PublishAllPorts *bool `pulumi:"publishAllPorts"`
	// If true, this volume will be readonly.
	// Defaults to false.
	ReadOnly      *bool `pulumi:"readOnly"`
	RemoveVolumes *bool `pulumi:"removeVolumes"`
	// The restart policy for the container. Must be
	// one of "no", "on-failure", "always", "unless-stopped".
	Restart *string `pulumi:"restart"`
	Rm      *bool   `pulumi:"rm"`
	// Set of string values to customize labels for MLS systems, such as SELinux. See https://docs.docker.com/engine/reference/run/#security-configuration.
	SecurityOpts []string `pulumi:"securityOpts"`
	// Size of `/dev/shm` in MBs.
	ShmSize *int `pulumi:"shmSize"`
	// If true, then the Docker container will be
	// started after creation. If false, then the container is only created.
	Start *bool `pulumi:"start"`
	// A map of kernel parameters (sysctls) to set in the container.
	Sysctls map[string]interface{} `pulumi:"sysctls"`
	// A map of container directories which should be replaced by `tmpfs mounts`, and their corresponding mount options.
	Tmpfs map[string]interface{} `pulumi:"tmpfs"`
	// See Ulimits below for
	// details.
	Ulimits []ContainerUlimit `pulumi:"ulimits"`
	// See File Upload below for details.
	Uploads []ContainerUpload `pulumi:"uploads"`
	// User used for run the first process. Format is
	// `user` or `user:group` which user and group can be passed literraly or
	// by name.
	User *string `pulumi:"user"`
	// Sets the usernamespace mode for the container when usernamespace remapping option is enabled.
	UsernsMode *string `pulumi:"usernsMode"`
	// See Volumes below for details.
	Volumes []ContainerVolume `pulumi:"volumes"`
	// The working directory for commands to run in
	WorkingDir *string `pulumi:"workingDir"`
}

// The set of arguments for constructing a Container resource.
type ContainerArgs struct {
	// If true attach to the container after its creation and waits the end of his execution.
	Attach pulumi.BoolPtrInput
	// See Capabilities below for details.
	Capabilities ContainerCapabilitiesPtrInput
	// The command to use to start the
	// container. For example, to run `/usr/bin/myprogram -f baz.conf` set the
	// command to be `["/usr/bin/myprogram", "-f", "baz.conf"]`.
	Command pulumi.StringArrayInput
	// A comma-separated list or hyphen-separated range of CPUs a container can use, e.g. `0-1`.
	CpuSet pulumi.StringPtrInput
	// CPU shares (relative weight) for the container.
	CpuShares pulumi.IntPtrInput
	// If defined will attempt to stop the container before destroying. Container will be destroyed after `n` seconds or on successful stop.
	DestroyGraceSeconds pulumi.IntPtrInput
	// See Devices below for details.
	Devices ContainerDeviceArrayInput
	// Set of DNS servers.
	Dns pulumi.StringArrayInput
	// Set of DNS options used by the DNS provider(s), see `resolv.conf` documentation for valid list of options.
	DnsOpts pulumi.StringArrayInput
	// Set of DNS search domains that are used when bare unqualified hostnames are used inside of the container.
	DnsSearches pulumi.StringArrayInput
	// Domain name of the container.
	Domainname pulumi.StringPtrInput
	// The command to use as the
	// Entrypoint for the container. The Entrypoint allows you to configure a
	// container to run as an executable. For example, to run `/usr/bin/myprogram`
	// when starting a container, set the entrypoint to be
	// `["/usr/bin/myprogram"]`.
	Entrypoints pulumi.StringArrayInput
	// Environment variables to set.
	Envs pulumi.StringArrayInput
	// Add additional groups to run as.
	GroupAdds pulumi.StringArrayInput
	// See Healthcheck below for details.
	Healthcheck ContainerHealthcheckPtrInput
	// Hostname of the container.
	Hostname pulumi.StringPtrInput
	// Hostname to add.
	Hosts ContainerHostArrayInput
	// The ID of the image to back this container.
	// The easiest way to get this value is to use the `RemoteImage` resource
	// as is shown in the example above.
	Image pulumi.StringInput
	// Configured whether an init process should be injected for this container. If unset this will default to the `dockerd` defaults.
	Init pulumi.BoolPtrInput
	// IPC sharing mode for the container. Possible values are: `none`, `private`, `shareable`, `container:<name|id>` or `host`.
	IpcMode pulumi.StringPtrInput
	// Adding labels.
	Labels ContainerLabelArrayInput
	// Set of links for link based
	// connectivity between containers that are running on the same host.
	//
	// Deprecated: The --link flag is a legacy feature of Docker. It may eventually be removed.
	Links pulumi.StringArrayInput
	// The logging driver to use for the container.
	// Defaults to "json-file".
	LogDriver pulumi.StringPtrInput
	// Key/value pairs to use as options for
	// the logging driver.
	LogOpts pulumi.MapInput
	// Save the container logs (`attach` must be enabled).
	Logs pulumi.BoolPtrInput
	// The maximum amount of times to an attempt
	// a restart when `restart` is set to "on-failure"
	MaxRetryCount pulumi.IntPtrInput
	// The memory limit for the container in MBs.
	Memory     pulumi.IntPtrInput
	MemorySwap pulumi.IntPtrInput
	// See Mounts below for details.
	Mounts  ContainerMountArrayInput
	MustRun pulumi.BoolPtrInput
	Name    pulumi.StringPtrInput
	// Network aliases of the container for user-defined networks only. *Deprecated:* use `networksAdvanced` instead.
	//
	// Deprecated: Use networks_advanced instead. Will be removed in v2.0.0
	NetworkAliases pulumi.StringArrayInput
	// Network mode of the container.
	NetworkMode pulumi.StringPtrInput
	// Id of the networks in which the
	// container is. *Deprecated:* use `networksAdvanced` instead.
	//
	// Deprecated: Use networks_advanced instead. Will be removed in v2.0.0
	Networks pulumi.StringArrayInput
	// See Networks Advanced below for details. If this block has priority to the deprecated `networkAlias` and `network` properties.
	NetworksAdvanced ContainerNetworksAdvancedArrayInput
	// The PID (Process) Namespace mode for the container. Either `container:<name|id>` or `host`.
	PidMode pulumi.StringPtrInput
	// See Ports below for details.
	Ports ContainerPortArrayInput
	// Run container in privileged mode.
	Privileged pulumi.BoolPtrInput
	// Publish all ports of the container.
	PublishAllPorts pulumi.BoolPtrInput
	// If true, this volume will be readonly.
	// Defaults to false.
	ReadOnly      pulumi.BoolPtrInput
	RemoveVolumes pulumi.BoolPtrInput
	// The restart policy for the container. Must be
	// one of "no", "on-failure", "always", "unless-stopped".
	Restart pulumi.StringPtrInput
	Rm      pulumi.BoolPtrInput
	// Set of string values to customize labels for MLS systems, such as SELinux. See https://docs.docker.com/engine/reference/run/#security-configuration.
	SecurityOpts pulumi.StringArrayInput
	// Size of `/dev/shm` in MBs.
	ShmSize pulumi.IntPtrInput
	// If true, then the Docker container will be
	// started after creation. If false, then the container is only created.
	Start pulumi.BoolPtrInput
	// A map of kernel parameters (sysctls) to set in the container.
	Sysctls pulumi.MapInput
	// A map of container directories which should be replaced by `tmpfs mounts`, and their corresponding mount options.
	Tmpfs pulumi.MapInput
	// See Ulimits below for
	// details.
	Ulimits ContainerUlimitArrayInput
	// See File Upload below for details.
	Uploads ContainerUploadArrayInput
	// User used for run the first process. Format is
	// `user` or `user:group` which user and group can be passed literraly or
	// by name.
	User pulumi.StringPtrInput
	// Sets the usernamespace mode for the container when usernamespace remapping option is enabled.
	UsernsMode pulumi.StringPtrInput
	// See Volumes below for details.
	Volumes ContainerVolumeArrayInput
	// The working directory for commands to run in
	WorkingDir pulumi.StringPtrInput
}

func (ContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerArgs)(nil)).Elem()
}

type ContainerInput interface {
	pulumi.Input

	ToContainerOutput() ContainerOutput
	ToContainerOutputWithContext(ctx context.Context) ContainerOutput
}

func (Container) ElementType() reflect.Type {
	return reflect.TypeOf((*Container)(nil)).Elem()
}

func (i Container) ToContainerOutput() ContainerOutput {
	return i.ToContainerOutputWithContext(context.Background())
}

func (i Container) ToContainerOutputWithContext(ctx context.Context) ContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerOutput)
}

type ContainerOutput struct {
	*pulumi.OutputState
}

func (ContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerOutput)(nil)).Elem()
}

func (o ContainerOutput) ToContainerOutput() ContainerOutput {
	return o
}

func (o ContainerOutput) ToContainerOutputWithContext(ctx context.Context) ContainerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ContainerOutput{})
}
