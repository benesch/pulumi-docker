// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Outputs
{

    [OutputType]
    public sealed class ServiceUpdateConfig
    {
        public readonly string? Delay;
        public readonly string? FailureAction;
        public readonly string? MaxFailureRatio;
        public readonly string? Monitor;
        public readonly string? Order;
        public readonly int? Parallelism;

        [OutputConstructor]
        private ServiceUpdateConfig(
            string? delay,

            string? failureAction,

            string? maxFailureRatio,

            string? monitor,

            string? order,

            int? parallelism)
        {
            Delay = delay;
            FailureAction = failureAction;
            MaxFailureRatio = maxFailureRatio;
            Monitor = monitor;
            Order = order;
            Parallelism = parallelism;
        }
    }
}
