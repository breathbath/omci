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

const EnhancedSecurityControlClassId uint16 = 332

var enhancedsecuritycontrolBME *BaseManagedEntityDefinition

// EnhancedSecurityControl (class ID #332) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type EnhancedSecurityControl struct {
	BaseManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	enhancedsecuritycontrolBME := &BaseManagedEntityDefinition{
		Name:     "EnhancedSecurityControl",
		ClassID:  332,
		MessageTypes: mapset.NewSetWith(
			Get,
			GetNext,
			Set,
		),
		AllowedAttributeMask: 0XFFF0,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, Read, false, false, false),
			1:  MultiByteField("OltCryptoCapabilities", 16, nil, Write, false, false, false),
			2:  TableField("OltRandomChallengeTable", TableInfo{17, nil, 17}, Read|Write, false, false),
			3:  ByteField("OltChallengeStatus", 0, Read|Write, false, false, false),
			4:  ByteField("OnuSelectedCryptoCapabilities", 0, Read, false, false, false),
			5:  TableField("OnuRandomChallengeTable", TableInfo{16, nil, 16}, Read, true, false),
			6:  TableField("OnuAuthenticationResultTable", TableInfo{16, nil, 16}, Read, true, false),
			7:  TableField("OltAuthenticationResultTable", TableInfo{17, nil, 17}, Write, false, false),
			8:  ByteField("OltResultStatus", 0, Read|Write, false, false, false),
			9:  ByteField("OnuAuthenticationStatus", 0, Read, true, false, false),
			10: MultiByteField("MasterSessionKeyName", 16, nil, Read, false, false, false),
			11: TableField("BroadcastKeyTable", TableInfo{18, nil, 18}, Read|Write, false, true),
			12: Uint16Field("EffectiveKeyLength", 0, Read, false, false, true),
		},
	}
}

// NewEnhancedSecurityControl (class ID 332 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewEnhancedSecurityControl(params ...ParamData) (IManagedEntity, error) {
	entity := &ManagedEntity {
	    Definition: enhancedsecuritycontrolBME,
	    Attributes: make(map[string]interface{}),
	}
	if err := entity.setAttributes(params...); err != nil {
	    return nil, err
	}
	return entity, nil
}
