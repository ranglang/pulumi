// *** WARNING: this file was generated by the Lumi IDL Compiler (LUMIDL). ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
    "errors"

    pbempty "github.com/golang/protobuf/ptypes/empty"
    pbstruct "github.com/golang/protobuf/ptypes/struct"
    "golang.org/x/net/context"

    "github.com/pulumi/lumi/pkg/resource"
    "github.com/pulumi/lumi/pkg/tokens"
    "github.com/pulumi/lumi/pkg/util/contract"
    "github.com/pulumi/lumi/pkg/util/mapper"
    "github.com/pulumi/lumi/sdk/go/pkg/lumirpc"
)

/* RPC stubs for Subnet resource provider */

// SubnetToken is the type token corresponding to the Subnet package type.
const SubnetToken = tokens.Type("aws:ec2/subnet:Subnet")

// SubnetProviderOps is a pluggable interface for Subnet-related management functionality.
type SubnetProviderOps interface {
    Check(ctx context.Context, obj *Subnet) ([]mapper.FieldError, error)
    Create(ctx context.Context, obj *Subnet) (resource.ID, error)
    Get(ctx context.Context, id resource.ID) (*Subnet, error)
    InspectChange(ctx context.Context,
        id resource.ID, old *Subnet, new *Subnet, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id resource.ID, old *Subnet, new *Subnet, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id resource.ID) error
}

// SubnetProvider is a dynamic gRPC-based plugin for managing Subnet resources.
type SubnetProvider struct {
    ops SubnetProviderOps
}

// NewSubnetProvider allocates a resource provider that delegates to a ops instance.
func NewSubnetProvider(ops SubnetProviderOps) lumirpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &SubnetProvider{ops: ops}
}

func (p *SubnetProvider) Check(
    ctx context.Context, req *lumirpc.CheckRequest) (*lumirpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(SubnetToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr == nil || len(decerr.Failures()) == 0 {
        failures, err := p.ops.Check(ctx, obj)
        if err != nil {
            return nil, err
        }
        if len(failures) > 0 {
            decerr = mapper.NewDecodeErr(failures)
        }
    }
    return resource.NewCheckResponse(decerr), nil
}

func (p *SubnetProvider) Name(
    ctx context.Context, req *lumirpc.NameRequest) (*lumirpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(SubnetToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    if obj.Name == nil || *obj.Name == "" {
        if req.Unknowns[Subnet_Name] {
            return nil, errors.New("Name property cannot be computed from unknown outputs")
        }
        return nil, errors.New("Name property cannot be empty")
    }
    return &lumirpc.NameResponse{Name: *obj.Name}, nil
}

func (p *SubnetProvider) Create(
    ctx context.Context, req *lumirpc.CreateRequest) (*lumirpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(SubnetToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    id, err := p.ops.Create(ctx, obj)
    if err != nil {
        return nil, err
    }
    return &lumirpc.CreateResponse{Id: string(id)}, nil
}

func (p *SubnetProvider) Get(
    ctx context.Context, req *lumirpc.GetRequest) (*lumirpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(SubnetToken))
    id := resource.ID(req.GetId())
    obj, err := p.ops.Get(ctx, id)
    if err != nil {
        return nil, err
    }
    return &lumirpc.GetResponse{
        Properties: resource.MarshalProperties(
            nil, resource.NewPropertyMap(obj), resource.MarshalOptions{}),
    }, nil
}

func (p *SubnetProvider) InspectChange(
    ctx context.Context, req *lumirpc.InspectChangeRequest) (*lumirpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(SubnetToken))
    id := resource.ID(req.GetId())
    old, oldprops, decerr := p.Unmarshal(req.GetOlds())
    if decerr != nil {
        return nil, decerr
    }
    new, newprops, decerr := p.Unmarshal(req.GetNews())
    if decerr != nil {
        return nil, decerr
    }
    var replaces []string
    diff := oldprops.Diff(newprops)
    if diff != nil {
        if diff.Changed("name") {
            replaces = append(replaces, "name")
        }
        if diff.Changed("cidrBlock") {
            replaces = append(replaces, "cidrBlock")
        }
        if diff.Changed("vpc") {
            replaces = append(replaces, "vpc")
        }
        if diff.Changed("availabilityZone") {
            replaces = append(replaces, "availabilityZone")
        }
    }
    more, err := p.ops.InspectChange(ctx, id, old, new, diff)
    if err != nil {
        return nil, err
    }
    return &lumirpc.InspectChangeResponse{
        Replaces: append(replaces, more...),
    }, err
}

func (p *SubnetProvider) Update(
    ctx context.Context, req *lumirpc.UpdateRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(SubnetToken))
    id := resource.ID(req.GetId())
    old, oldprops, err := p.Unmarshal(req.GetOlds())
    if err != nil {
        return nil, err
    }
    new, newprops, err := p.Unmarshal(req.GetNews())
    if err != nil {
        return nil, err
    }
    diff := oldprops.Diff(newprops)
    if err := p.ops.Update(ctx, id, old, new, diff); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *SubnetProvider) Delete(
    ctx context.Context, req *lumirpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(SubnetToken))
    id := resource.ID(req.GetId())
    if err := p.ops.Delete(ctx, id); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *SubnetProvider) Unmarshal(
    v *pbstruct.Struct) (*Subnet, resource.PropertyMap, mapper.DecodeError) {
    var obj Subnet
    props := resource.UnmarshalProperties(nil, v, resource.MarshalOptions{RawResources: true})
    result := mapper.MapIU(props.Mappable(), &obj)
    return &obj, props, result
}

/* Marshalable Subnet structure(s) */

// Subnet is a marshalable representation of its corresponding IDL type.
type Subnet struct {
    Name *string `json:"name,omitempty"`
    CIDRBlock string `json:"cidrBlock"`
    VPC resource.ID `json:"vpc"`
    AvailabilityZone *string `json:"availabilityZone,omitempty"`
    MapPublicIpOnLaunch *bool `json:"mapPublicIpOnLaunch,omitempty"`
}

// Subnet's properties have constants to make dealing with diffs and property bags easier.
const (
    Subnet_Name = "name"
    Subnet_CIDRBlock = "cidrBlock"
    Subnet_VPC = "vpc"
    Subnet_AvailabilityZone = "availabilityZone"
    Subnet_MapPublicIpOnLaunch = "mapPublicIpOnLaunch"
)

