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

// Dot1XConfigurationProfileClassID is the 16-bit ID for the OMCI
// Managed entity Dot1X configuration profile
const Dot1XConfigurationProfileClassID ClassID = ClassID(291)

var dot1xconfigurationprofileBME *ManagedEntityDefinition

// Dot1XConfigurationProfile (class ID #291)
//	An instance of this ME represents a set of attributes that control an ONU's 802.1X operation
//	with regard to IEEE 802 services. An instance of this ME is created by the ONU if it is capable
//	of supporting [IEEE 802.1X] authentication of CPE.
//
//	Relationships
//		One instance of this ME governs the ONU's 802.1X CPE authentication behaviour.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute provides a unique number for each instance of this ME. There
//			is at most one instance, number 0. (R) (mandatory) (2-bytes)
//
//		Circuit Id Prefix
//			Circuit ID prefix: This attribute is a pointer to a large string ME whose content appears as the
//			prefix of the NAS port ID in radius access-request messages. The remainder of the NAS port ID
//			field is local information (for example, slot-port, appended by the ONU itself). The default
//			value of this attribute is the null pointer 0. (R,-W) (mandatory) (2-bytes)
//
//		Fallback Policy
//			Fallback policy: When set to 1 (deny), this attribute causes IEEE-802.1X conversations to fail
//			when no external authentication server is accessible, such that no Ethernet service is provided.
//			The default value 0 causes IEEE-802.1X conversations to succeed when no external authentication
//			server is accessible. (R,-W) (mandatory) (1-byte)
//
//		Auth Server 1
//			Auth server 1: This attribute is a pointer to a large string ME that contains the URI of the
//			first choice radius authentication server. The value 0 indicates that no radius authentication
//			server is specified. (R,-W) (mandatory) (2-bytes)
//
//		Shared Secret Auth1
//			Shared secret auth1: This attribute is the shared secret for the first radius authentication
//			server. It is a null-terminated character string. (R,-W) (mandatory) (25-bytes)
//
//		Auth Server 2
//			Auth server 2:	(R,-W) (optional) (2-bytes)
//
//		Shared Secret Auth2
//			Shared secret auth2:	(R,-W) (optional) (25-bytes)
//
//		Auth Server 3
//			Auth server 3:	(R,-W) (optional) (2-bytes)
//
//		Shared Secret Auth3
//			Shared secret auth3:	(R,-W) (optional) (25-bytes)
//
//		Olt Proxy Address
//			OLT proxy address: This attribute indicates the IP address of a possible proxy at the OLT for
//			IEEE-802.1X radius messages. The default value 0.0.0.0 indicates that no proxy is required.
//			(R,-W) (optional) (4-bytes)
//
//		Calling Station Id Format
//			Other values are reserved.
//
type Dot1XConfigurationProfile struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	dot1xconfigurationprofileBME = &ManagedEntityDefinition{
		Name:    "Dot1XConfigurationProfile",
		ClassID: 291,
		MessageTypes: mapset.NewSetWith(
			Get,
			Set,
		),
		AllowedAttributeMask: 0xffc0,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read), false, false, false, false, 0),
			1:  Uint16Field("CircuitIdPrefix", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 1),
			2:  ByteField("FallbackPolicy", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 2),
			3:  Uint16Field("AuthServer1", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 3),
			4:  MultiByteField("SharedSecretAuth1", 25, nil, mapset.NewSetWith(Read, Write), false, false, false, false, 4),
			5:  Uint16Field("AuthServer2", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 5),
			6:  MultiByteField("SharedSecretAuth2", 25, nil, mapset.NewSetWith(Read, Write), false, false, true, false, 6),
			7:  Uint16Field("AuthServer3", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 7),
			8:  MultiByteField("SharedSecretAuth3", 25, nil, mapset.NewSetWith(Read, Write), false, false, true, false, 8),
			9:  Uint32Field("OltProxyAddress", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 9),
			10: Uint16Field("CallingStationIdFormat", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 10),
		},
		Access:  CreatedByOnu,
		Support: UnknownSupport,
	}
}

// NewDot1XConfigurationProfile (class ID 291) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewDot1XConfigurationProfile(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*dot1xconfigurationprofileBME, params...)
}
