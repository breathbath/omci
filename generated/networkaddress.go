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

// NetworkAddressClassID is the 16-bit ID for the OMCI
// Managed entity Network address
const NetworkAddressClassID ClassID = ClassID(137)

var networkaddressBME *ManagedEntityDefinition

// NetworkAddress (class ID #137)
//	The network address ME associates a network address with security methods required to access a
//	server. It is conditionally required for ONUs that support VoIP services. The address may take
//	the form of a URL, a fully qualified path or IP address represented as an ACII string.
//
//	If a non-OMCI interface is used to manage VoIP signalling, this ME is unnecessary.
//
//	Instances of this ME are created and deleted by the OLT or the ONU, depending on the method used
//	and case.
//
//	Relationships
//		Any ME that requires a network address may link to an instance of this ME.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. Instances of
//			this ME created autonomously by the ONU have IDs in the range 0..0x7FFF. Instances created by
//			the OLT have IDs in the range 0x8000..0xFFFE. The value 0xFFFF is reserved. (R, setbycreate)
//			(mandatory) (2-bytes)
//
//		Security Pointer
//			Security pointer: This attribute points to an authentication security method ME. The
//			authentication security method indicates the username and password to be used when retrieving
//			the network address indicated by this ME. A null pointer indicates that security attributes are
//			not defined for this network address. (R,-W, setbycreate) (mandatory) (2-bytes)
//
//		Address Pointer
//			Address pointer: This attribute points to the large string ME that contains the network address.
//			It may contain a fully qualified domain name, URI or IP address. The URI may also contain a port
//			identifier (e.g., "x.y.z.com:5060"). A null pointer indicates that no network address is
//			defined. (R,-W, setbycreate) (mandatory) (2-bytes)
//
type NetworkAddress struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	networkaddressBME = &ManagedEntityDefinition{
		Name:    "NetworkAddress",
		ClassID: 137,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0xc000,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, false, 0),
			1: Uint16Field("SecurityPointer", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 1),
			2: Uint16Field("AddressPointer", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 2),
		},
		Access:  CreatedByOlt,
		Support: UnknownSupport,
	}
}

// NewNetworkAddress (class ID 137) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewNetworkAddress(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*networkaddressBME, params...)
}
