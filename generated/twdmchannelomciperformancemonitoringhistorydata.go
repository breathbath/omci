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

const TwdmChannelOmciPerformanceMonitoringHistoryDataClassId ClassID = ClassID(452)

var twdmchannelomciperformancemonitoringhistorydataBME *ManagedEntityDefinition

// TwdmChannelOmciPerformanceMonitoringHistoryData (class ID #452) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type TwdmChannelOmciPerformanceMonitoringHistoryData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	twdmchannelomciperformancemonitoringhistorydataBME = &ManagedEntityDefinition{
		Name:    "TwdmChannelOmciPerformanceMonitoringHistoryData",
		ClassID: 452,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			GetCurrentData,
			Set,
		),
		AllowedAttributeMask: 0XF800,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, false, 0),
			1: ByteField("IntervalEndTime", 0, mapset.NewSetWith(Read), false, false, false, false, 1),
			2: Uint16Field("ThresholdData12Id", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 2),
			3: Uint32Field("OmciBaselineMessageCount", 0, mapset.NewSetWith(Read), false, false, false, false, 3),
			4: Uint32Field("OmciExtendedMessageCount", 0, mapset.NewSetWith(Read), false, false, false, false, 4),
			5: Uint32Field("OmciMicErrorCount", 0, mapset.NewSetWith(Read), false, false, false, false, 5),
		},
	}
}

// NewTwdmChannelOmciPerformanceMonitoringHistoryData (class ID 452 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewTwdmChannelOmciPerformanceMonitoringHistoryData(params ...ParamData) (*ManagedEntity, error) {
	return NewManagedEntity(twdmchannelomciperformancemonitoringhistorydataBME, params...)
}
