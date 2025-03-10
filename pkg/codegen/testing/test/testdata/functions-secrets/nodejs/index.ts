// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { FuncWithSecretsArgs, FuncWithSecretsResult, FuncWithSecretsOutputArgs } from "./funcWithSecrets";
export const funcWithSecrets: typeof import("./funcWithSecrets").funcWithSecrets = null as any;
export const funcWithSecretsOutput: typeof import("./funcWithSecrets").funcWithSecretsOutput = null as any;

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;

utilities.lazyLoad(exports, ["funcWithSecrets","funcWithSecretsOutput"], () => require("./funcWithSecrets"));
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));
pulumi.runtime.registerResourcePackage("mypkg", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:mypkg") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
