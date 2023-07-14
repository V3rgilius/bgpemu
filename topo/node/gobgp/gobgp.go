// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package gobgp

import (
	"fmt"

	ktpb "github.com/openconfig/kne/proto/topo"
	node "github.com/v3rgilius/bgpemu/topo/node"
)

func New(nodeImpl *node.Impl) (node.Node, error) {
	if nodeImpl == nil {
		return nil, fmt.Errorf("nodeImpl cannot be nil")
	}
	if nodeImpl.Proto == nil {
		return nil, fmt.Errorf("nodeImpl.Proto cannot be nil")
	}
	defaults(nodeImpl.Proto)
	n := &Node{
		Impl: nodeImpl,
	}
	return n, nil
}

type Node struct {
	*node.Impl
}

func defaults(pb *ktpb.Node) *ktpb.Node {
	if pb.Config == nil {
		pb.Config = &ktpb.Config{}
	}
	if pb.Config.Image == "" {
		pb.Config.Image = "midnightreiter/gobgp:v3.11.0"
	}
	if len(pb.GetConfig().GetCommand()) == 0 {
		pb.Config.Command = []string{"/bin/sleep", "2000000"}
	}
	if pb.Config.EntryCommand == "" {
		pb.Config.EntryCommand = fmt.Sprintf("kubectl exec -it %s -- /bin/bash", pb.Name)
	}
	return pb
}

func init() {
	node.Vendor(ktpb.Vendor_GOBGP, New)
}
