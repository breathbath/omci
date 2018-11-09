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

// EnhancedSecurityControl (class ID 332 defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type EnhancedSecurityControl struct {
	BaseManagedEntityDefinition
}

// NewEnhancedSecurityControl (class ID 332 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewEnhancedSecurityControl(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "EnhancedSecurityControl",
		ClassID:  332,
		EntityID: eid,
		MessageTypes: []MsgType{
			Set,
			Get,
			GetNext,
		},
		AttributeMask: 0,
		Attributes: []*AttributeDefinition{
			Uint16Field("ManagedEntityId", 0, Read),
			UnknownField("OltCryptoCapabilities", 0, Write),
			UnknownField("OltRandomChallengeTable", 0, Read|Write),
			ByteField("OltChallengeStatus", 0, Read|Write),
			ByteField("OnuSelectedCryptoCapabilities", 0, Read),
			UnknownField("OnuRandomChallengeTable", 0, Read),
			UnknownField("OnuAuthenticationResultTable", 0, Read),
			UnknownField("OltAuthenticationResultTable", 0, Write),
			ByteField("OltResultStatus", 0, Read|Write),
			ByteField("OnuAuthenticationStatus", 0, Read),
			UnknownField("MasterSessionKeyName", 0, Read),
			UnknownField("BroadcastKeyTable", 0, Read|Write),
			Uint16Field("EffectiveKeyLength", 0, Read),
		},
	}
	entity.computeAttributeMask()
	return &EnhancedSecurityControl{entity}, nil
}
