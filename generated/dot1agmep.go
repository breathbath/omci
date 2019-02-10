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

const Dot1AgMepClassId uint16 = 302

var dot1agmepBME *BaseManagedEntityDefinition

// Dot1AgMep (class ID #302) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type Dot1AgMep struct {
	BaseManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	dot1agmepBME = &BaseManagedEntityDefinition{
		Name:    "Dot1AgMep",
		ClassID: 302,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFFFC,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, Read|SetByCreate, false, false, false),
			1:  Uint16Field("Layer2EntityPointer", 0, Read|SetByCreate|Write, false, false, false),
			2:  ByteField("Layer2Type", 0, Read|SetByCreate|Write, false, false, false),
			3:  Uint16Field("MaPointer", 0, Read|SetByCreate|Write, false, false, false),
			4:  Uint16Field("MepId", 0, Read|SetByCreate|Write, false, false, false),
			5:  ByteField("MepControl", 0, Read|SetByCreate|Write, false, false, false),
			6:  Uint16Field("PrimaryVlan", 0, Read|SetByCreate|Write, false, false, false),
			7:  ByteField("AdministrativeState", 0, Read|SetByCreate|Write, false, false, false),
			8:  ByteField("CcmAndLtmPriority", 0, Read|SetByCreate|Write, false, false, false),
			9:  Uint64Field("EgressIdentifier", 0, Read|SetByCreate|Write, false, false, false),
			10: MultiByteField("PeerMepIds", 24, nil, Read|Write, false, false, false),
			11: ByteField("EthAisControl", 0, Read|SetByCreate|Write, false, false, false),
			12: ByteField("FaultAlarmThreshold", 0, Read|SetByCreate|Write, false, false, true),
			13: Uint16Field("AlarmDeclarationSoakTime", 0, Read|Write, false, false, false),
			14: Uint16Field("AlarmClearSoakTime", 0, Read|Write, false, false, false),
		},
	}
}

// NewDot1AgMep (class ID 302 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewDot1AgMep(params ...ParamData) (*ManagedEntity, error) {
	entity := &ManagedEntity{
		Definition: dot1agmepBME,
		Attributes: make(map[string]interface{}),
	}
	if err := entity.setAttributes(params...); err != nil {
		return nil, err
	}
	return entity, nil
}
