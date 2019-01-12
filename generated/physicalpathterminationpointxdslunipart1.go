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

const PhysicalPathTerminationPointXdslUniPart1ClassId uint16 = 98

// PhysicalPathTerminationPointXdslUniPart1 (class ID #98) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type PhysicalPathTerminationPointXdslUniPart1 struct {
	BaseManagedEntityDefinition
}

// NewPhysicalPathTerminationPointXdslUniPart1 (class ID 98 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewPhysicalPathTerminationPointXdslUniPart1(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "PhysicalPathTerminationPointXdslUniPart1",
		ClassID:  98,
		EntityID: eid,
		MessageTypes: []MsgType{
			Get,
			Set,
		},
		AllowedAttributeMask: 0,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, Read, false, false, false, false),
			1:  ByteField("LoopbackConfiguration", 0, Read|Write, false, false, false, false),
			2:  ByteField("AdministrativeState", 0, Read|Write, false, false, false, false),
			3:  ByteField("OperationalState", 0, Read, true, false, false, true),
			4:  Uint16Field("XdslLineConfigurationProfile", 0, Read|Write, false, false, false, false),
			5:  Uint16Field("XdslSubcarrierMaskingDownstreamProfile", 0, Read|Write, false, false, false, false),
			6:  Uint16Field("XdslSubcarrierMaskingUpstreamProfile", 0, Read|Write, false, false, false, false),
			7:  Uint16Field("XdslDownstreamPowerSpectralDensityPsdMaskProfile", 0, Read|Write, false, false, false, false),
			8:  Uint16Field("XdslDownstreamRfiBandsProfile", 0, Read|Write, false, false, false, false),
			9:  ByteField("Arc", 0, Read|Write, true, false, false, true),
			10: ByteField("ArcInterval", 0, Read|Write, false, false, false, true),
			11: ByteField("ModemType", 0, Read|Write, false, false, false, true),
			12: Uint16Field("UpstreamPsdMaskProfile", 0, Read|Write, false, false, false, true),
			13: Uint16Field("NetworkSpecificExtensionsPointer", 0, Read|Write, false, false, false, true),
		},
	}
	entity.computeAttributeMask()
	return &PhysicalPathTerminationPointXdslUniPart1{entity}, nil
}
