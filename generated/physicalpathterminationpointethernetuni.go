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

const PhysicalPathTerminationPointEthernetUniClassId ClassID = ClassID(11)

var physicalpathterminationpointethernetuniBME *ManagedEntityDefinition

// PhysicalPathTerminationPointEthernetUni (class ID #11) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type PhysicalPathTerminationPointEthernetUni struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	physicalpathterminationpointethernetuniBME = &ManagedEntityDefinition{
		Name:    "PhysicalPathTerminationPointEthernetUni",
		ClassID: 11,
		MessageTypes: mapset.NewSetWith(
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFFFE,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read), false, false, false, false, 0),
			1:  ByteField("ExpectedType", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 1),
			2:  ByteField("SensedType", 0, mapset.NewSetWith(Read), true, false, false, false, 2),
			3:  ByteField("AutoDetectionConfiguration", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 3),
			4:  ByteField("EthernetLoopbackConfiguration", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 4),
			5:  ByteField("AdministrativeState", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 5),
			6:  ByteField("OperationalState", 0, mapset.NewSetWith(Read), true, false, true, false, 6),
			7:  ByteField("ConfigurationInd", 0, mapset.NewSetWith(Read), false, false, false, false, 7),
			8:  Uint16Field("MaxFrameSize", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 8),
			9:  ByteField("DteOrDceInd", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 9),
			10: Uint16Field("PauseTime", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 10),
			11: ByteField("BridgedOrIpInd", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 11),
			12: ByteField("Arc", 0, mapset.NewSetWith(Read, Write), true, false, true, false, 12),
			13: ByteField("ArcInterval", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 13),
			14: ByteField("PppoeFilter", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 14),
			15: ByteField("PowerControl", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 15),
		},
	}
}

// NewPhysicalPathTerminationPointEthernetUni (class ID 11 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewPhysicalPathTerminationPointEthernetUni(params ...ParamData) (*ManagedEntity, error) {
	return NewManagedEntity(physicalpathterminationpointethernetuniBME, params...)
}
