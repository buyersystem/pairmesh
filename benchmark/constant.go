// Copyright 2022 PairMesh, Inc.
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

package benchmark

// ModeType as string alias, represents type of a mode
type ModeType string

// ModeType represents type of a mode
const (
	ModeTypeEcho  ModeType = "echo"
	ModeTypeRelay ModeType = "relay"
)

// Some handy constant variables for benchmark tooling
const (
	BufferSize     = 512
	MaxClientCount = 1000
	MaxPayload     = 9000
)
