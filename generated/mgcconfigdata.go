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

// MgcConfigDataClassID is the 16-bit ID for the OMCI
// Managed entity MGC config data
const MgcConfigDataClassID ClassID = ClassID(155)

var mgcconfigdataBME *ManagedEntityDefinition

// MgcConfigData (class ID #155)
//	The MGC config data ME defines the MGC configuration associated with an MG subscriber. It is
//	conditionally required for ONUs that support ITU-T H.248 VoIP services. If a non-OMCI interface
//	is used to manage VoIP signalling, this ME is unnecessary.
//
//	Instances of this ME are created and deleted by the OLT.
//
//	Relationships
//		An instance of this ME may be associated with one or more VoIP voice CTP MEs.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. (R, setbycreate)
//			(mandatory) (2-bytes)
//
//		Primary Mgc
//			Primary MGC: This attribute points to a network address ME that contains the name (IP-address or
//			resolved name) of the primary MGC that controls the signalling messages. The port is optional
//			and defaults to 2944 for text message formats and 2955 for binary message formats. (R,-W,
//			setbycreate) (mandatory) (2-bytes)
//
//		Secondary Mgc
//			Secondary MGC: This attribute points to a network address ME that contains the name (IP-address
//			or resolved name) of the secondary or backup MGC that controls the signalling messages. The port
//			is optional and defaults to 2944 for text message formats and 2955 for binary message formats.
//			(R,-W, setbycreate) (mandatory) (2-bytes)
//
//		Tcp_Udp Pointer
//			TCP/UDP pointer: This attribute points to the TCP/UDP config data ME to be used for
//			communication with the MGC. (R,-W, setbycreate) (mandatory) (2-bytes)
//
//		Version
//			Version:	This integer attribute reports the version of the Megaco protocol in use. The ONU
//			should deny an attempt by the OLT to set or create a value that it does not support. The value 0
//			indicates that no particular version is specified. (R,-W, setbycreate) (mandatory) (1-byte)
//
//		Message Format
//			The default value is recommended to be 0. (R,-W, setbycreate) (mandatory) (1-byte)
//
//		Maximum Retry Time
//			Maximum retry time: This attribute specifies the maximum retry time for MGC transactions, in
//			seconds. The default value 0 specifies vendor-specific implementation. (R,-W) (optional)
//			(2-bytes)
//
//		Maximum Retry Attempts
//			Maximum retry attempts: This attribute specifies the maximum number of times a message is
//			retransmitted to the MGC. The recommended default value 0 specifies vendor-specific
//			implementation. (R,-W, setbycreate) (optional) (2-bytes)
//
//		Service Change Delay
//			Service change delay: This attribute specifies the service status delay time for changes in line
//			service status. This attribute is specified in seconds. The default value 0 specifies no delay.
//			(R,-W) (optional) (2-bytes)
//
//		Termination Id Base
//			Termination ID base: This attribute specifies the base string for the ITU-T H.248 physical
//			termination ID(s) for this ONU. This string is intended to uniquely identify an ONU. Vendor-
//			specific termination identifiers (port IDs) are optionally added to this string to uniquely
//			identify a termination on a specific ONU. (R,-W) (optional) (25-bytes)
//
//		Softswitch
//			Softswitch:	This attribute identifies the gateway softswitch vendor. The format is four ASCII
//			coded alphabetic characters [A..Z] as defined in [ATIS0300220]. A value of four null bytes
//			indicates an unknown or unspecified vendor. (R,-W, setbycreate) (mandatory) (4-bytes)
//
//		Message Id Pointer
//			Message ID pointer: This attribute points to a large string whose value specifies the message
//			identifier string for ITU-T H.248 messages originated by the ONU. (R, W, setbycreate) (optional)
//			(2 bytes)
//
type MgcConfigData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	mgcconfigdataBME = &ManagedEntityDefinition{
		Name:    "MgcConfigData",
		ClassID: 155,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0xffe0,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", PointerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, 0),
			1:  Uint16Field("PrimaryMgc", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 1),
			2:  Uint16Field("SecondaryMgc", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 2),
			3:  Uint16Field("TcpUdpPointer", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 3),
			4:  ByteField("Version", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 4),
			5:  ByteField("MessageFormat", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 5),
			6:  Uint16Field("MaximumRetryTime", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, true, false, 6),
			7:  Uint16Field("MaximumRetryAttempts", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, true, false, 7),
			8:  Uint16Field("ServiceChangeDelay", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, true, false, 8),
			9:  MultiByteField("TerminationIdBase", OctetsAttributeType, 25, toOctets("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=="), mapset.NewSetWith(Read, Write), false, true, false, 9),
			10: Uint32Field("Softswitch", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 10),
			11: Uint16Field("MessageIdPointer", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, true, false, 11),
		},
		Access:  CreatedByOlt,
		Support: UnknownSupport,
	}
}

// NewMgcConfigData (class ID 155) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewMgcConfigData(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*mgcconfigdataBME, params...)
}
