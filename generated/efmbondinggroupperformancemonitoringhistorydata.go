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

const EfmBondingGroupPerformanceMonitoringHistoryDataClassId uint16 = 421

var efmbondinggroupperformancemonitoringhistorydataBME *BaseManagedEntityDefinition

// EfmBondingGroupPerformanceMonitoringHistoryData (class ID #421) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type EfmBondingGroupPerformanceMonitoringHistoryData struct {
	BaseManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	efmbondinggroupperformancemonitoringhistorydataBME := &BaseManagedEntityDefinition{
		Name:     "EfmBondingGroupPerformanceMonitoringHistoryData",
		ClassID:  421,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFFF0,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, Read|SetByCreate, false, false, false),
			1:  ByteField("IntervalEndTime", 0, Read, false, false, false),
			2:  Uint16Field("ThresholdData12Id", 0, Read|SetByCreate|Write, false, false, false),
			3:  Uint32Field("RxBadFragments", 0, Read, false, false, false),
			4:  Uint32Field("RxLostFragments", 0, Read, false, false, false),
			5:  Uint32Field("RxLostStarts", 0, Read, false, false, false),
			6:  Uint32Field("RxLostEnds", 0, Read, false, false, false),
			7:  Uint32Field("RxFrames", 0, Read, false, false, false),
			8:  Uint32Field("TxFrames", 0, Read, false, false, false),
			9:  Uint64Field("RxBytes", 0, Read, false, false, false),
			10: Uint64Field("TxBytes", 0, Read, false, false, false),
			11: Uint32Field("TxDiscardedFrames", 0, Read, false, false, false),
			12: Uint32Field("TxDiscardedBytes", 0, Read, false, false, false),
		},
	}
}

// NewEfmBondingGroupPerformanceMonitoringHistoryData (class ID 421 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewEfmBondingGroupPerformanceMonitoringHistoryData(params ...ParamData) (IManagedEntity, error) {
	entity := &ManagedEntity {
	    Definition: efmbondinggroupperformancemonitoringhistorydataBME,
	    Attributes: make(map[string]interface{}),
	}
	if err := entity.setAttributes(params...); err != nil {
	    return nil, err
	}
	return entity, nil
}
