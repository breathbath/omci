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

const Vdsl2LineConfigurationExtensions3ClassId uint16 = 410

var vdsl2lineconfigurationextensions3BME *BaseManagedEntityDefinition

// Vdsl2LineConfigurationExtensions3 (class ID #410) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type Vdsl2LineConfigurationExtensions3 struct {
	BaseManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	vdsl2lineconfigurationextensions3BME := &BaseManagedEntityDefinition{
		Name:     "Vdsl2LineConfigurationExtensions3",
		ClassID:  410,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFFF0,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, Read|SetByCreate, false, false, false),
			1:  ByteField("Ripolicyds", 0, Read|Write, false, false, true),
			2:  ByteField("Ripolicyus", 0, Read|Write, false, false, true),
			3:  ByteField("ReinitTimeThresholdds", 0, Read|Write, false, false, true),
			4:  ByteField("ReinitTimeThresholdus", 0, Read|Write, false, false, true),
			5:  ByteField("Rxrefvnsfus", 0, Read|Write, false, false, true),
			6:  ByteField("Txrefvnsfds", 0, Read|Write, false, false, true),
			7:  ByteField("RtxModeds", 0, Read|Write, false, false, false),
			8:  ByteField("RtxModeus", 0, Read|Write, false, false, false),
			9:  ByteField("LeftrThresh", 0, Read|Write, false, false, false),
			10: ByteField("MaxdelayoctetSplitParameterMdosplit", 0, Read|Write, false, false, true),
			11: ByteField("AttndrMethodAttndrMethod", 0, Read|Write, false, false, true),
			12: ByteField("AttndrMaxdelayoctetSplitParameterAttndrMdosplit", 0, Read|Write, false, false, true),
		},
	}
}

// NewVdsl2LineConfigurationExtensions3 (class ID 410 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewVdsl2LineConfigurationExtensions3(params ...ParamData) (IManagedEntity, error) {
	entity := &ManagedEntity {
	    Definition: vdsl2lineconfigurationextensions3BME,
	    Attributes: make(map[string]interface{}),
	}
	if err := entity.setAttributes(params...); err != nil {
	    return nil, err
	}
	return entity, nil
}
