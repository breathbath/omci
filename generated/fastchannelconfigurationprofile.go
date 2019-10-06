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

const FastChannelConfigurationProfileClassId ClassID = ClassID(432)

var fastchannelconfigurationprofileBME *ManagedEntityDefinition

// FastChannelConfigurationProfile (class ID #432)
//	This ME contains the FAST channel configuration profile for an xDSL UNI. An instance of this ME
//	is created and deleted by the OLT.
//
//	Relationships
//		An instance of this ME may be associated with zero or more instances of the PPTP xDSL UNI part
//		1.
//
//	Attributes
//		Maximum Net Data Rate Maxndr
//			Maximum net data rate (MAXNDR): This attribute specifies the value of the maximum net data rate.
//			See clause 11.4.2.2 of [ITU-T G.9701]. Valid values range from 0 (0-kbit/s) to 4294967295
//			(2^32-1-kbit/s). See clause 7.2.1.1 of [ITUT-G.997.2]. (R, W) (mandatory) (4 bytes)
//
//		Minimum Expected Throughput Minetr
//			Minimum expected throughput (MINETR): This attribute specifies the value of the minimum expected
//			throughput. See clause 11.4.2.1 of [ITU-T G.9701]. Valid values range from 0 (0-kbit/s) to
//			4294967295 (2^32-1-kbit/s). See clause-7.2.1.2 of [ITU-T G.997.2]. (R, W) (mandatory) (4 bytes)
//
//		Maximum Gamma Data Rate Maxgdr
//			Maximum gamma data rate (MAXGDR): This attribute specifies the maximum value of the GDR (see
//			clause 7.11.1.3). The GDR shall not exceed MAXGDR at the start of showtime and during showtime.
//			Valid values range from 0 (0-kbit/s) to 4294967295 (2^32-1-kbit/s). See clause 7.2.1.3 of [ITU-T
//			G.997.2]. (R, W) (mandatory) (4 bytes)
//
//		Minimum Gamma Data Rate Mingdr
//			Minimum gamma data rate (MINGDR): This attribute specifies the minimum value of the GDR (see
//			clause 7.11.1.3). The GDR may be lower than MINGDR. If the GDR is lower than MINGDR at
//			initialization or when GDR becomes lower than MINGDR during showtime, a TCA occurs. Valid values
//			range from 0 (0-kbit/s) to 4294967295 (2^32-1-kbit/s). See clause 7.2.1.4 of [ITU-T G.997.2].
//			(R, W) (mandatory) (4 bytes)
//
//		Maximum Delay Delaymax
//			Maximum delay (DELAYMAX): This attribute specifies the maximum allowed delay for retransmission.
//			See clause 9.8 of [ITU-T G.9701]. The ITUT-G.9701 control parameter delay_max is set to the same
//			value as the maximum delay. See clause 11.4.2.3 of [ITU-T G.9701]. Valid values range from 4
//			(1-ms) to 252 (63-ms) in steps of 0.25-ms. See clause 7.2.2.1 of [ITUT-G.997.2]. (R, W)
//			(mandatory) (4 bytes)
//
//		Minimum Impulse Noise Protection Against Shine Inpmin_Shine
//			Minimum impulse noise protection against SHINE (INPMIN_SHINE): This attribute specifies the
//			minimum INP against SHINE. See clause 9.8 of [ITU-T G.9701]. The ITU-T G.9701 control parameter
//			INP_min_shine is set to the same value as the minimum INP against SHINE. See clause 11.4.2.4 of
//			[ITU-T G.9701]. Valid values range from 0 to 520 (520 symbol periods). See clause 7.2.2.2 of
//			[ITUT G.997.2]. (R, W) (mandatory) (2-bytes)
//
//		Shine Ratio Shineratio
//			SHINE ratio (SHINERATIO): This attribute specifies the SHINE ratio that is used in the
//			definition of the expected throughput rate (ETR). See clause 9.8 of [ITUT-G.9701]. The ITU-T
//			G.9701 control parameter SHINEratio is set to the same value as the SHINE ratio. See clause
//			11.4.2.5 of [ITU-T G.9701]. The value is expressed in units of 0.001, Valid values range from 0
//			to 100 (0.01) in steps of 0.001. See clause 7.2.2.3 of [ITU-T G.997.2]. (R, W) (mandatory)
//			(1-byte)
//
//		Minimum Impulse Noise Protection Against Rein Inpmin_Rein
//			Minimum impulse noise protection against REIN (INPMIN_REIN): This attribute specifies the
//			minimum INP against REIN. See clause 9.8 of [ITU-T G.9701]. The ITU-T G.9701 control parameter
//			INP_min_rein is set to the same value as the minimum INP against REIN. See clause 11.4.2.6 of
//			[ITU-T G.9701]. Valid values range from 0 to 63 (63-symbol periods). See clause 7.2.2.4 of
//			[ITU-T G.997.2]. (R, W) (mandatory) (1 byte)
//
//		Rein Inter_Arrival Time Iat_Rein
//			(R, W) (mandatory) (1 byte)
//
//		Minimum Reed_Solomon Rfec_Nfec Ratio Rnratio
//			Minimum Reed-Solomon RFEC/NFEC ratio (RNRATIO): This attribute specifies the minimal required
//			ratio, RFEC/NFEC, of Reed-Solomon code parameters. The ITU-T G.9701 control parameter rnratio is
//			set to the same value as the minimum Reed-Solomon RFEC/NFEC ratio. See clause 11.4.2.8 of
//			[ITUT-G.9701]. The value is expressed in units of 1/32, Valid values range from 0 to 8 (1/4).
//			See clause 7.2.2.6 of [ITU-T G.997.2]. (R, W) (mandatory) (1 byte)
//
//		Rtx_Tc Testmode Rtx_Testmode
//			RTX-TC testmode (RTX_TESTMODE): This Boolean attribute specifies whether the retransmission test
//			mode defined in clause 9.8.3.1.2 [ITU-T G.9701] is enabled-(true) or disabled (disabled). See
//			clause 7.2.2.7 of [ITU-T G.997.2]. (R,-W) (optional) (1 byte)
//
type FastChannelConfigurationProfile struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	fastchannelconfigurationprofileBME = &ManagedEntityDefinition{
		Name:    "FastChannelConfigurationProfile",
		ClassID: 432,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFFC0,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint32Field("MaximumNetDataRateMaxndr", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 0),
			1:  Uint32Field("MinimumExpectedThroughputMinetr", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 1),
			2:  Uint32Field("MaximumGammaDataRateMaxgdr", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 2),
			3:  Uint32Field("MinimumGammaDataRateMingdr", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 3),
			4:  Uint32Field("MaximumDelayDelaymax", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 4),
			5:  Uint16Field("MinimumImpulseNoiseProtectionAgainstShineInpminShine", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 5),
			6:  ByteField("ShineRatioShineratio", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 6),
			7:  ByteField("MinimumImpulseNoiseProtectionAgainstReinInpminRein", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 7),
			8:  ByteField("ReinInterArrivalTimeIatRein", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 8),
			9:  ByteField("MinimumReedSolomonRfecNfecRatioRnratio", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 9),
			10: ByteField("RtxTcTestmodeRtxTestmode", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 10),
		},
	}
}

// NewFastChannelConfigurationProfile (class ID 432 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewFastChannelConfigurationProfile(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*fastchannelconfigurationprofileBME, params...)
}
