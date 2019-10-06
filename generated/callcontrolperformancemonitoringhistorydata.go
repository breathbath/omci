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

const CallControlPerformanceMonitoringHistoryDataClassId ClassID = ClassID(140)

var callcontrolperformancemonitoringhistorydataBME *ManagedEntityDefinition

// CallControlPerformanceMonitoringHistoryData (class ID #140)
//	This ME collects PM data related to the call control channel. Instances of this ME are created
//	and deleted by the OLT.
//
//	For a complete discussion of generic PM architecture, refer to clause I.4.
//
//	Relationships
//		An instance of this ME is associated with an instance of the PPTP POTS UNI ME.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an
//			identical ID, this ME is implicitly linked to an instance of the PPTP POTS UNI. (R, setbycreate)
//			(mandatory) (2-bytes)
//
//		Interval End Time
//			Interval end time: This attribute identifies the most recently finished 15-min interval. (R)
//			(mandatory) (1-byte)
//
//		Threshold Data 1_2 Id
//			Threshold data 1/2 ID: This attribute points to an instance of the threshold data 1 ME that
//			contains PM threshold values. Since no threshold value attribute number exceeds 7, a threshold
//			data 2 ME is optional. (R,-W, setbycreate) (mandatory) (2-bytes)
//
//		Call Setup Failures
//			Call setup failures: This attribute counts call set-up failures. (R) (mandatory) (4-bytes)
//
//		Call Setup Timer
//			Call setup timer: This attribute is a high water-mark that records the longest duration of a
//			single call set-up detected during this interval. Time is measured in milliseconds from the time
//			an initial set-up was requested by the subscriber until the time at which a response was
//			provided to the subscriber in the form of busy tone, audible ring tone, etc. (R) (mandatory)
//			(4-bytes)
//
//		Call Terminate Failures
//			Call terminate failures: This attribute counts the number of calls that were terminated with
//			cause. (R) (mandatory) (4-bytes)
//
//		Analog Port Releases
//			Analog port releases: This attribute counts the number of analogue port releases without
//			dialling detected (abandoned calls). (R) (mandatory) (4-bytes)
//
//		Analog Port Off_Hook Timer
//			Analog port off-hook timer: This attribute is a high water-mark that records the longest period
//			of a single off-hook detected on the analogue port. Time is measured in milliseconds. (R)
//			(mandatory) (4-bytes)
//
type CallControlPerformanceMonitoringHistoryData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	callcontrolperformancemonitoringhistorydataBME = &ManagedEntityDefinition{
		Name:    "CallControlPerformanceMonitoringHistoryData",
		ClassID: 140,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFE00,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, false, 0),
			1: ByteField("IntervalEndTime", 0, mapset.NewSetWith(Read), false, false, false, false, 1),
			2: Uint16Field("ThresholdData12Id", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 2),
			3: Uint32Field("CallSetupFailures", 0, mapset.NewSetWith(Read), false, false, false, false, 3),
			4: Uint32Field("CallSetupTimer", 0, mapset.NewSetWith(Read), false, false, false, false, 4),
			5: Uint32Field("CallTerminateFailures", 0, mapset.NewSetWith(Read), false, false, false, false, 5),
			6: Uint32Field("AnalogPortReleases", 0, mapset.NewSetWith(Read), false, false, false, false, 6),
			7: Uint32Field("AnalogPortOffHookTimer", 0, mapset.NewSetWith(Read), false, false, false, false, 7),
		},
	}
}

// NewCallControlPerformanceMonitoringHistoryData (class ID 140 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewCallControlPerformanceMonitoringHistoryData(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*callcontrolperformancemonitoringhistorydataBME, params...)
}
