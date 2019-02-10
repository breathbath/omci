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

const SnmpConfigurationDataClassId uint16 = 335

var snmpconfigurationdataBME *BaseManagedEntityDefinition

// SnmpConfigurationData (class ID #335) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type SnmpConfigurationData struct {
	BaseManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	snmpconfigurationdataBME = &BaseManagedEntityDefinition{
		Name:    "SnmpConfigurationData",
		ClassID: 335,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFF00,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, Read|SetByCreate, false, false, false),
			1: Uint16Field("SnmpVersion", 0, Read|SetByCreate|Write, false, false, false),
			2: Uint16Field("SnmpAgentAddress", 0, Read|SetByCreate|Write, false, false, false),
			3: Uint32Field("SnmpServerAddress", 0, Read|SetByCreate|Write, false, false, false),
			4: Uint16Field("SnmpServerPort", 0, Read|SetByCreate|Write, false, false, false),
			5: Uint16Field("SecurityNamePointer", 0, Read|SetByCreate|Write, false, false, false),
			6: Uint16Field("CommunityForRead", 0, Read|SetByCreate|Write, false, false, false),
			7: Uint16Field("CommunityForWrite", 0, Read|SetByCreate|Write, false, false, false),
			8: Uint16Field("SysNamePointer", 0, Read|SetByCreate|Write, false, false, false),
		},
	}
}

// NewSnmpConfigurationData (class ID 335 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewSnmpConfigurationData(params ...ParamData) (*ManagedEntity, error) {
	entity := &ManagedEntity{
		Definition: snmpconfigurationdataBME,
		Attributes: make(map[string]interface{}),
	}
	if err := entity.setAttributes(params...); err != nil {
		return nil, err
	}
	return entity, nil
}
