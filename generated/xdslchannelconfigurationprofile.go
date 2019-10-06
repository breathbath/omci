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

const XdslChannelConfigurationProfileClassId ClassID = ClassID(107)

var xdslchannelconfigurationprofileBME *ManagedEntityDefinition

// XdslChannelConfigurationProfile (class ID #107)
//	This ME contains the channel configuration profile for an xDSL UNI. An instance of this ME is
//	created and deleted by the OLT.
//
//	NOTE - If [ITUT G.997.1] compatibility is required, bit rates should only be set to integer
//	multiples of 1000-bits/s. The ONU may reject attempts to set other values for bit rate
//	attributes.
//
//	Relationships
//		An instance of this ME may be associated with zero or more instances of the PPTP xDSL UNI part
//		1.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. The value 0 is
//			reserved. (R, setbycreate) (mandatory) (2-bytes)
//
//		Minimum Data Rate
//			Minimum data rate: This parameter specifies the minimum desired net data rate for the bearer
//			channel. It is coded in bits per second. (R,-W, setbycreate) (mandatory) (4-bytes)
//
//		Maximum Data Rate
//			Maximum data rate: This parameter specifies the maximum desired net data rate for the bearer
//			channel. It is coded in bits per second. (R,-W, setbycreate) (mandatory) (4-bytes)
//
//		Rate Adaptation Ratio
//			Rate adaptation ratio: This attribute specifies the weight that should be taken into account
//			when performing rate adaptation in the direction of the bearer channel. The attribute is defined
//			as a percentage. The value 20, for example, means that 20% of the available data rate (in excess
//			of the minimum data rate summed over all bearer channels) is assigned to this bearer channel and
//			80% to the other bearer channels. The OLT must ensure that the sum of rate adaptation ratios
//			over all bearers in one direction is 100%. (R,-W, setbycreate) (optional) (1-byte)
//
//		Maximum Interleaving Delay
//			The delay is coded in milliseconds, varying from 2 to 63, with special meaning assigned to
//			values 0, 1 and 255. The value 0 indicates that no delay bound is imposed. The value 1 indicates
//			the fast latency path is to be used in the ITUT G.992.1 operating mode and S and D are to be
//			selected such that S- 1 and D-= 1 in ITU-T G.992.2, ITUT G.992.3, ITUT G.992.4, ITUT G.992.5 and
//			ITUT G.993.2 operating modes. The value 255 indicates a delay bound of 1-ms in ITUT-G.993.2
//			operation. (R,-W, setbycreate) (mandatory) (1-byte)
//
//		Data Rate Threshold Upshift
//			Data rate threshold upshift: This attribute is a threshold on the cumulative data rate upshift
//			achieved over one or more bearer channel data rate adaptations. An upshift rate change (DRT up)
//			notification is issued by the PPTP xDSL UNI part 1 when the actual data rate exceeds the data
//			rate at the last entry into showtime by more than the threshold. The data rate threshold is
//			coded in bits per second. (R,-W, setbycreate) (mandatory for xDSL standards that use this
//			attribute) (4-bytes)
//
//		Data Rate Threshold Downshift
//			Data rate threshold downshift: This attribute is a threshold on the cumulative data rate
//			downshift achieved over one or more bearer channel data rate adaptations. A downshift rate
//			change (DRT down) notification is issued by the PPTP xDSL UNI part 1 when the actual data rate
//			is below the data rate at the last entry into showtime by more than the threshold. The data rate
//			threshold is coded in bits per second. (R,-W, setbycreate) (mandatory for xDSL standards that
//			use this attribute) (4-bytes)
//
//		Minimum Reserved Data Rate
//			Minimum reserved data rate: This attribute specifies the desired minimum reserved net data rate
//			for the bearer channel. The rate is coded in bits per second. This attribute is needed only if
//			the rate adaptation mode is set to dynamic in the xDSL line configuration profile part 1. (R,-W,
//			setbycreate) (optional) (4-bytes)
//
//		Minimum Data Rate In Low _ Power State
//			Minimum data rate in low-power state: This parameter specifies the minimum desired net data rate
//			for the bearer channel during the low-power state (L1/L2). The power management low-power states
//			L1 and L2 are defined in [ITUT-G.992.2] and [ITUT G.992.3], respectively. The data rate is coded
//			in bits per second. (R,-W, setbycreate) (mandatory) (4-byte)
//
//		Minimum Impulse Noise Protection
//			(R,-W, setbycreate) (optional for [ITU-T G.992.1], mandatory for other xDSL standards that use
//			this attribute) (1-byte)
//
//		Maximum Bit Error Ratio
//			(R,-W, setbycreate) (mandatory for standards that use this attribute) (1-byte)
//
//		Minimum Impulse Noise Protection 8_Khz
//			Minimum impulse noise protection 8-kHz: The INPmin8 attribute specifies the minimum INP for the
//			bearer channel if it is transported over DMT symbols with a subcarrier spacing of 8.625-kHz. It
//			is only valid for [ITUT-G.993.2]. INP is expressed in DMT symbols with a subcarrier spacing of
//			8.625-kHz. It can take any integer value from 0 (default) to 16, inclusive. (R, W) (mandatory
//			for [ITUT-G.993.2]) (1 byte)
//
//		Maximum Delay Variation
//			Maximum delay variation: The DVMAX attribute specifies the maximum value for delay variation
//			allowed in an OLR procedure. Its value ranges from 1 (0.1-ms) to 254 (25.4-ms). The special
//			value 255 specifies that no delay variation bound is imposed. (R, W) (optional: used by
//			[ITUT-G.993.2]) (1 byte)
//
//		Channel Initialization Policy Selection
//			Channel initialization policy selection: The CIPOLICY attribute specifies the policy to
//			determine transceiver configuration at initialization. Valid values are 0..1, as defined in the
//			Recommendations that use this attribute. (R,-W) (optional) (1-byte)
//
//		Minimum Sos Bit Rate Downstream
//			Minimum SOS bit rate downstream: The MIN-SOS-BR-ds attribute specifies the minimum net data rate
//			required for a valid SOS request in the downstream direction. The value is coded as an unsigned
//			integer representing the data rate as a multiple of 8-kbit/s. (R,-W) (optional) (4-bytes)
//
//		Minimum Sos Bit Rate Upstream
//			Minimum SOS bit rate upstream: The MIN-SOS-BR-us attribute specifies the minimum net data rate
//			required for a valid SOS request in the upstream direction. The value is coded as an unsigned
//			integer representing the data rate as a multiple of 8-kbit/s. (R,-W) (optional) (4-bytes)
//
type XdslChannelConfigurationProfile struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	xdslchannelconfigurationprofileBME = &ManagedEntityDefinition{
		Name:    "XdslChannelConfigurationProfile",
		ClassID: 107,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFFFE,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, false, 0),
			1:  Uint32Field("MinimumDataRate", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 1),
			2:  Uint32Field("MaximumDataRate", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 2),
			3:  ByteField("RateAdaptationRatio", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, true, false, 3),
			4:  ByteField("MaximumInterleavingDelay", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 4),
			5:  Uint32Field("DataRateThresholdUpshift", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 5),
			6:  Uint32Field("DataRateThresholdDownshift", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 6),
			7:  Uint32Field("MinimumReservedDataRate", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, true, false, 7),
			8:  Uint32Field("MinimumDataRateInLowPowerState", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 8),
			9:  ByteField("MinimumImpulseNoiseProtection", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 9),
			10: ByteField("MaximumBitErrorRatio", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 10),
			11: ByteField("MinimumImpulseNoiseProtection8Khz", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 11),
			12: ByteField("MaximumDelayVariation", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 12),
			13: ByteField("ChannelInitializationPolicySelection", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 13),
			14: Uint32Field("MinimumSosBitRateDownstream", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 14),
			15: Uint32Field("MinimumSosBitRateUpstream", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 15),
		},
	}
}

// NewXdslChannelConfigurationProfile (class ID 107 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewXdslChannelConfigurationProfile(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*xdslchannelconfigurationprofileBME, params...)
}
