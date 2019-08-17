/*
 * Copyright (c) 2018 - present.  Boling Consulting Solutions (bcsw.net)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/*
 * NOTE: This file was generated, manual edits will be overwritten!
 *
 * Generated by 'goCodeGenerator.py':
 *              https://github.com/cboling/OMCI-parser/README.md
 */
package generated

import "github.com/deckarep/golang-set"

const InterworkingVccTerminationPointClassId ClassID = ClassID(14)

var interworkingvccterminationpointBME *ManagedEntityDefinition

// InterworkingVccTerminationPoint (class ID #14) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type InterworkingVccTerminationPoint struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	interworkingvccterminationpointBME = &ManagedEntityDefinition{
		Name:    "InterworkingVccTerminationPoint",
		ClassID: 14,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFF80,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, false, 0),
			1: Uint16Field("VciValue", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 1),
			2: Uint16Field("VpNetworkCtpConnectivityPointer", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 2),
			3: ByteField("Deprecated1", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, true, 3),
			4: Uint16Field("Deprecated2", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, true, 4),
			5: Uint16Field("Aal5ProfilePointer", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 5),
			6: Uint16Field("Deprecated3", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, true, 6),
			7: ByteField("AalLoopbackConfiguration", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 7),
			8: ByteField("PptpCounter", 0, mapset.NewSetWith(Read), false, false, true, false, 8),
			9: ByteField("OperationalState", 0, mapset.NewSetWith(Read), true, false, true, false, 9),
		},
	}
}

// NewInterworkingVccTerminationPoint (class ID 14 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewInterworkingVccTerminationPoint(params ...ParamData) (*ManagedEntity, error) {
	return NewManagedEntity(interworkingvccterminationpointBME, params...)
}
