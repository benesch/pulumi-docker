// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class ContainerLabelArgs extends com.pulumi.resources.ResourceArgs {

    public static final ContainerLabelArgs Empty = new ContainerLabelArgs();

    @Import(name="label", required=true)
    private Output<String> label;

    public Output<String> label() {
        return this.label;
    }

    @Import(name="value", required=true)
    private Output<String> value;

    public Output<String> value() {
        return this.value;
    }

    private ContainerLabelArgs() {}

    private ContainerLabelArgs(ContainerLabelArgs $) {
        this.label = $.label;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ContainerLabelArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ContainerLabelArgs $;

        public Builder() {
            $ = new ContainerLabelArgs();
        }

        public Builder(ContainerLabelArgs defaults) {
            $ = new ContainerLabelArgs(Objects.requireNonNull(defaults));
        }

        public Builder label(Output<String> label) {
            $.label = label;
            return this;
        }

        public Builder label(String label) {
            return label(Output.of(label));
        }

        public Builder value(Output<String> value) {
            $.value = value;
            return this;
        }

        public Builder value(String value) {
            return value(Output.of(value));
        }

        public ContainerLabelArgs build() {
            $.label = Objects.requireNonNull($.label, "expected parameter 'label' to be non-null");
            $.value = Objects.requireNonNull($.value, "expected parameter 'value' to be non-null");
            return $;
        }
    }

}
