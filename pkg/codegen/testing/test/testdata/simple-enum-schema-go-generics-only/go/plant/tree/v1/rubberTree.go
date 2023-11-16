// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"simple-enum-schema-go-generics-only/plant"
	"simple-enum-schema-go-generics-only/plant/internal"
)

type RubberTree struct {
	pulumi.CustomResourceState

	Container pulumix.GPtrOutput[plant.Container, plant.ContainerOutput] `pulumi:"container"`
	Diameter  pulumix.Output[Diameter]                                   `pulumi:"diameter"`
	Farm      pulumix.Output[*string]                                    `pulumi:"farm"`
	Size      pulumix.Output[*TreeSize]                                  `pulumi:"size"`
	Type      pulumix.Output[RubberTreeVariety]                          `pulumi:"type"`
}

// NewRubberTree registers a new resource with the given unique name, arguments, and options.
func NewRubberTree(ctx *pulumi.Context,
	name string, args *RubberTreeArgs, opts ...pulumi.ResourceOption) (*RubberTree, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Container != nil {
		args.Container = pulumix.Apply(args.Container, func(o *plant.ContainerArgs) *plant.ContainerArgs { return o.Defaults() })
	}
	if args.Diameter == nil {
		args.Diameter = pulumix.Val(Diameter(6.0))
	}
	if args.Farm == nil {
		args.Farm = pulumix.Ptr("(unknown)")
	}
	if args.Size == nil {
		args.Size = pulumix.Ptr(TreeSize("medium"))
	}
	if args.Type == nil {
		args.Type = pulumix.Val(RubberTreeVariety("Burgundy"))
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RubberTree
	err := ctx.RegisterResource("plant:tree/v1:RubberTree", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRubberTree gets an existing RubberTree resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRubberTree(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RubberTreeState, opts ...pulumi.ResourceOption) (*RubberTree, error) {
	var resource RubberTree
	err := ctx.ReadResource("plant:tree/v1:RubberTree", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RubberTree resources.
type rubberTreeState struct {
	Farm *string `pulumi:"farm"`
}

type RubberTreeState struct {
	Farm pulumix.Input[*string]
}

func (RubberTreeState) ElementType() reflect.Type {
	return reflect.TypeOf((*rubberTreeState)(nil)).Elem()
}

type rubberTreeArgs struct {
	Container *plant.Container  `pulumi:"container"`
	Diameter  Diameter          `pulumi:"diameter"`
	Farm      *string           `pulumi:"farm"`
	Size      *TreeSize         `pulumi:"size"`
	Type      RubberTreeVariety `pulumi:"type"`
}

// The set of arguments for constructing a RubberTree resource.
type RubberTreeArgs struct {
	Container pulumix.Input[*plant.ContainerArgs]
	Diameter  pulumix.Input[Diameter]
	Farm      pulumix.Input[*string]
	Size      pulumix.Input[*TreeSize]
	Type      pulumix.Input[RubberTreeVariety]
}

func (RubberTreeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rubberTreeArgs)(nil)).Elem()
}

type RubberTreeOutput struct{ *pulumi.OutputState }

func (RubberTreeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RubberTree)(nil)).Elem()
}

func (o RubberTreeOutput) ToRubberTreeOutput() RubberTreeOutput {
	return o
}

func (o RubberTreeOutput) ToRubberTreeOutputWithContext(ctx context.Context) RubberTreeOutput {
	return o
}

func (o RubberTreeOutput) ToOutput(ctx context.Context) pulumix.Output[RubberTree] {
	return pulumix.Output[RubberTree]{
		OutputState: o.OutputState,
	}
}

func (o RubberTreeOutput) Container() pulumix.GPtrOutput[plant.Container, plant.ContainerOutput] {
	value := pulumix.Apply[RubberTree](o, func(v RubberTree) pulumix.GPtrOutput[plant.Container, plant.ContainerOutput] { return v.Container })
	unwrapped := pulumix.Flatten[*plant.Container, pulumix.GPtrOutput[plant.Container, plant.ContainerOutput]](value)
	return pulumix.GPtrOutput[plant.Container, plant.ContainerOutput]{OutputState: unwrapped.OutputState}
}

func (o RubberTreeOutput) Diameter() pulumix.Output[Diameter] {
	value := pulumix.Apply[RubberTree](o, func(v RubberTree) pulumix.Output[Diameter] { return v.Diameter })
	return pulumix.Flatten[Diameter, pulumix.Output[Diameter]](value)
}

func (o RubberTreeOutput) Farm() pulumix.Output[*string] {
	value := pulumix.Apply[RubberTree](o, func(v RubberTree) pulumix.Output[*string] { return v.Farm })
	return pulumix.Flatten[*string, pulumix.Output[*string]](value)
}

func (o RubberTreeOutput) Size() pulumix.Output[*TreeSize] {
	value := pulumix.Apply[RubberTree](o, func(v RubberTree) pulumix.Output[*TreeSize] { return v.Size })
	return pulumix.Flatten[*TreeSize, pulumix.Output[*TreeSize]](value)
}

func (o RubberTreeOutput) Type() pulumix.Output[RubberTreeVariety] {
	value := pulumix.Apply[RubberTree](o, func(v RubberTree) pulumix.Output[RubberTreeVariety] { return v.Type })
	return pulumix.Flatten[RubberTreeVariety, pulumix.Output[RubberTreeVariety]](value)
}

func init() {
	pulumi.RegisterOutputType(RubberTreeOutput{})
}
