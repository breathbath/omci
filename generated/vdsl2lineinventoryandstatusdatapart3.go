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

// Vdsl2LineInventoryAndStatusDataPart3ClassID is the 16-bit ID for the OMCI
// Managed entity VDSL2 line inventory and status data part 3
const Vdsl2LineInventoryAndStatusDataPart3ClassID ClassID = ClassID(170)

var vdsl2lineinventoryandstatusdatapart3BME *ManagedEntityDefinition

// Vdsl2LineInventoryAndStatusDataPart3 (class ID #170)
//	This ME extends the other xDSL line inventory and status data MEs with attributes specific to
//	VDSL2. This ME contains per-band attributes for both directions. These same attributes are
//	defined in the xDSL line inventory and status data part 2 ME, but only for a single band. [ITUT
//	G.993.2] allows for VDSL2 to have as many as five bands upstream and as many as five bands
//	downstream.
//
//	Relationships
//		This is one of the status data MEs associated with an xDSL UNI. It is required only if VDSL2 is
//		supported by the PPTP. The ONU automatically creates or deletes an instance of this ME upon
//		creation or deletion of a PPTP xDSL UNI part 1 that supports these attributes.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an
//			identical ID, this ME is implicitly linked to an instance of the PPTP xDSL UNI part 1 ME. (R)
//			(mandatory) (2-bytes)
//
//		Downstream Line Attenuation Per Band
//			Downstream line attenuation per band: The LATNds attribute is defined per usable band. It is the
//			squared magnitude of the channel characteristics function, H(f), averaged over this band, and
//			measured during loop diagnostic mode and initialization. The exact definition is included in the
//			relevant xDSL Recommendation. The upstream line attenuation per band ranges from 0 (0.0-dB) to
//			1270 (+127.0-dB). The special value 0xFFFF indicates that the line attenuation per band is out
//			of the range to be represented. (R) (mandatory) (3-bands * 2-bytes-=-6-bytes)
//
//		Upstream Line Attenuation Per Band
//			Upstream line attenuation per band: The LATNus attribute is defined per usable band. It is the
//			squared magnitude of the channel characteristics function H(f) averaged over this band, and
//			measured during loop diagnostic mode and initialization. The exact definition is included in the
//			relevant xDSL Recommendation. The upstream line attenuation per band ranges from 0 (0.0-dB) to
//			1270 (+127.0-dB). The special value 0xFFFF indicates that line attenuation per band is out of
//			range to be represented. (R) (mandatory) (4-bands * 2-bytes-=-8-bytes)
//
//		Downstream Signal Attenuation Per Band
//			NOTE 1 - During showtime, only a subset of the subcarriers may be transmitted by the xTU-C, as
//			compared to loop diagnostic mode and initialization. Therefore, the downstream signal
//			attenuation value during showtime may be significantly lower than the downstream signal
//			attenuation value during loop diagnostic mode and initialization.
//
//		Upstream Signal Attenuation Per Band
//			NOTE 2 - During showtime, only a subset of the subcarriers may be transmitted by the xTU-R, as
//			compared to loop diagnostic mode and initialization. Therefore, the upstream signal attenuation
//			value during showtime may be significantly lower than the upstream signal attenuation value
//			during loop diagnostic mode and initialization.
//
//		Downstream Snr Margin Per Band
//			Downstream SNR margin per band: The SNRMpbds attribute is defined per usable band. The
//			downstream SNR margin per band is the maximum increase of noise power received at the xTU-R,
//			such that the BER requirements are met for all downstream bearer channels. Each array value
//			ranges from 0 (-64.0-dB) to 1270 (+63.0-dB). The special value 0xFFFF indicates that the
//			attribute is out of range to be represented. (R) (mandatory) (3 bands * 2-bytes-=-6-bytes)
//
//		Upstream Snr Margin Per Band
//			Upstream SNR margin per band: The SNRMpbus attribute is defined per usable band. The upstream
//			SNR margin per band is the maximum increase of noise power received at the xTU-C, such that the
//			BER requirements are met for all upstream bearer channels. Each array value ranges from 0
//			(-64.0-dB) to 1270 (+63.0-dB). The special value 0xFFFF indicates that the attribute is out of
//			range to be represented. (R) (mandatory) (4 bands * 2-bytes-= 8-bytes)
//
type Vdsl2LineInventoryAndStatusDataPart3 struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	vdsl2lineinventoryandstatusdatapart3BME = &ManagedEntityDefinition{
		Name:    "Vdsl2LineInventoryAndStatusDataPart3",
		ClassID: 170,
		MessageTypes: mapset.NewSetWith(
			Get,
		),
		AllowedAttributeMask: 0xfc00,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read), false, false, false, false, 0),
			1: MultiByteField("DownstreamLineAttenuationPerBand", 3, nil, mapset.NewSetWith(Read), false, false, false, false, 1),
			2: Uint32Field("UpstreamLineAttenuationPerBand", 0, mapset.NewSetWith(Read), false, false, false, false, 2),
			3: MultiByteField("DownstreamSignalAttenuationPerBand", 3, nil, mapset.NewSetWith(Read), false, false, false, false, 3),
			4: Uint32Field("UpstreamSignalAttenuationPerBand", 0, mapset.NewSetWith(Read), false, false, false, false, 4),
			5: MultiByteField("DownstreamSnrMarginPerBand", 3, nil, mapset.NewSetWith(Read), false, false, false, false, 5),
			6: Uint32Field("UpstreamSnrMarginPerBand", 0, mapset.NewSetWith(Read), false, false, false, false, 6),
		},
		Access:  CreatedByOnu,
		Support: UnknownSupport,
	}
}

// NewVdsl2LineInventoryAndStatusDataPart3 (class ID 170) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewVdsl2LineInventoryAndStatusDataPart3(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*vdsl2lineinventoryandstatusdatapart3BME, params...)
}
