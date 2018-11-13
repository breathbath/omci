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

const SipUserDataClassId uint16 = 153

// SipUserData (class ID #153) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type SipUserData struct {
	BaseManagedEntityDefinition
}

// NewSipUserData (class ID 153 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewSipUserData(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "SipUserData",
		ClassID:  153,
		EntityID: eid,
		MessageTypes: []MsgType{
			Set,
			Get,
			Create,
			Delete,
		},
		AllowedAttributeMask: 0,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, Read|SetByCreate),
			1: Uint16Field("SipAgentPointer", 0, Read|Write|SetByCreate),
			2: Uint16Field("UserPartAor", 0, Read|Write|SetByCreate),
			3: MultiByteField("SipDisplayName", 25, nil, Read|Write),
			4: Uint16Field("UsernameAndPassword", 0, Read|Write|SetByCreate),
			5: Uint16Field("VoicemailServerSipUri", 0, Read|Write|SetByCreate),
			6: Uint32Field("VoicemailSubscriptionExpirationTime", 0, Read|Write|SetByCreate),
			7: Uint16Field("NetworkDialPlanPointer", 0, Read|Write|SetByCreate),
			8: Uint16Field("ApplicationServicesProfilePointer", 0, Read|Write|SetByCreate),
			9: Uint16Field("FeatureCodePointer", 0, Read|Write|SetByCreate),
			10: Uint16Field("PptpPointer", 0, Read|Write|SetByCreate),
			11: ByteField("ReleaseTimer", 0, Read|Write),
			12: ByteField("ReceiverOffHookRohTimer", 0, Read|Write),
		},
	}
	entity.computeAttributeMask()
	return &SipUserData{entity}, nil
}
