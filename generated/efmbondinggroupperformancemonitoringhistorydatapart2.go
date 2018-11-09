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

// EfmBondingGroupPerformanceMonitoringHistoryDataPart2 (class ID 422 defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type EfmBondingGroupPerformanceMonitoringHistoryDataPart2 struct {
	BaseManagedEntityDefinition
}

// NewEfmBondingGroupPerformanceMonitoringHistoryDataPart2 (class ID 422 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewEfmBondingGroupPerformanceMonitoringHistoryDataPart2(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "EfmBondingGroupPerformanceMonitoringHistoryDataPart2",
		ClassID:  422,
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
			ByteField("IntervalEndTime", 0, Read),
			Uint16Field("ThresholdData12Id", 0, Read|Write|SetByCreate),
			Uint32Field("RxUnicastFrames", 0, Read),
			Uint32Field("TxUnicastFrames", 0, Read),
			Uint64Field("RxUnicastBytes", 0, Read),
			Uint64Field("TxUnicastBytes", 0, Read),
			Uint32Field("RxBroadcastFrames", 0, Read),
			Uint32Field("TxBroadcastFrames", 0, Read),
			Uint64Field("RxBroadcastBytes", 0, Read),
			Uint64Field("TxBroadcastBytes", 0, Read),
			Uint32Field("RxMulticastFrames", 0, Read),
			Uint32Field("TxMulticastFrames", 0, Read),
			Uint64Field("RxMulticastBytes", 0, Read),
			Uint64Field("TxMulticastBytes", 0, Read),
		},
	}
	entity.computeAttributeMask()
	return &EfmBondingGroupPerformanceMonitoringHistoryDataPart2{entity}, nil
}
