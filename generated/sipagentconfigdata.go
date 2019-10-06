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

const SipAgentConfigDataClassId ClassID = ClassID(150)

var sipagentconfigdataBME *ManagedEntityDefinition

// SipAgentConfigData (class ID #150)
//	The SIP agent config data ME models a SIP signalling agent. It defines the configuration
//	necessary to establish communication for signalling between the SIP user agent (UA) and a SIP
//	server.
//
//	NOTE 1 - If a non-OMCI interface is used to manage SIP for VoIP, this ME is unnecessary. The
//	non-OMCI interface supplies the necessary data, which may be read back to the OLT via the SIP
//	config portal ME.
//
//	Instances of this ME are created and deleted by the OLT.
//
//	Relationships
//		An instance of this ME serves one or more SIP user data MEs and points to a TCP/UDP config data
//		that carries signalling messages. Other pointers establish additional agent parameters such as
//		proxy servers.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. (R, setbycreate)
//			(mandatory) (2-bytes)
//
//		Proxy Server Address Pointer
//			Proxy server address pointer: This attribute points to a large string ME that contains the name
//			(IP address or URI) of the SIP proxy server for SIP signalling messages. (R,-W, setbycreate)
//			(mandatory) (2-bytes)
//
//		Outbound Proxy Address Pointer
//			Outbound proxy address pointer: An outbound SIP proxy may or may not be required within a given
//			network. If an outbound SIP proxy is used, the outbound proxy address pointer attribute must be
//			set to point to a valid large string ME that contains the name (IP address or URI) of the
//			outbound proxy server for SIP signalling messages. If an outbound SIP proxy is not used, the
//			outbound proxy address pointer attribute must be set to a null pointer. (R,-W, setbycreate)
//			(mandatory) (2-bytes)
//
//		Primary Sip Dns
//			Primary SIP DNS: This attribute specifies the primary SIP DNS IP address. If the value of this
//			attribute is 0, the primary DNS server is defined in the corresponding IP host config data or
//			IPv6 host config data ME. If the value is non-zero, it takes precedence over the primary DNS
//			server defined in the IP host config data or IPv6 host config data ME. (R,-W, set-by-create)
//			(mandatory) (4-bytes)
//
//		Secondary Sip Dns
//			Secondary SIP DNS: This attribute specifies the secondary SIP DNS IP address. If the value of
//			this attribute is 0, the secondary DNS server is defined in the corresponding IP host config
//			data or IPv6 host config data ME. If the value is non-zero, it takes precedence over the
//			secondary DNS server defined in the IP host config data or IPv6 host config data ME. (R,-W, set-
//			by-create) (mandatory) (4-bytes)
//
//		Tcp_Udp Pointer
//			TCP/UDP pointer: This pointer associates the SIP agent with the TCP/UDP config data ME to be
//			used for communication with the SIP server. The default value is 0xFFFF, a null pointer. (R,-W)
//			(mandatory) (2-bytes)
//
//		Sip Reg Exp Time
//			SIP reg exp time: This attribute specifies the SIP registration expiration time in seconds. If
//			its value is 0, the SIP agent does not add an expiration time to the registration requests and
//			does not perform reregistration. The default value is 3600-s. (R,-W) (mandatory) (4-bytes)
//
//		Sip Rereg Head Start Time
//			SIP rereg head start time: This attribute specifies the time in seconds prior to timeout that
//			causes the SIP agent to start the re-registration process. The default value is 360-s. (R,-W)
//			(mandatory) (4-bytes)
//
//		Host Part Uri
//			Host part URI: This attribute points to a large string ME that contains the host or domain part
//			of the SIP address of record for users connected to this ONU. A null pointer indicates that the
//			current address in the IP host config ME is to be used. (R,-W, setbycreate) (mandatory)
//			(2-bytes)
//
//		Sip Status
//			(R) (mandatory) (1-byte)
//
//		Sip Registrar
//			SIP registrar: This attribute points to a network address ME that contains the name (IP address
//			or resolved name) of the registrar server for SIP signalling messages. Examples: "10.10.10.10"
//			and "proxy.voip.net". (R,-W, set-by-create) (mandatory) (2-bytes)
//
//		Softswitch
//			Softswitch:	This attribute identifies the SIP gateway softswitch vendor. The format is four
//			ASCII coded alphabetic characters [A..Z] as defined in [ATIS0300220]. A value of four null bytes
//			indicates an unknown or unspecified vendor. (R,-W, setbycreate) (mandatory) (4-bytes)
//
//		Sip Response Table
//			NOTE 2 - This model assumes that SIP response tones and text are common to all POTS lines that
//			share a given SIP agent.
//
//		Sip Option Transmit Control
//			SIP option transmit control: This Boolean attribute specifies that the ONU is (true) or is not
//			(false) enabled to transmit SIP options. The default value is recommended to be false. (R, W,
//			setbycreate) (optional) (1 byte)
//
//		Sip Uri Format
//			SIP URI format: This attribute specifies the format of the URI in outgoing SIP messages. The
//			recommended default value 0 specifies TEL URIs; the value 1 specifies SIP URIs. Other values are
//			reserved. (R, W, setbycreate) (optional) (1 byte)
//
//		Redundant Sip Agent Pointer
//			Redundant SIP agent pointer: This attribute points to another SIP agent config data ME, which is
//			understood to provide redundancy. The initial SIP agent is determined by the pointer from the
//			SIP user data ME. It is the manager's responsibility to provision a group of redundant SIP
//			agents with mutually consistent attributes. (R, W, setbycreate) (optional) (2 bytes)
//
type SipAgentConfigData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	sipagentconfigdataBME = &ManagedEntityDefinition{
		Name:    "SipAgentConfigData",
		ClassID: 150,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFFFE,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, false, 0),
			1:  Uint16Field("ProxyServerAddressPointer", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 1),
			2:  Uint16Field("OutboundProxyAddressPointer", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 2),
			3:  Uint32Field("PrimarySipDns", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 3),
			4:  Uint32Field("SecondarySipDns", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 4),
			5:  Uint16Field("TcpUdpPointer", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 5),
			6:  Uint32Field("SipRegExpTime", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 6),
			7:  Uint32Field("SipReregHeadStartTime", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 7),
			8:  Uint16Field("HostPartUri", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 8),
			9:  ByteField("SipStatus", 0, mapset.NewSetWith(Read), true, false, false, false, 9),
			10: Uint16Field("SipRegistrar", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 10),
			11: Uint32Field("Softswitch", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 11),
			12: MultiByteField("SipResponseTable", 5, nil, mapset.NewSetWith(Read, Write), false, false, true, false, 12),
			13: ByteField("SipOptionTransmitControl", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, true, false, 13),
			14: ByteField("SipUriFormat", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, true, false, 14),
			15: Uint16Field("RedundantSipAgentPointer", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, true, false, 15),
		},
	}
}

// NewSipAgentConfigData (class ID 150 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewSipAgentConfigData(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*sipagentconfigdataBME, params...)
}
