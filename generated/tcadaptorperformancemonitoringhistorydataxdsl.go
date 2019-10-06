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

const TcAdaptorPerformanceMonitoringHistoryDataXdslClassId ClassID = ClassID(116)

var tcadaptorperformancemonitoringhistorydataxdslBME *ManagedEntityDefinition

// TcAdaptorPerformanceMonitoringHistoryDataXdsl (class ID #116)
//	This ME collects PM data of an xTUC to xTUR ATM data path. Instances of this ME are created and
//	deleted by the OLT.
//
//	For a complete discussion of generic PM architecture, refer to clause I.4.
//
//	Relationships
//		An instance of this ME is associated with an xDSL UNI.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID:	This attribute uniquely identifies each instance of this ME. Through an
//			identical ID, this ME is implicitly linked to an instance of the PPTP xDSL UNI part 1. (R)
//			(mandatory) (2-bytes)
//
//		Interval End Time
//			Interval end time: This attribute identifies the most recently finished 15-min interval. (R)
//			(mandatory) (1-byte)
//
//		Threshold Data 1 _2 Id
//			Threshold data1/2 ID: This attribute points to an instance of the threshold data1 ME that
//			contains PM threshold values. Since no threshold value attribute number exceeds 7, a threshold
//			data 2 ME is optional. (R,-W, setbycreate) (mandatory) (2-bytes)
//
//		Near_End Hec Violation Count
//			Near-end HEC violation count: This attribute counts near-end HEC anomalies in the ATM data path.
//			(R) (mandatory) (2-bytes)
//
//		Near_End Delineated Total Cell Count Cd P
//			Near-end delineated total cell count (CDP): This attribute counts the total number of cells
//			passed through the cell delineation and HEC function process operating on the ATM data path
//			while in the SYNC state. (R) (mandatory) (4-bytes)
//
//		Near_End User Total Cell Count Cu_P
//			Near-end user total cell count(CU-P): This attribute counts the total number of cells in the ATM
//			data path delivered at the V-C interface. (R) (mandatory) (4-bytes)
//
//		Near_End Idle Cell Bit Error Count
//			Near-end idle cell bit error count: This attribute counts cells with bit errors in the ATM data
//			path idle payload received at the near end. (R) (mandatory) (2-bytes)
//
//		Far_End Hec Violation Count
//			Far-end HEC violation count: This attribute counts far-end HEC anomalies in the ATM data path.
//			(R) (mandatory) (2-bytes)
//
//		Far_End Delineated Total Cell Count Cd_Pfe
//			Far-end delineated total cell count (CD-PFE): This attribute counts the total number of cells
//			passed through the cell delineation process and HEC function operating on the ATM data path
//			while in the SYNC state. (R) (mandatory) (4-bytes)
//
//		Far_End User Total Cell Count Cu_Pfe
//			Far-end user total cell count (CU-PFE): This attribute counts the total number of cells in the
//			ATM data path delivered at the T-R interface. (R) (mandatory) (4-bytes)
//
//		Far_End Idle Cell Bit Error Count
//			Far-end idle cell bit error count: This attribute counts cells with bit errors in the ATM data
//			path idle payload received at the far end. (R) (mandatory) (2-bytes)
//
type TcAdaptorPerformanceMonitoringHistoryDataXdsl struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	tcadaptorperformancemonitoringhistorydataxdslBME = &ManagedEntityDefinition{
		Name:    "TcAdaptorPerformanceMonitoringHistoryDataXdsl",
		ClassID: 116,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFFC0,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read), false, false, false, false, 0),
			1:  ByteField("IntervalEndTime", 0, mapset.NewSetWith(Read), false, false, false, false, 1),
			2:  Uint16Field("ThresholdData12Id", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 2),
			3:  Uint16Field("NearEndHecViolationCount", 0, mapset.NewSetWith(Read), false, false, false, false, 3),
			4:  Uint32Field("NearEndDelineatedTotalCellCountCdP", 0, mapset.NewSetWith(Read), false, false, false, false, 4),
			5:  Uint32Field("NearEndUserTotalCellCountCuP", 0, mapset.NewSetWith(Read), false, false, false, false, 5),
			6:  Uint16Field("NearEndIdleCellBitErrorCount", 0, mapset.NewSetWith(Read), false, false, false, false, 6),
			7:  Uint16Field("FarEndHecViolationCount", 0, mapset.NewSetWith(Read), false, false, false, false, 7),
			8:  Uint32Field("FarEndDelineatedTotalCellCountCdPfe", 0, mapset.NewSetWith(Read), false, false, false, false, 8),
			9:  Uint32Field("FarEndUserTotalCellCountCuPfe", 0, mapset.NewSetWith(Read), false, false, false, false, 9),
			10: Uint16Field("FarEndIdleCellBitErrorCount", 0, mapset.NewSetWith(Read), false, false, false, false, 10),
		},
	}
}

// NewTcAdaptorPerformanceMonitoringHistoryDataXdsl (class ID 116 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewTcAdaptorPerformanceMonitoringHistoryDataXdsl(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*tcadaptorperformancemonitoringhistorydataxdslBME, params...)
}
