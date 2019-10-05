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

const FastXtuRPerformanceMonitoringHistoryDataClassId ClassID = ClassID(438)

var fastxturperformancemonitoringhistorydataBME *ManagedEntityDefinition

// FastXtuRPerformanceMonitoringHistoryData (class ID #438)
//	This ME collects PM data of the xTU C to xTU R path as seen from the xTU-R. Instances of this ME
//	are created and deleted by the OLT. For a complete discussion of generic PM architecture, refer
//	to clause I.4.
//
//	Relationships
//		An instance of this ME is associated with an xDSL UNI.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an
//			identical ID, this ME is implicitly linked to an instance of the PPTP xDSL UNI part 1. (R, set-
//			by-create) (mandatory) (2 bytes)
//
//		Interval End Time
//			Interval end time: This attribute identifies the most recently finished 15 min interval. (R)
//			(mandatory) (1 byte)
//
//		Threshold Data 1_2 Id
//			Threshold data 1/2 ID: This attribute points to an instance of the threshold data 1 and 2 MEs
//			that contain PM threshold values. (R, W, set-by-create) (mandatory) (2 bytes)
//
//		Successful Fra Counter
//			Successful FRA counter: This attribute counts the successful FRA primitives (success_FRA). The
//			successful FRA primitive (success_FRA) is defined in clause 11.3.1.6 of [ITU-T G.9701]. See
//			clause 7.7.22 of [ITU-T G.997.2] (R) (mandatory) (4 bytes)
//
//		Successful Rpa Counter
//			Successful RPA counter: This attribute counts the successful RPA primitives (success_RPA). The
//			successful RPA primitive (success_RPA) is defined in clause 11.3.1.6 of [ITU-T G.9701]. See
//			clause 7.7.23 of [ITU-T G.997.2] (R) (optional) (4 bytes)
//
type FastXtuRPerformanceMonitoringHistoryData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	fastxturperformancemonitoringhistorydataBME = &ManagedEntityDefinition{
		Name:    "FastXtuRPerformanceMonitoringHistoryData",
		ClassID: 438,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XF000,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, false, 0),
			1: ByteField("IntervalEndTime", 0, mapset.NewSetWith(Read), false, false, false, false, 1),
			2: Uint16Field("ThresholdData12Id", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 2),
			3: Uint32Field("SuccessfulFraCounter", 0, mapset.NewSetWith(Read), false, false, false, false, 3),
			4: Uint32Field("SuccessfulRpaCounter", 0, mapset.NewSetWith(Read), false, false, true, false, 4),
		},
	}
}

// NewFastXtuRPerformanceMonitoringHistoryData (class ID 438 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewFastXtuRPerformanceMonitoringHistoryData(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(fastxturperformancemonitoringhistorydataBME, params...)
}
