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

// Vdsl2LineInventoryAndStatusDataPart2ClassID is the 16-bit ID for the OMCI
// Managed entity VDSL2 line inventory and status data part 2
const Vdsl2LineInventoryAndStatusDataPart2ClassID ClassID = ClassID(169)

var vdsl2lineinventoryandstatusdatapart2BME *ManagedEntityDefinition

// Vdsl2LineInventoryAndStatusDataPart2 (class ID #169)
//	This ME extends the xDSL line configuration MEs. The ME name was chosen because its attributes
//	were initially unique to ITU-T G.993.2 VDSL2. Due to continuing standards development, some
//	attributes - and therefore this ME - have also become applicable to other Recommendations,
//	specifically [ITU-T G.992.3] and [ITU-T G.992.5].
//
//	This ME contains upstream attributes.
//
//	Relationships
//		This is one of the status data MEs associated with an xDSL UNI. It is meaningful if the PPTP
//		supports [ITU-T G.992.3], [ITU-T G.992.5] or [ITU-T G.993.2]. The ONU automatically creates or
//		deletes an instance of this ME upon creation and deletion of a PPTP xDSL UNI part 1 that
//		supports these attributes.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an
//			identical ID, this ME is implicitly linked to an instance of the PPTP xDSL UNI part 1 ME. (R)
//			(mandatory) (2-bytes)
//
//		Vdsl2 Transmission System Capability Xtu_R
//			VDSL2 transmission system capability xTU-R: This attribute extends the xTU-R transmission system
//			capability attribute of the xDSL line inventory and status data part 1 to include xTUR VDSL2
//			capabilities. It is a defined by bits 57..64 of Table 9.7.12-1. (R) (mandatory) (1-byte)
//
//		Actsnrmodeus
//			(R) (mandatory) (1-byte)
//
//		Upbokle
//			UPBOKLE:	This attribute contains the electrical length estimated by the VTU-O expressed in
//			decibels at 1-MHz, kl0 (see O-UPDATE in clause 12.3.3.2.1.2 of [ITUT-G.993.2]). This is the
//			final electrical length that would have been sent from the VTU-O to the VTU-R if the electrical
//			length were not forced by the OLT. The value lies in the range 0 (0.0-dB) to 1280 (128.0-dB) (R)
//			(mandatory) (2-bytes)
//
//		Hlingus
//			HLINGus:	This attribute is the number of subcarriers per group used to report HLINpsus. (R)
//			(mandatory) (1-byte)
//
//		Hloggus
//			HLOGGus:	This attribute is the number of subcarriers per group used to report HLOGpsus. (R)
//			(mandatory) (1-byte)
//
//		Qlngus
//			QLNGus:	This attribute is the number of subcarriers per group used to report QLNpsus. (R)
//			(mandatory) (1-byte)
//
//		Snrgus
//			SNRGus:	This attribute is the number of subcarriers per group used to report SNRpsus. (R)
//			(mandatory) (1-byte)
//
//		Mrefpsdus Table
//			(R) (mandatory) (3 * N bytes, where N is the number of breakpoints)
//
//		Trellisus
//			(R) (mandatory for ITU-T G.993.2 VDSL2, optional for others) (1-byte)
//
//		Actualce
//			ACTUALCE: This attribute reports the cyclic extension used on the line. It is coded as an
//			unsigned integer from 2 to 16 in units of N/32 samples, where 2N is the IDFT size. (R)
//			(mandatory) (1-byte)
//
//		Upbokle_R
//			UPBOKLE-R: This attribute contains the electrical length estimated by the VTU-R expressed in
//			decibels at 1-MHz. This is the value contained in the message RMSG1 (see clause 12.3.3.2.2.1of
//			[ITUT G.993.2]). Its value lies in the range 0 (0.0-dB) to 1280 (128.0-dB) (R) (optional)
//			(2-bytes)
//
//		Actual Rate Adaptation Mode Upstream
//			(R) (optional) (1-byte)
//
//		Actual Impulse Noise Protection Roc Upstream
//			Actual impulse noise protection ROC upstream: The ACTINP-ROC-us attribute reports the actual INP
//			of the ROC in the upstream direction expressed in multiples of T4k. The INP of this attribute is
//			equal to the integer value multiplied by 0.1 symbols. Valid values and usage are given in clause
//			7.5.1.34.2 of [ITUT-G.997.1]. (R) (optional) (1-byte)
//
//		Snr Margin Roc Upstream
//			SNR margin ROC upstream: The SNRM-ROC-us attribute reports the actual signal-to-noise margin of
//			the ROC in the upstream direction. Its value ranges from 0  (-64.0-dB) to 1270 (+63.0-dB). The
//			special value 0xFFFF indicates that the attribute is out of range. (R) (optional) (2-bytes)
//
type Vdsl2LineInventoryAndStatusDataPart2 struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	vdsl2lineinventoryandstatusdatapart2BME = &ManagedEntityDefinition{
		Name:    "Vdsl2LineInventoryAndStatusDataPart2",
		ClassID: 169,
		MessageTypes: mapset.NewSetWith(
			Get,
			GetNext,
		),
		AllowedAttributeMask: 0xfffc,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read), false, false, false, false, 0),
			1:  ByteField("Vdsl2TransmissionSystemCapabilityXtuR", 0, mapset.NewSetWith(Read), false, false, false, false, 1),
			2:  ByteField("Actsnrmodeus", 0, mapset.NewSetWith(Read), false, false, false, false, 2),
			3:  Uint16Field("Upbokle", 0, mapset.NewSetWith(Read), false, false, false, false, 3),
			4:  ByteField("Hlingus", 0, mapset.NewSetWith(Read), false, false, false, false, 4),
			5:  ByteField("Hloggus", 0, mapset.NewSetWith(Read), false, false, false, false, 5),
			6:  ByteField("Qlngus", 0, mapset.NewSetWith(Read), false, false, false, false, 6),
			7:  ByteField("Snrgus", 0, mapset.NewSetWith(Read), false, false, false, false, 7),
			8:  TableField("MrefpsdusTable", TableInfo{nil, 3}, mapset.NewSetWith(Read), false, false, false, 8),
			9:  ByteField("Trellisus", 0, mapset.NewSetWith(Read), false, false, false, false, 9),
			10: ByteField("Actualce", 0, mapset.NewSetWith(Read), false, false, false, false, 10),
			11: Uint16Field("UpbokleR", 0, mapset.NewSetWith(Read), false, false, true, false, 11),
			12: ByteField("ActualRateAdaptationModeUpstream", 0, mapset.NewSetWith(Read), false, false, true, false, 12),
			13: ByteField("ActualImpulseNoiseProtectionRocUpstream", 0, mapset.NewSetWith(Read), false, false, true, false, 13),
			14: Uint16Field("SnrMarginRocUpstream", 0, mapset.NewSetWith(Read), false, false, true, false, 14),
		},
		Access:  CreatedByOnu,
		Support: UnknownSupport,
	}
}

// NewVdsl2LineInventoryAndStatusDataPart2 (class ID 169) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewVdsl2LineInventoryAndStatusDataPart2(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*vdsl2lineinventoryandstatusdatapart2BME, params...)
}
