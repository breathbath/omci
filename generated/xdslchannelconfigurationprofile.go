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
 *              https://github.com/cboling/OMCI-parser
 */
package generated

// XdslChannelConfigurationProfile (class ID 107 defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type XdslChannelConfigurationProfile struct {
	BaseManagedEntityDefinition
}

// NewXdslChannelConfigurationProfile (class ID 107 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewXdslChannelConfigurationProfile(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "XdslChannelConfigurationProfile",
		ClassID:  107,
		EntityID: eid,
		MessageTypes: []MsgType{
			Set,
			Get,
			Create,
			Delete,
		},
		AttributeMask: 0,
		Attributes: []*AttributeDefinition{
			Uint16Field("ManagedEntityId", 0, Read|SetByCreate),
			Uint32Field("MinimumDataRate", 0, Read|Write|SetByCreate),
			Uint32Field("MaximumDataRate", 0, Read|Write|SetByCreate),
			ByteField("RateAdaptationRatio", 0, Read|Write|SetByCreate),
			ByteField("MaximumInterleavingDelay", 0, Read|Write|SetByCreate),
			Uint32Field("DataRateThresholdUpshift", 0, Read|Write|SetByCreate),
			Uint32Field("DataRateThresholdDownshift", 0, Read|Write|SetByCreate),
			Uint32Field("MinimumReservedDataRate", 0, Read|Write|SetByCreate),
			Uint32Field("MinimumDataRateInLowPowerState", 0, Read|Write|SetByCreate),
			ByteField("MinimumImpulseNoiseProtection", 0, Read|Write|SetByCreate),
			ByteField("MaximumBitErrorRatio", 0, Read|Write|SetByCreate),
			ByteField("MinimumImpulseNoiseProtection8Khz", 0, Read|Write),
			ByteField("MaximumDelayVariation", 0, Read|Write),
			ByteField("ChannelInitializationPolicySelection", 0, Read|Write),
			Uint32Field("MinimumSosBitRateDownstream", 0, Read|Write),
			Uint32Field("MinimumSosBitRateUpstream", 0, Read|Write),
		},
	}
	entity.computeAttributeMask()
	return &XdslChannelConfigurationProfile{entity}, nil
}
