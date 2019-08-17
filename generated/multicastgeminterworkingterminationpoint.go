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

const MulticastGemInterworkingTerminationPointClassId ClassID = ClassID(281)

var multicastgeminterworkingterminationpointBME *ManagedEntityDefinition

// MulticastGemInterworkingTerminationPoint (class ID #281) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type MulticastGemInterworkingTerminationPoint struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	multicastgeminterworkingterminationpointBME = &ManagedEntityDefinition{
		Name:    "MulticastGemInterworkingTerminationPoint",
		ClassID: 281,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			GetNext,
			Set,
		),
		AllowedAttributeMask: 0XFE00,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, false, 0),
			1: Uint16Field("GemPortNetworkCtpConnectivityPointer", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 1),
			2: ByteField("InterworkingOption", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 2),
			3: Uint16Field("ServiceProfilePointer", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 3),
			4: ByteField("PptpCounter", 0, mapset.NewSetWith(Read), false, false, true, false, 4),
			5: ByteField("OperationalState", 0, mapset.NewSetWith(Read), true, false, true, false, 5),
			6: Uint16Field("GalProfilePointer", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 6),
			7: TableField("Ipv6MulticastAddressTable", TableInfo{nil, 24}, mapset.NewSetWith(Read, Write), false, true, false, 7),
		},
	}
}

// NewMulticastGemInterworkingTerminationPoint (class ID 281 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewMulticastGemInterworkingTerminationPoint(params ...ParamData) (*ManagedEntity, error) {
	return NewManagedEntity(multicastgeminterworkingterminationpointBME, params...)
}
