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

// EnhancedFecPerformanceMonitoringHistoryDataClassID is the 16-bit ID for the OMCI
// Managed entity Enhanced FEC performance monitoring history data
const EnhancedFecPerformanceMonitoringHistoryDataClassID ClassID = ClassID(453)

var enhancedfecperformancemonitoringhistorydataBME *ManagedEntityDefinition

// EnhancedFecPerformanceMonitoringHistoryData (class ID #453)
//	This ME collects PM data associated with PON downstream FEC counters for XGS-PON and subsequent
//	ITU-T PON systems. Instances of this ME are created and deleted by the OLT.
//
//	For a complete discussion of generic PM architecture, refer to clause I.4.
//
//	Relationships
//		An instance of this ME is associated with an instance of the ANI-G ME or an instance of the TWDM
//		channel ME.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an
//			identical ID, this ME is implicitly linked to an instance of the ANI-G or a TWDM channel. (R,
//			setbycreate) (mandatory) (2-bytes)
//
//		Interval End Time
//			Interval end time: This attribute identifies the most recently finished 15-min interval. (R)
//			(mandatory) (1-byte)
//
//		Threshold Data 64 Bit Id
//			Threshold data 64-bit ID: This attribute points to an instance of the threshold data 64-bit ME
//			that contains PM threshold values. (R,-W, setbycreate) (mandatory) (2-bytes)
//
//		Corrected Bytes
//			Corrected bytes: This attribute counts the number of bytes that were corrected by the FEC
//			function. (R) (mandatory) (8-bytes)
//
//		Corrected Code Words
//			Corrected code words: This attribute counts the code words that were corrected by the FEC
//			function. (R) (mandatory) (8-bytes)
//
//		Uncorrectable Code Words
//			Uncorrectable code words: This attribute counts errored code words that could not be corrected
//			by the FEC function. (R) (mandatory) (8-bytes)
//
//		Total Code Words
//			Total code words: This attribute counts the total received code words. (R) (mandatory) (8-bytes)
//
//		Fec Seconds
//			FEC seconds:	This attribute counts seconds during which there was an FEC anomaly. (R)
//			(mandatory) (2-bytes)
//
type EnhancedFecPerformanceMonitoringHistoryData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	enhancedfecperformancemonitoringhistorydataBME = &ManagedEntityDefinition{
		Name:    "EnhancedFecPerformanceMonitoringHistoryData",
		ClassID: 453,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0xfe00,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, false, 0),
			1: ByteField("IntervalEndTime", 0, mapset.NewSetWith(Read), false, false, false, false, 1),
			2: Uint16Field("ThresholdData64BitId", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 2),
			3: Uint64Field("CorrectedBytes", 0, mapset.NewSetWith(Read), false, true, false, false, 3),
			4: Uint64Field("CorrectedCodeWords", 0, mapset.NewSetWith(Read), false, true, false, false, 4),
			5: Uint64Field("UncorrectableCodeWords", 0, mapset.NewSetWith(Read), false, true, false, false, 5),
			6: Uint64Field("TotalCodeWords", 0, mapset.NewSetWith(Read), false, true, false, false, 6),
			7: Uint16Field("FecSeconds", 0, mapset.NewSetWith(Read), false, true, false, false, 7),
		},
		Access:  CreatedByOlt,
		Support: UnknownSupport,
	}
}

// NewEnhancedFecPerformanceMonitoringHistoryData (class ID 453) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewEnhancedFecPerformanceMonitoringHistoryData(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*enhancedfecperformancemonitoringhistorydataBME, params...)
}
