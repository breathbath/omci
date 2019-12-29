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

// VoipVoiceCtpClassID is the 16-bit ID for the OMCI
// Managed entity VoIP voice CTP
const VoipVoiceCtpClassID ClassID = ClassID(139)

var voipvoicectpBME *ManagedEntityDefinition

// VoipVoiceCtp (class ID #139)
//	The VoIP voice CTP defines the attributes necessary to associate a specified VoIP service (SIP,
//	ITUT-H.248) with a POTS UNI. This entity is conditionally required for ONUs that offer VoIP
//	services. If a non-OMCI interface is used to manage VoIP signalling, this ME is unnecessary.
//
//	An instance of this ME is created and deleted by the OLT. A VoIP voice CTP ME is needed for each
//	PPTP POTS UNI served by VoIP.
//
//	Relationships
//		An instance of this ME links a PPTP POTS UNI ME with a VoIP media profile and a SIP user data or
//		media gateway controller (MGC) config data ME.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. (R, setbycreate)
//			(mandatory) (2-bytes)
//
//		User Protocol Pointer
//			User protocol pointer: This attribute points to signalling protocol data. If the signalling
//			protocol used attribute of the VoIP config data ME specifies that the ONU's signalling protocol
//			is SIP, this attribute points to a SIP user data ME, which in turn points to a SIP agent config
//			data ME. If the signalling protocol is ITU-T-H.248, this attribute points directly to an MGC
//			config data ME. (R,-W, setbycreate) (mandatory) (2-bytes)
//
//		Pptp Pointer
//			PPTP pointer: This attribute points to the PPTP POTS UNI ME that serves the analogue telephone
//			port. (R,-W, setbycreate) (mandatory) (2-bytes)
//
//		V O Ip Media Profile Pointer
//			VoIP media profile pointer: This attribute points to an associated VoIP media profile. (R,-W,
//			setbycreate) (mandatory) (2-bytes)
//
//		Signalling Code
//			(R,-W, setbycreate) (mandatory) (1-byte)
//
type VoipVoiceCtp struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	voipvoicectpBME = &ManagedEntityDefinition{
		Name:    "VoipVoiceCtp",
		ClassID: 139,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0xf000,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, false, 0),
			1: Uint16Field("UserProtocolPointer", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 1),
			2: Uint16Field("PptpPointer", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 2),
			3: Uint16Field("VOIpMediaProfilePointer", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 3),
			4: ByteField("SignallingCode", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 4),
		},
		Access:  CreatedByOlt,
		Support: UnknownSupport,
	}
}

// NewVoipVoiceCtp (class ID 139) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewVoipVoiceCtp(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*voipvoicectpBME, params...)
}
