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

// TwdmChannelTuningPerformanceMonitoringHistoryDataPart2 (class ID 450 defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type TwdmChannelTuningPerformanceMonitoringHistoryDataPart2 struct {
	BaseManagedEntityDefinition
}

// NewTwdmChannelTuningPerformanceMonitoringHistoryDataPart2 (class ID 450 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewTwdmChannelTuningPerformanceMonitoringHistoryDataPart2(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "TwdmChannelTuningPerformanceMonitoringHistoryDataPart2",
		ClassID:  450,
		EntityID: eid,
		MessageTypes: []MsgType{
			Set,
			Get,
			Create,
			GetCurrentData,
			Delete,
		},
		AttributeMask: 0,
		Attributes: []*AttributeDefinition{
			Uint16Field("ManagedEntityId", 0, Read|SetByCreate),
			ByteField("IntervalEndTime", 0, Read),
			Uint16Field("ThresholdData12Id", 0, Read|Write|SetByCreate),
			Uint32Field("TuningControlRequestsRejectedDsAlbl", 0, Read),
			Uint32Field("TuningControlRequestsRejectedDsVoid", 0, Read),
			Uint32Field("TuningControlRequestsRejectedDsPart", 0, Read),
			Uint32Field("TuningControlRequestsRejectedDsTunr", 0, Read),
			Uint32Field("TuningControlRequestsRejectedDsLnrt", 0, Read),
			Uint32Field("TuningControlRequestsRejectedDsLncd", 0, Read),
			Uint32Field("TuningControlRequestsRejectedUsAlbl", 0, Read),
			Uint32Field("TuningControlRequestsRejectedUsVoid", 0, Read),
			Uint32Field("TuningControlRequestsRejectedUsTunr", 0, Read),
			Uint32Field("TuningControlRequestsRejectedUsClbr", 0, Read),
			Uint32Field("TuningControlRequestsRejectedUsLktp", 0, Read),
			Uint32Field("TuningControlRequestsRejectedUsLnrt", 0, Read),
			Uint32Field("TuningControlRequestsRejectedUsLncd", 0, Read),
		},
	}
	entity.computeAttributeMask()
	return &TwdmChannelTuningPerformanceMonitoringHistoryDataPart2{entity}, nil
}
