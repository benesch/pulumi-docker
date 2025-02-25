// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker
{
    /// <summary>
    /// &lt;!-- Bug: Type and Name are switched --&gt;
    /// Pulls a Docker image to a given Docker host from a Docker Registry.
    ///  This resource will *not* pull new layers of the image automatically unless used in conjunction with docker.RegistryImage data source to update the `pull_triggers` field.
    /// 
    /// ## Example Usage
    /// ### Basic
    /// 
    /// Finds and downloads the latest `ubuntu:precise` image but does not check
    /// for further updates of the image
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Docker = Pulumi.Docker;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var ubuntu = new Docker.RemoteImage("ubuntu", new()
    ///     {
    ///         Name = "ubuntu:precise",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Dynamic updates
    /// 
    /// To be able to update an image dynamically when the `sha256` sum changes,
    /// you need to use it in combination with `docker.RegistryImage` as follows:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Docker = Pulumi.Docker;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var ubuntuRegistryImage = Docker.GetRegistryImage.Invoke(new()
    ///     {
    ///         Name = "ubuntu:precise",
    ///     });
    /// 
    ///     var ubuntuRemoteImage = new Docker.RemoteImage("ubuntuRemoteImage", new()
    ///     {
    ///         Name = ubuntuRegistryImage.Apply(getRegistryImageResult =&gt; getRegistryImageResult.Name),
    ///         PullTriggers = new[]
    ///         {
    ///             ubuntuRegistryImage.Apply(getRegistryImageResult =&gt; getRegistryImageResult.Sha256Digest),
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [DockerResourceType("docker:index/remoteImage:RemoteImage")]
    public partial class RemoteImage : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Configuration to build an image. Please see [docker build command reference](https://docs.docker.com/engine/reference/commandline/build/#options) too.
        /// </summary>
        [Output("build")]
        public Output<Outputs.RemoteImageBuild?> Build { get; private set; } = null!;

        /// <summary>
        /// If true, then the image is removed forcibly when the resource is destroyed.
        /// </summary>
        [Output("forceRemove")]
        public Output<bool?> ForceRemove { get; private set; } = null!;

        /// <summary>
        /// The ID of the image (as seen when executing `docker inspect` on the image). Can be used to reference the image via its ID in other resources.
        /// </summary>
        [Output("imageId")]
        public Output<string> ImageId { get; private set; } = null!;

        /// <summary>
        /// If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker local storage on destroy operation.
        /// </summary>
        [Output("keepLocally")]
        public Output<bool?> KeepLocally { get; private set; } = null!;

        /// <summary>
        /// The ID of the image in the form of `sha256:&lt;hash&gt;` image digest. Do not confuse it with the default `latest` tag.
        /// </summary>
        [Output("latest")]
        public Output<string> Latest { get; private set; } = null!;

        /// <summary>
        /// The name of the Docker image, including any tags or SHA256 repo digests.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("output")]
        public Output<string> Output { get; private set; } = null!;

        /// <summary>
        /// A value which cause an image pull when changed
        /// </summary>
        [Output("pullTrigger")]
        public Output<string?> PullTrigger { get; private set; } = null!;

        /// <summary>
        /// List of values which cause an image pull when changed. This is used to store the image digest from the registry when using the docker*registry*image.
        /// </summary>
        [Output("pullTriggers")]
        public Output<ImmutableArray<string>> PullTriggers { get; private set; } = null!;

        /// <summary>
        /// The image sha256 digest in the form of `repo[:tag]@sha256:&lt;hash&gt;`.
        /// </summary>
        [Output("repoDigest")]
        public Output<string> RepoDigest { get; private set; } = null!;

        /// <summary>
        /// A map of arbitrary strings that, when changed, will force the `docker.RemoteImage` resource to be replaced. This can be used to rebuild an image when contents of source code folders change
        /// </summary>
        [Output("triggers")]
        public Output<ImmutableDictionary<string, object>?> Triggers { get; private set; } = null!;


        /// <summary>
        /// Create a RemoteImage resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RemoteImage(string name, RemoteImageArgs args, CustomResourceOptions? options = null)
            : base("docker:index/remoteImage:RemoteImage", name, args ?? new RemoteImageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RemoteImage(string name, Input<string> id, RemoteImageState? state = null, CustomResourceOptions? options = null)
            : base("docker:index/remoteImage:RemoteImage", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RemoteImage resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RemoteImage Get(string name, Input<string> id, RemoteImageState? state = null, CustomResourceOptions? options = null)
        {
            return new RemoteImage(name, id, state, options);
        }
    }

    public sealed class RemoteImageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration to build an image. Please see [docker build command reference](https://docs.docker.com/engine/reference/commandline/build/#options) too.
        /// </summary>
        [Input("build")]
        public Input<Inputs.RemoteImageBuildArgs>? Build { get; set; }

        /// <summary>
        /// If true, then the image is removed forcibly when the resource is destroyed.
        /// </summary>
        [Input("forceRemove")]
        public Input<bool>? ForceRemove { get; set; }

        /// <summary>
        /// If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker local storage on destroy operation.
        /// </summary>
        [Input("keepLocally")]
        public Input<bool>? KeepLocally { get; set; }

        /// <summary>
        /// The name of the Docker image, including any tags or SHA256 repo digests.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// A value which cause an image pull when changed
        /// </summary>
        [Input("pullTrigger")]
        public Input<string>? PullTrigger { get; set; }

        [Input("pullTriggers")]
        private InputList<string>? _pullTriggers;

        /// <summary>
        /// List of values which cause an image pull when changed. This is used to store the image digest from the registry when using the docker*registry*image.
        /// </summary>
        public InputList<string> PullTriggers
        {
            get => _pullTriggers ?? (_pullTriggers = new InputList<string>());
            set => _pullTriggers = value;
        }

        [Input("triggers")]
        private InputMap<object>? _triggers;

        /// <summary>
        /// A map of arbitrary strings that, when changed, will force the `docker.RemoteImage` resource to be replaced. This can be used to rebuild an image when contents of source code folders change
        /// </summary>
        public InputMap<object> Triggers
        {
            get => _triggers ?? (_triggers = new InputMap<object>());
            set => _triggers = value;
        }

        public RemoteImageArgs()
        {
        }
        public static new RemoteImageArgs Empty => new RemoteImageArgs();
    }

    public sealed class RemoteImageState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration to build an image. Please see [docker build command reference](https://docs.docker.com/engine/reference/commandline/build/#options) too.
        /// </summary>
        [Input("build")]
        public Input<Inputs.RemoteImageBuildGetArgs>? Build { get; set; }

        /// <summary>
        /// If true, then the image is removed forcibly when the resource is destroyed.
        /// </summary>
        [Input("forceRemove")]
        public Input<bool>? ForceRemove { get; set; }

        /// <summary>
        /// The ID of the image (as seen when executing `docker inspect` on the image). Can be used to reference the image via its ID in other resources.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker local storage on destroy operation.
        /// </summary>
        [Input("keepLocally")]
        public Input<bool>? KeepLocally { get; set; }

        /// <summary>
        /// The ID of the image in the form of `sha256:&lt;hash&gt;` image digest. Do not confuse it with the default `latest` tag.
        /// </summary>
        [Input("latest")]
        public Input<string>? Latest { get; set; }

        /// <summary>
        /// The name of the Docker image, including any tags or SHA256 repo digests.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("output")]
        public Input<string>? Output { get; set; }

        /// <summary>
        /// A value which cause an image pull when changed
        /// </summary>
        [Input("pullTrigger")]
        public Input<string>? PullTrigger { get; set; }

        [Input("pullTriggers")]
        private InputList<string>? _pullTriggers;

        /// <summary>
        /// List of values which cause an image pull when changed. This is used to store the image digest from the registry when using the docker*registry*image.
        /// </summary>
        public InputList<string> PullTriggers
        {
            get => _pullTriggers ?? (_pullTriggers = new InputList<string>());
            set => _pullTriggers = value;
        }

        /// <summary>
        /// The image sha256 digest in the form of `repo[:tag]@sha256:&lt;hash&gt;`.
        /// </summary>
        [Input("repoDigest")]
        public Input<string>? RepoDigest { get; set; }

        [Input("triggers")]
        private InputMap<object>? _triggers;

        /// <summary>
        /// A map of arbitrary strings that, when changed, will force the `docker.RemoteImage` resource to be replaced. This can be used to rebuild an image when contents of source code folders change
        /// </summary>
        public InputMap<object> Triggers
        {
            get => _triggers ?? (_triggers = new InputMap<object>());
            set => _triggers = value;
        }

        public RemoteImageState()
        {
        }
        public static new RemoteImageState Empty => new RemoteImageState();
    }
}
