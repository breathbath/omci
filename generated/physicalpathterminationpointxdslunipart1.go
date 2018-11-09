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
 *              https://github.com/cboling/OMCI-parser
 */
package generated

// PhysicalPathTerminationPointXdslUniPart1 (class ID 98 defines the basic
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
			Set,
			Get,
		},
		AttributeMask: 0,
		Attributes: []*AttributeDefinition{
			Uint16Field("ManagedEntityId", 0, Read),
			ByteField("LoopbackConfiguration", 0, Read|Write),
			ByteField("AdministrativeState", 0, Read|Write),
			ByteField("OperationalState", 0, Read),
			Uint16Field("XdslLineConfigurationProfile", 0, Read|Write),
			Uint16Field("XdslSubcarrierMaskingDownstreamProfile", 0, Read|Write),
			Uint16Field("XdslSubcarrierMaskingUpstreamProfile", 0, Read|Write),
			Uint16Field("XdslDownstreamPowerSpectralDensityPsdMaskProfile", 0, Read|Write),
			Uint16Field("XdslDownstreamRfiBandsProfile", 0, Read|Write),
			ByteField("Arc", 0, Read|Write),
			ByteField("ArcInterval", 0, Read|Write),
			ByteField("ModemType", 0, Read|Write),
			Uint16Field("UpstreamPsdMaskProfile", 0, Read|Write),
			Uint16Field("NetworkSpecificExtensionsPointer", 0, Read|Write),
		},
	}
	entity.computeAttributeMask()
	return &PhysicalPathTerminationPointXdslUniPart1{entity}, nil
}
