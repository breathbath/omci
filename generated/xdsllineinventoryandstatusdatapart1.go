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

// XdslLineInventoryAndStatusDataPart1ClassID is the 16-bit ID for the OMCI
// Managed entity xDSL line inventory and status data part 1
const XdslLineInventoryAndStatusDataPart1ClassID ClassID = ClassID(100)

var xdsllineinventoryandstatusdatapart1BME *ManagedEntityDefinition

// XdslLineInventoryAndStatusDataPart1 (class ID #100)
//	This ME contains part 1 of the line inventory and status data for an xDSL UNI. The ONU
//	automatically creates or deletes an instance of this ME upon the creation or deletion of a PPTP
//	xDSL UNI part 1.
//
//	Relationships
//		An instance of this ME is associated with an xDSL UNI.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an
//			identical ID, this ME is implicitly linked to an instance of the PPTP xDSL UNI part 1. (R)
//			(mandatory) (2-bytes)
//
//		Xtu_C G.994.1 Vendor Id
//			xTU-C G.994.1 vendor ID: This is the vendor ID as inserted by the xTU-C in the ITUT-G.994.1 CL
//			message. It comprises 8 octets, including a country code followed by a (regionally allocated)
//			provider code, as defined in [ITUT-T.35]. (R) (mandatory) (8-bytes)
//
//		Xtu_R G.994.1 Vendor Id
//			xTU-R G.994.1 vendor ID: This is the vendor ID as inserted by the xTU-R in the ITUT-G.994.1 CLR
//			message. It comprises 8 binary octets, with the same format as the xTUC ITUT G.994.1 vendor ID.
//			(R) (mandatory) (8-bytes)
//
//		Xtu_C System Vendor Id
//			xTU-C system vendor ID: This is the vendor ID as inserted by the xTU-C in the overhead messages
//			of [ITU-T G.992.3] and [ITU-T G.992.4]. It comprises 8 binary octets, with the same format as
//			the xTU-C ITUT-G.994.1 vendor ID. (R) (mandatory) (8-bytes)
//
//		Xtu_R System Vendor Id
//			xTU-R system vendor ID: This is the vendor ID as inserted by the xTU-R in the embedded
//			operations channel and overhead messages of [ITU-T G.992.3] and [ITUT-G.992.4]. It comprises 8
//			binary octets, with the same format as the xTU-C ITUT-G.994.1 vendor ID. (R) (mandatory)
//			(8-bytes)
//
//		Xtu_C Version Number
//			xTU-C version number: This is the vendorspecific version number as inserted by the xTUC in the
//			overhead messages of [ITU-T G.992.3] and [ITU-T G.992.4]. It comprises up to 16 binary octets.
//			(R) (mandatory) (16-bytes)
//
//		Xtu_R Version Number
//			xTU-R version number: This is the version number as inserted by the xTUR in the embedded
//			operations channel of [ITU-T G.992.1] or [ITU-T G.992.2], or the overhead messages of [ITU-T
//			G.992.3], [ITU-T G.992.4], [ITU-T G.992.5] and [ITU-T G.993.2]. The attribute value may be
//			vendor-specific, but is recommended to comprise up to 16 ASCII characters, null-terminated if it
//			is shorter than 16. The string should contain the xTU-R firmware version and the xTU-R model,
//			encoded in that order and separated by a space character: "<xTU-R firmware version><xTU-R
//			model>". It is recognized that legacy xTU-Rs may not support this format. (R) (mandatory)
//			(16-bytes)
//
//		Xtu_C Serial Number Part 1
//			xTU-C serial number part 1: The vendorspecific serial number inserted by the xTU-C in the
//			overhead messages of [ITU-T G.992.3] and [ITU-T G.992.4] comprises up to 32 ASCII characters,
//			null terminated if it is shorter than 32 characters. This attribute contains the first 16
//			characters. (R) (mandatory) (16-bytes)
//
//		Xtu_C Serial Number Part 2
//			xTU-C serial number part 2: This attribute contains the second 16 characters of the xTU-C serial
//			number. (R) (mandatory) (16-bytes)
//
//		Xtu_R Serial Number Part 1
//			xTU-R serial number part 1: The serial number inserted by the xTU-R in the embedded operations
//			channel of [ITU-T G.992.1] or [ITU-T G.992.2], or the overhead messages of [ITU-T G.992.3],
//			[ITU-T G.992.4], [ITU-T G.992.5] and [ITUT-G.993.2], comprises up to 32 ASCII characters,
//			nullterminated if it is shorter than 32. It is recommended that the equipment serial number, the
//			equipment model and the equipment firmware version, encoded in that order and separated by space
//			characters, be contained: "<equipment serial number><equipment model><equipment firmware
//			version>". It is recognized that legacy xTU-Rs may not support this format. This attribute
//			contains the first 16 characters. (R) (mandatory) (16-bytes)
//
//		Xtu_R Serial Number Part 2
//			xTU-R serial number part 2: This attribute contains the second 16 characters of the xTU-R serial
//			number. (R) (mandatory) (16-bytes)
//
//		Xtu_C Self Test Results
//			xTU-C selftest results: This parameter reports the xTU-C self-test result. It is coded in two
//			fields. The most significant octet is 0 if the self-test passed and 1 if it failed. The three
//			least significant octets are a vendor-discretionary integer that can be interpreted in
//			combination with [ITU-T G.994.1] and the system vendor ID. (R) (mandatory) (4-bytes)
//
//		Xtu_R Self Test Results
//			xTU-R selftest results: This parameter defines the xTU-R self-test result. It is coded in two
//			fields. The most significant octet is 0 if the self-test passed and 1 if it failed. The three
//			least significant octets are a vendor-discretionary integer that can be interpreted in
//			combination with [ITU-T G.994.1] and the system vendor ID. (R) (mandatory) (4-bytes)
//
//		Xtu_C Transmission System Capability
//			NOTE 1 - This attribute is only 7-bytes long. An eighth byte identifying VDSL2 capabilities is
//			defined in the VDSL2 line inventory and status data part 1 ME.
//
//		Xtu_R Transmission System Capability
//			NOTE 2 - This attribute is only 7-bytes long. An eighth byte identifying VDSL2 capabilities is
//			defined in the VDSL2 line inventory and status data part 2 ME.
//
//		Initialization Success_Failure Cause
//			(R) (mandatory) (1-byte)
//
type XdslLineInventoryAndStatusDataPart1 struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	xdsllineinventoryandstatusdatapart1BME = &ManagedEntityDefinition{
		Name:    "XdslLineInventoryAndStatusDataPart1",
		ClassID: 100,
		MessageTypes: mapset.NewSetWith(
			Get,
		),
		AllowedAttributeMask: 0xfffe,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", PointerAttributeType, 0, mapset.NewSetWith(Read), false, false, false, 0),
			1:  Uint64Field("XtuCG9941VendorId", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read), false, false, false, 1),
			2:  Uint64Field("XtuRG9941VendorId", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read), false, false, false, 2),
			3:  Uint64Field("XtuCSystemVendorId", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read), false, false, false, 3),
			4:  Uint64Field("XtuRSystemVendorId", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read), false, false, false, 4),
			5:  MultiByteField("XtuCVersionNumber", OctetsAttributeType, 16, toOctets("AAAAAAAAAAAAAAAAAAAAAA=="), mapset.NewSetWith(Read), false, false, false, 5),
			6:  MultiByteField("XtuRVersionNumber", OctetsAttributeType, 16, toOctets("AAAAAAAAAAAAAAAAAAAAAA=="), mapset.NewSetWith(Read), false, false, false, 6),
			7:  MultiByteField("XtuCSerialNumberPart1", OctetsAttributeType, 16, toOctets("AAAAAAAAAAAAAAAAAAAAAA=="), mapset.NewSetWith(Read), false, false, false, 7),
			8:  MultiByteField("XtuCSerialNumberPart2", OctetsAttributeType, 16, toOctets("AAAAAAAAAAAAAAAAAAAAAA=="), mapset.NewSetWith(Read), false, false, false, 8),
			9:  MultiByteField("XtuRSerialNumberPart1", OctetsAttributeType, 16, toOctets("AAAAAAAAAAAAAAAAAAAAAA=="), mapset.NewSetWith(Read), false, false, false, 9),
			10: MultiByteField("XtuRSerialNumberPart2", OctetsAttributeType, 16, toOctets("AAAAAAAAAAAAAAAAAAAAAA=="), mapset.NewSetWith(Read), false, false, false, 10),
			11: Uint32Field("XtuCSelfTestResults", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read), false, false, false, 11),
			12: Uint32Field("XtuRSelfTestResults", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read), false, false, false, 12),
			13: MultiByteField("XtuCTransmissionSystemCapability", OctetsAttributeType, 7, toOctets("AAAAAAAAAA=="), mapset.NewSetWith(Read), false, false, false, 13),
			14: MultiByteField("XtuRTransmissionSystemCapability", OctetsAttributeType, 7, toOctets("AAAAAAAAAA=="), mapset.NewSetWith(Read), false, false, false, 14),
			15: ByteField("InitializationSuccessFailureCause", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read), false, false, false, 15),
		},
		Access:  CreatedByOnu,
		Support: UnknownSupport,
	}
}

// NewXdslLineInventoryAndStatusDataPart1 (class ID 100) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewXdslLineInventoryAndStatusDataPart1(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*xdsllineinventoryandstatusdatapart1BME, params...)
}
