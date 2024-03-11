// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mypkg

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"secrets/mypkg/internal"
)

type Resource struct {
	pulumi.CustomResourceState

	Config      ConfigOutput             `pulumi:"config"`
	ConfigArray ConfigArrayOutput        `pulumi:"configArray"`
	ConfigMap   ConfigMapOutput          `pulumi:"configMap"`
	Foo         pulumi.StringOutput      `pulumi:"foo"`
	FooArray    pulumi.StringArrayOutput `pulumi:"fooArray"`
	FooMap      pulumi.StringMapOutput   `pulumi:"fooMap"`
}

// NewResource registers a new resource with the given unique name, arguments, and options.
func NewResource(ctx *pulumi.Context,
	name string, args *ResourceArgs, opts ...pulumi.ResourceOption) (*Resource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Config == nil {
		return nil, errors.New("invalid value for required argument 'Config'")
	}
	if args.ConfigArray == nil {
		return nil, errors.New("invalid value for required argument 'ConfigArray'")
	}
	if args.ConfigMap == nil {
		return nil, errors.New("invalid value for required argument 'ConfigMap'")
	}
	if args.Foo == nil {
		return nil, errors.New("invalid value for required argument 'Foo'")
	}
	if args.FooArray == nil {
		return nil, errors.New("invalid value for required argument 'FooArray'")
	}
	if args.FooMap == nil {
		return nil, errors.New("invalid value for required argument 'FooMap'")
	}
	if args.Config != nil {
		args.Config = pulumi.ToSecret(args.Config).(ConfigInput)
	}
	if args.ConfigArray != nil {
		args.ConfigArray = pulumi.ToSecret(args.ConfigArray).(ConfigArrayInput)
	}
	if args.ConfigMap != nil {
		args.ConfigMap = pulumi.ToSecret(args.ConfigMap).(ConfigMapInput)
	}
	if args.Foo != nil {
		args.Foo = pulumi.ToSecret(args.Foo).(pulumi.StringInput)
	}
	if args.FooArray != nil {
		args.FooArray = pulumi.ToSecret(args.FooArray).(pulumi.StringArrayInput)
	}
	if args.FooMap != nil {
		args.FooMap = pulumi.ToSecret(args.FooMap).(pulumi.StringMapInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"config",
		"configArray",
		"configMap",
		"foo",
		"fooArray",
		"fooMap",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Resource
	err := ctx.RegisterResource("mypkg::Resource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResource gets an existing Resource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceState, opts ...pulumi.ResourceOption) (*Resource, error) {
	var resource Resource
	err := ctx.ReadResource("mypkg::Resource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Resource resources.
type resourceState struct {
}

type ResourceState struct {
}

func (ResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceState)(nil)).Elem()
}

type resourceArgs struct {
	Config      Config            `pulumi:"config"`
	ConfigArray []Config          `pulumi:"configArray"`
	ConfigMap   map[string]Config `pulumi:"configMap"`
	Foo         string            `pulumi:"foo"`
	FooArray    []string          `pulumi:"fooArray"`
	FooMap      map[string]string `pulumi:"fooMap"`
}

// The set of arguments for constructing a Resource resource.
type ResourceArgs struct {
	Config      ConfigInput
	ConfigArray ConfigArrayInput
	ConfigMap   ConfigMapInput
	Foo         pulumi.StringInput
	FooArray    pulumi.StringArrayInput
	FooMap      pulumi.StringMapInput
}

func (ResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceArgs)(nil)).Elem()
}

type ResourceInput interface {
	pulumi.Input

	ToResourceOutput() ResourceOutput
	ToResourceOutputWithContext(ctx context.Context) ResourceOutput
}

func (*Resource) ElementType() reflect.Type {
	return reflect.TypeOf((**Resource)(nil)).Elem()
}

func (i *Resource) ToResourceOutput() ResourceOutput {
	return i.ToResourceOutputWithContext(context.Background())
}

func (i *Resource) ToResourceOutputWithContext(ctx context.Context) ResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceOutput)
}

func (i *Resource) ToOutput(ctx context.Context) pulumix.Output[*Resource] {
	return pulumix.Output[*Resource]{
		OutputState: i.ToResourceOutputWithContext(ctx).OutputState,
	}
}

type ResourceOutput struct{ *pulumi.OutputState }

func (ResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Resource)(nil)).Elem()
}

func (o ResourceOutput) ToResourceOutput() ResourceOutput {
	return o
}

func (o ResourceOutput) ToResourceOutputWithContext(ctx context.Context) ResourceOutput {
	return o
}

func (o ResourceOutput) ToOutput(ctx context.Context) pulumix.Output[*Resource] {
	return pulumix.Output[*Resource]{
		OutputState: o.OutputState,
	}
}

func (o ResourceOutput) Config() ConfigOutput {
	return o.ApplyT(func(v *Resource) ConfigOutput { return v.Config }).(ConfigOutput)
}

func (o ResourceOutput) ConfigArray() ConfigArrayOutput {
	return o.ApplyT(func(v *Resource) ConfigArrayOutput { return v.ConfigArray }).(ConfigArrayOutput)
}

func (o ResourceOutput) ConfigMap() ConfigMapOutput {
	return o.ApplyT(func(v *Resource) ConfigMapOutput { return v.ConfigMap }).(ConfigMapOutput)
}

func (o ResourceOutput) Foo() pulumi.StringOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringOutput { return v.Foo }).(pulumi.StringOutput)
}

func (o ResourceOutput) FooArray() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringArrayOutput { return v.FooArray }).(pulumi.StringArrayOutput)
}

func (o ResourceOutput) FooMap() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringMapOutput { return v.FooMap }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceInput)(nil)).Elem(), &Resource{})
	pulumi.RegisterOutputType(ResourceOutput{})
}