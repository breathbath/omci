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

// XdslChannelConfigurationProfile (class ID #107) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
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
func NewXdslChannelConfigurationProfile(params ...ParamData) (*ManagedEntity, error) {
	return NewManagedEntity(xdslchannelconfigurationprofileBME, params...)
}
