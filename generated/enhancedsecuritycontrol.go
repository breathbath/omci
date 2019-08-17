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

const EnhancedSecurityControlClassId ClassID = ClassID(332)

var enhancedsecuritycontrolBME *ManagedEntityDefinition

// EnhancedSecurityControl (class ID #332) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type EnhancedSecurityControl struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	enhancedsecuritycontrolBME = &ManagedEntityDefinition{
		Name:    "EnhancedSecurityControl",
		ClassID: 332,
		MessageTypes: mapset.NewSetWith(
			Get,
			GetNext,
			Set,
		),
		AllowedAttributeMask: 0XFFF0,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read), false, false, false, false, 0),
			1:  MultiByteField("OltCryptoCapabilities", 16, nil, mapset.NewSetWith(Write), false, false, false, false, 1),
			2:  TableField("OltRandomChallengeTable", TableInfo{nil, 17}, mapset.NewSetWith(Read, Write), false, false, false, 2),
			3:  ByteField("OltChallengeStatus", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 3),
			4:  ByteField("OnuSelectedCryptoCapabilities", 0, mapset.NewSetWith(Read), false, false, false, false, 4),
			5:  TableField("OnuRandomChallengeTable", TableInfo{nil, 16}, mapset.NewSetWith(Read), true, false, false, 5),
			6:  TableField("OnuAuthenticationResultTable", TableInfo{nil, 16}, mapset.NewSetWith(Read), true, false, false, 6),
			7:  TableField("OltAuthenticationResultTable", TableInfo{nil, 17}, mapset.NewSetWith(Read, Write), false, false, false, 7),
			8:  ByteField("OltResultStatus", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 8),
			9:  ByteField("OnuAuthenticationStatus", 0, mapset.NewSetWith(Read), true, false, false, false, 9),
			10: MultiByteField("MasterSessionKeyName", 16, nil, mapset.NewSetWith(Read), false, false, false, false, 10),
			11: TableField("BroadcastKeyTable", TableInfo{nil, 18}, mapset.NewSetWith(Read, Write), false, true, false, 11),
			12: Uint16Field("EffectiveKeyLength", 0, mapset.NewSetWith(Read), false, false, true, false, 12),
		},
	}
}

// NewEnhancedSecurityControl (class ID 332 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewEnhancedSecurityControl(params ...ParamData) (*ManagedEntity, error) {
	return NewManagedEntity(enhancedsecuritycontrolBME, params...)
}
