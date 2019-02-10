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

const ThresholdData2ClassId uint16 = 274

var thresholddata2BME *BaseManagedEntityDefinition

// ThresholdData2 (class ID #274) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type ThresholdData2 struct {
	BaseManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	thresholddata2BME = &BaseManagedEntityDefinition{
		Name:    "ThresholdData2",
		ClassID: 274,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFE00,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, Read|SetByCreate, false, false, false),
			1: Uint32Field("ThresholdValue8", 0, Read|SetByCreate|Write, false, false, false),
			2: Uint32Field("ThresholdValue9", 0, Read|SetByCreate|Write, false, false, false),
			3: Uint32Field("ThresholdValue10", 0, Read|SetByCreate|Write, false, false, false),
			4: Uint32Field("ThresholdValue11", 0, Read|SetByCreate|Write, false, false, false),
			5: Uint32Field("ThresholdValue12", 0, Read|SetByCreate|Write, false, false, false),
			6: Uint32Field("ThresholdValue13", 0, Read|SetByCreate|Write, false, false, false),
			7: Uint32Field("ThresholdValue14", 0, Read|SetByCreate|Write, false, false, false),
		},
	}
}

// NewThresholdData2 (class ID 274 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewThresholdData2(params ...ParamData) (*ManagedEntity, error) {
	entity := &ManagedEntity{
		Definition: thresholddata2BME,
		Attributes: make(map[string]interface{}),
	}
	if err := entity.setAttributes(params...); err != nil {
		return nil, err
	}
	return entity, nil
}
