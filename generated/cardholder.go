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

const CardholderClassId uint16 = 5

// Cardholder (class ID #5) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type Cardholder struct {
	BaseManagedEntityDefinition
}

// NewCardholder (class ID 5 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewCardholder(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "Cardholder",
		ClassID:  5,
		EntityID: eid,
		MessageTypes: []MsgType{
			Get,
			Set,
		},
		AllowedAttributeMask: 0,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, Read, false, false, false, false),
			1: ByteField("ActualPlugInUnitType", 0, Read, true, false, false, false),
			2: ByteField("ExpectedPlugInUnitType", 0, Read|Write, false, false, false, false),
			3: ByteField("ExpectedPortCount", 0, Read|Write, false, false, false, true),
			4: MultiByteField("ExpectedEquipmentId", 20, nil, Read|Write, false, false, false, true),
			5: MultiByteField("ActualEquipmentId", 20, nil, Read, true, false, false, true),
			6: ByteField("ProtectionProfilePointer", 0, Read, false, false, false, true),
			7: ByteField("InvokeProtectionSwitch", 0, Read|Write, false, false, false, true),
			8: ByteField("AlarmReportingControl", 0, Read|Write, true, false, false, true),
			9: ByteField("ArcInterval", 0, Read|Write, false, false, false, true),
		},
	}
	entity.computeAttributeMask()
	return &Cardholder{entity}, nil
}
