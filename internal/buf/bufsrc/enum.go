// Copyright 2020 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bufsrc

type enum struct {
	namedDescriptor

	values         []EnumValue
	allowAlias     bool
	allowAliasPath []int32
	reservedRanges []ReservedRange
	reservedNames  []ReservedName
}

func newEnum(
	namedDescriptor namedDescriptor,
	allowAlias bool,
	allowAliasPath []int32,
) *enum {
	return &enum{
		namedDescriptor: namedDescriptor,
		allowAlias:      allowAlias,
		allowAliasPath:  allowAliasPath,
	}
}

func (e *enum) Values() []EnumValue {
	return e.values
}

func (e *enum) AllowAlias() bool {
	return e.allowAlias
}

func (e *enum) AllowAliasLocation() Location {
	return e.getLocation(e.allowAliasPath)
}

func (e *enum) ReservedRanges() []ReservedRange {
	return e.reservedRanges
}

func (e *enum) ReservedNames() []ReservedName {
	return e.reservedNames
}

func (e *enum) addValue(value EnumValue) {
	e.values = append(e.values, value)
}

func (e *enum) addReservedRange(reservedRange ReservedRange) {
	e.reservedRanges = append(e.reservedRanges, reservedRange)
}

func (e *enum) addReservedName(reservedName ReservedName) {
	e.reservedNames = append(e.reservedNames, reservedName)
}