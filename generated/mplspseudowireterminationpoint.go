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

const MplsPseudowireTerminationPointClassId uint16 = 333

// MplsPseudowireTerminationPoint (class ID #333) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type MplsPseudowireTerminationPoint struct {
	BaseManagedEntityDefinition
}

// NewMplsPseudowireTerminationPoint (class ID 333 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewMplsPseudowireTerminationPoint(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "MplsPseudowireTerminationPoint",
		ClassID:  333,
		EntityID: eid,
		MessageTypes: []MsgType{
			Create,
			Delete,
			Get,
			Set,
		},
		AllowedAttributeMask: 0,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, Read|SetByCreate, false, false, false, false),
			1:  ByteField("TpType", 0, Read|SetByCreate|Write, false, false, false, false),
			2:  Uint16Field("TpPointer", 0, Read|SetByCreate|Write, false, false, false, false),
			3:  ByteField("MplsLabelIndicator", 0, Read|SetByCreate|Write, false, false, false, false),
			4:  ByteField("MplsPwDirection", 0, Read|SetByCreate|Write, false, false, false, false),
			5:  Uint32Field("MplsPwUplinkLabel", 0, Read|SetByCreate|Write, false, false, false, false),
			6:  Uint32Field("MplsPwDownlinkLabel", 0, Read|SetByCreate|Write, false, false, false, false),
			7:  ByteField("MplsPwTc", 0, Read|SetByCreate|Write, false, false, false, false),
			8:  ByteField("MplsTunnelDirection", 0, Read|SetByCreate|Write, false, false, false, false),
			9:  Uint32Field("MplsTunnelUplinkLabel", 0, Read|SetByCreate|Write, false, false, false, false),
			10: Uint32Field("MplsTunnelDownlinkLabel", 0, Read|SetByCreate|Write, false, false, false, false),
			11: ByteField("MplsTunnelTc", 0, Read|SetByCreate|Write, false, false, false, false),
			12: Uint16Field("PseudowireType", 0, Read|SetByCreate|Write, false, false, false, false),
			13: ByteField("PseudowireControlWordPreference", 0, Read|SetByCreate|Write, false, false, false, true),
			14: ByteField("AdministrativeState", 0, Read|Write, false, false, false, true),
			15: ByteField("OperationalState", 0, Read, true, false, false, true),
		},
	}
	entity.computeAttributeMask()
	return &MplsPseudowireTerminationPoint{entity}, nil
}
