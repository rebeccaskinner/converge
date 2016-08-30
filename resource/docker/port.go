// Copyright © 2016 Asteris, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docker

import (
	"fmt"

	dc "github.com/fsouza/go-dockerclient"
)

// Port represents a docker port specification
type Port struct {
	p dc.Port
}

// NewPort creates a new Port
func NewPort(s string) Port {
	return Port{p: dc.Port(s)}
}

// Port returns the port as a string
func (p Port) Port() string {
	return p.p.Port()
}

// Proto returns the protocol
func (p Port) Proto() string {
	return p.p.Proto()
}

// String returns the port as a string with Port and Protocol
func (p Port) String() string {
	return fmt.Sprintf("%s/%s", p.Port(), p.Proto())
}

// ToDockerClientPort returns a dc.Port from a Port
func (p Port) ToDockerClientPort() dc.Port {
	return dc.Port(p.String())
}
