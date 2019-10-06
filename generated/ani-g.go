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

const AniGClassId ClassID = ClassID(263)

var anigBME *ManagedEntityDefinition

// AniG (class ID #263)
//	This ME organizes data associated with each access network interface supported by a GPON ONU.
//	The ONU automatically creates one instance of this ME for each PON physical port.
//
//	Relationships
//		An instance of this ME is associated with each instance of a physical PON interface.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. Its value
//			indicates the physical position of the PON interface. The first byte is the slot ID, defined in
//			clause 9.1.5. The second byte is the port ID. (R) (mandatory) (2-bytes)
//
//		Sr Indication
//			SR indication: This Boolean attribute indicates the ONU's capability to report queue status for
//			DBA. The value true means that status reporting is available for all TCONTs that are associated
//			with the ANI. (R) (mandatory) (1-byte)
//
//		Total TCont Number
//			Total TCONT number: This attribute indicates the total number of T-CONTs that can be supported
//			on this ANI. (R) (mandatory) (2-bytes)
//
//		Gem Block Length
//			In all other ITU-T PON systems, the unit for queue occupancy reporting is fixed in at 4-bytes by
//			the respective TC layer specification.
//
//		Piggyback Dba Reporting
//			(R) (mandatory) (1-byte)
//
//		Deprecated
//			Deprecated:	This attribute should be set to 0 by the ONU and ignored by the OLT. (R) (mandatory)
//			(1-byte)
//
//		Signal Fail Threshold
//			Signal fail (SF) threshold: This attribute specifies the downstream bit error rate (BER)
//			threshold to detect the SF alarm. When this value is y, the BER threshold is 10-y. Valid values
//			are 3..8. Upon ME instantiation, the ONU sets this attribute to 5. (R,-W) (mandatory) (1-byte)
//
//		Signal Degrade Sd Threshold
//			Signal degrade (SD) threshold: This attribute specifies the downstream BER threshold to detect
//			the SD alarm. When this value is x, the BER threshold for SD is 10-x. Valid values are 4..10.
//			The SD threshold must be lower than the SF threshold; i.e., x-> y. Upon ME instantiation, the
//			ONU sets this attribute to 9. (R,-W) (mandatory) (1-byte)
//
//		Arc
//			ARC:	See clause A.1.4.3. (R,-W) (optional) (1-byte)
//
//		Arc Interval
//			ARC interval: See clause A.1.4.3. (R,-W) (optional) (1-byte)
//
//		Optical Signal Level
//			Optical signal level: This attribute reports the current measurement of the total downstream
//			optical signal level. Its value is a 2s complement integer referred to 1- mW (i.e., 1-dBm), with
//			0.002 dB granularity. (R) (optional) (2-bytes)
//
//		Lower Optical Threshold
//			Lower optical threshold: This attribute specifies the optical level the ONU uses to declare the
//			downstream low received optical power alarm. Valid values are  -127 dBm (coded as 254) to 0 dBm
//			(coded as 0) in 0.5 dB increments. The default value 0xFF selects the ONU's internal policy.
//			(R,-W) (optional) (1-byte)
//
//		Upper Optical Threshold
//			Upper optical threshold: This attribute specifies the optical level the ONU uses to declare the
//			downstream high received optical power alarm. Valid values are  -127 dBm (coded as 254) to 0 dBm
//			(coded as 0) in 0.5 dB increments. The default value 0xFF selects the ONU's internal policy.
//			(R,-W) (optional) (1-byte)
//
//		Onu Response Time
//			(R) (optional) (2-bytes)
//
//		Transmit Optical Level
//			Transmit optical level: This attribute reports the current measurement of mean optical launch
//			power. Its value is a 2s complement integer referred to 1-mW (i.e., 1-dBm), with 0.002 dB
//			granularity. (R) (optional) (2-bytes)
//
//		Lower Transmit Power Threshold
//			Lower transmit power threshold: This attribute specifies the minimum mean optical launch power
//			that the ONU uses to declare the low transmit optical power alarm. Its value is a 2s complement
//			integer referred to 1-mW (i.e., dBm), with 0.5-dB granularity. The default value -63.5 (0x81)
//			selects the ONU's internal policy. (R,-W) (optional) (1-byte)
//
//		Upper Transmit Power Threshold
//			Upper transmit power threshold: This attribute specifies the maximum mean optical launch power
//			that the ONU uses to declare the high transmit optical power alarm. Its value is a 2s complement
//			integer referred to 1-mW (i.e., dBm), with 0.5-dB granularity. The default value -63.5 (0x81)
//			selects the ONU's internal policy. (R,-W) (optional) (1-byte)
//
type AniG struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	anigBME = &ManagedEntityDefinition{
		Name:    "AniG",
		ClassID: 263,
		MessageTypes: mapset.NewSetWith(
			Get,
			Set,
			Test,
		),
		AllowedAttributeMask: 0XFFFF,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read), false, false, false, false, 0),
			1:  ByteField("SrIndication", 0, mapset.NewSetWith(Read), false, false, false, false, 1),
			2:  Uint16Field("TotalTcontNumber", 0, mapset.NewSetWith(Read), false, false, false, false, 2),
			3:  Uint16Field("GemBlockLength", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 3),
			4:  ByteField("PiggybackDbaReporting", 0, mapset.NewSetWith(Read), false, false, false, false, 4),
			5:  ByteField("Deprecated", 0, mapset.NewSetWith(Read), false, false, false, true, 5),
			6:  ByteField("SignalFailThreshold", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 6),
			7:  ByteField("SignalDegradeSdThreshold", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 7),
			8:  ByteField("Arc", 0, mapset.NewSetWith(Read, Write), true, false, true, false, 8),
			9:  ByteField("ArcInterval", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 9),
			10: Uint16Field("OpticalSignalLevel", 0, mapset.NewSetWith(Read), false, false, true, false, 10),
			11: ByteField("LowerOpticalThreshold", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 11),
			12: ByteField("UpperOpticalThreshold", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 12),
			13: Uint16Field("OnuResponseTime", 0, mapset.NewSetWith(Read), false, false, true, false, 13),
			14: Uint16Field("TransmitOpticalLevel", 0, mapset.NewSetWith(Read), false, false, true, false, 14),
			15: ByteField("LowerTransmitPowerThreshold", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 15),
			16: ByteField("UpperTransmitPowerThreshold", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 16),
		},
	}
}

// NewAniG (class ID 263 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewAniG(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*anigBME, params...)
}
