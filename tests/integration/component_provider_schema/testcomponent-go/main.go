// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.
//go:build !all
// +build !all

package main

import (
	"fmt"
	"os"

	"github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/cmdutil"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	pulumiprovider "github.com/pulumi/pulumi/sdk/v3/go/pulumi/provider"
)

const providerName = "testcomponent"
const version = "0.0.1"

func main() {
	var schema string
	if _, ok := os.LookupEnv("INCLUDE_SCHEMA"); ok {
		schema = `{"hello": "world"}`
	}
	err := provider.MainWithOptions(provider.Options{
		Name:    providerName,
		Version: version,
		Schema:  []byte(schema),
		Construct: func(ctx *pulumi.Context, typ, name string,
			inputs pulumiprovider.ConstructInputs, options pulumi.ResourceOption) (*pulumiprovider.ConstructResult, error) {
			return nil, fmt.Errorf("unknown resource type %s", typ)
		},
	})
	if err != nil {
		cmdutil.ExitError(err.Error())
	}
}
