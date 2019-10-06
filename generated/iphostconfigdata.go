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

const IpHostConfigDataClassId ClassID = ClassID(134)

var iphostconfigdataBME *ManagedEntityDefinition

// IpHostConfigData (class ID #134)
//	The IP host config data configures IPv4 based services offered on the ONU. The ONU automatically
//	creates instances of this ME if IP host services are available. A possible IPv6 stack is
//	supported through the IPv6 host config data ME. In this clause, references to IP addresses are
//	understood to mean IPv4.
//
//	Relationships
//		An instance of this ME is associated with the ONU ME. Any number of TCP/UDP config data MEs can
//		point to the IP host config data, to model any number of ports and protocols. Performance may be
//		monitored through an implicitly linked IP host PM history data ME.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. The ONU creates
//			as many instances as there are independent IPv4 stacks on the ONU. To facilitate discovery, IP
//			host config data MEs should be numbered from 0 upwards. The ONU should create IP(v4) and IPv6
//			host config data MEs with separate ME IDs, such that other MEs can use a single TP type
//			attribute to link with either. (R) (mandatory) (2 bytes)
//
//		Ip Options
//			(R,-W) (mandatory) (1-byte)
//
//		Mac Address
//			MAC address: This attribute indicates the MAC address used by the IP node. (R) (mandatory)
//			(6-bytes)
//
//		Onu Identifier
//			Onu identifier: A unique ONU identifier string. If set to a non-null value, this string is used
//			instead of the MAC address in retrieving dynamic host configuration protocol (DHCP) parameters.
//			If the string is shorter than 25 characters, it must be null terminated. Its default value is 25
//			null bytes. (R,-W) (mandatory) (25-bytes)
//
//		Ip Address
//			IP address:	The address used for IP host services; this attribute has the default value 0.
//			(R,-W) (mandatory) (4-bytes)
//
//		Mask
//			Mask:	The subnet mask for IP host services; this attribute has the default value 0. (R,-W)
//			(mandatory) (4-bytes)
//
//		Gateway
//			Gateway:	The default gateway address used for IP host services; this attribute has the default
//			value 0. (R,-W) (mandatory) (4-bytes)
//
//		Primary Dns
//			Primary DNS: The address of the primary DNS server; this attribute has the default value 0.
//			(R,-W) (mandatory) (4-bytes)
//
//		Secondary Dns
//			Secondary DNS: The address of the secondary DNS server; this attribute has the default value 0.
//			(R,-W) (mandatory) (4-bytes)
//
//		Current Address
//			Current address: Current address of the IP host service. (R) (optional) (4-bytes)
//
//		Current Mask
//			Current mask: Current subnet mask for the IP host service. (R) (optional) (4-bytes)
//
//		Current Gateway
//			Current gateway: Current default gateway address for the IP host service. (R) (optional)
//			(4-bytes)
//
//		Current Primary Dns
//			Current primary DNS: Current primary DNS server address. (R) (optional) (4-bytes)
//
//		Current Secondary Dns
//			Current secondary DNS: Current secondary DNS server address. (R) (optional) (4-bytes)
//
//		Domain Name
//			Domain name: If DHCP indicates a domain name, it is presented here. If no domain name is
//			indicated, this attribute is set to a null string. If the string is shorter than 25-bytes, it
//			must be null terminated. The default value is 25 null bytes. (R) (mandatory) (25-bytes)
//
//		Host Name
//			Host name:	If DHCP indicates a host name, it is presented here. If no host name is indicated,
//			this attribute is set to a null string. If the string is shorter than 25-bytes, it must be null
//			terminated. The default value is 25 null bytes. (R) (mandatory) (25-bytes)
//
//		Relay Agent Options
//			2/3/4:atm/123.4567
//
type IpHostConfigData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	iphostconfigdataBME = &ManagedEntityDefinition{
		Name:    "IpHostConfigData",
		ClassID: 134,
		MessageTypes: mapset.NewSetWith(
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFFFF,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read), false, false, false, false, 0),
			1:  ByteField("IpOptions", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 1),
			2:  MultiByteField("MacAddress", 6, nil, mapset.NewSetWith(Read), false, false, false, false, 2),
			3:  MultiByteField("OnuIdentifier", 25, nil, mapset.NewSetWith(Read, Write), false, false, false, false, 3),
			4:  Uint32Field("IpAddress", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 4),
			5:  Uint32Field("Mask", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 5),
			6:  Uint32Field("Gateway", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 6),
			7:  Uint32Field("PrimaryDns", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 7),
			8:  Uint32Field("SecondaryDns", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 8),
			9:  Uint32Field("CurrentAddress", 0, mapset.NewSetWith(Read), true, false, true, false, 9),
			10: Uint32Field("CurrentMask", 0, mapset.NewSetWith(Read), true, false, true, false, 10),
			11: Uint32Field("CurrentGateway", 0, mapset.NewSetWith(Read), true, false, true, false, 11),
			12: Uint32Field("CurrentPrimaryDns", 0, mapset.NewSetWith(Read), true, false, true, false, 12),
			13: Uint32Field("CurrentSecondaryDns", 0, mapset.NewSetWith(Read), true, false, true, false, 13),
			14: MultiByteField("DomainName", 25, nil, mapset.NewSetWith(Read), true, false, false, false, 14),
			15: MultiByteField("HostName", 25, nil, mapset.NewSetWith(Read), true, false, false, false, 15),
			16: Uint16Field("RelayAgentOptions", 0, mapset.NewSetWith(Read, Write), true, false, true, false, 16),
		},
	}
}

// NewIpHostConfigData (class ID 134 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewIpHostConfigData(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*iphostconfigdataBME, params...)
}
