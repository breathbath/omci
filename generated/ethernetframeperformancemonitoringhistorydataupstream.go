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

const EthernetFramePerformanceMonitoringHistoryDataUpstreamClassId ClassID = ClassID(322)

var ethernetframeperformancemonitoringhistorydataupstreamBME *ManagedEntityDefinition

// EthernetFramePerformanceMonitoringHistoryDataUpstream (class ID #322)
//	This ME collects PM data associated with upstream Ethernet frame delivery. It is based on the
//	Etherstats group of [IETF RFC 2819]. Instances of this ME are created and deleted by the OLT.
//
//	For a complete discussion of generic PM architecture, refer to clause I.4.
//
//	NOTE 1 - Implementers are encouraged to consider the Ethernet frame extended PM ME defined in
//	clause-9.3.32, which collects the same counters in a more generalized way.
//
//	Relationships
//		An instance of this ME is associated with an instance of a MAC bridge port configuration data.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an
//			identical ID, this ME is implicitly linked to an instance of a MAC bridge port configuration
//			data. (R, setbycreate) (mandatory) (2-bytes)
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
//		Drop Events
//			Drop events:	The total number of events in which packets were dropped due to a lack of
//			resources. This is not necessarily the number of packets dropped; it is the number of times this
//			event was detected. (R) (mandatory) (4-bytes)
//
//		Octets
//			Octets:	The total number of upstream octets received, including those in bad packets, excluding
//			framing bits, but including FCS. (R) (mandatory) (4-bytes)
//
//		Packets
//			Packets:	The total number of upstream packets received, including bad packets, broadcast packets
//			and multicast packets. (R) (mandatory) (4-bytes)
//
//		Broadcast Packets
//			Broadcast packets: The total number of upstream good packets received that were directed to the
//			broadcast address. This does not include multicast packets. (R) (mandatory) (4-bytes)
//
//		Multicast Packets
//			Multicast packets: The total number of upstream good packets received that were directed to a
//			multicast address. This does not include broadcast packets. (R) (mandatory) (4-bytes)
//
//		Crc Errored Packets
//			CRC errored packets: The total number of upstream packets received that had a length (excluding
//			framing bits, but including FCS octets) of between 64-octets and 1518-octets, inclusive, but had
//			either a bad FCS with an integral number of octets (FCS error) or a bad FCS with a non-integral
//			number of octets (alignment error). (R) (mandatory) (4-bytes)
//
//		Undersize Packets
//			Undersize packets: The total number of upstream packets received that were less than 64-octets
//			long, but were otherwise well formed (excluding framing bits, but including FCS). (R)
//			(mandatory) (4-bytes)
//
//		Oversize Packets
//			NOTE 2 - If 2-000-byte Ethernet frames are supported, counts in this performance parameter are
//			not necessarily errors.
//
//		Packets 64 Octets
//			Packets 64 octets: The total number of upstream received packets (including bad packets) that
//			were 64-octets long, excluding framing bits but including FCS. (R) (mandatory) (4-bytes)
//
//		Packets 65 To 127 Octets
//			Packets 65 to 127 octets: The total number of upstream received packets (including bad packets)
//			that were 65..127 octets long, excluding framing bits but including FCS. (R) (mandatory)
//			(4-bytes)
//
//		Packets 128 To 255 Octets
//			Packets 128 to 255 octets: The total number of upstream packets (including bad packets) received
//			that were 128..255 octets long, excluding framing bits but including FCS. (R) (mandatory)
//			(4-bytes)
//
//		Packets 256 To 511 Octets
//			Packets 256 to 511 octets: The total number of upstream packets (including bad packets) received
//			that were 256..511 octets long, excluding framing bits but including FCS. (R) (mandatory)
//			(4-bytes)
//
//		Packets 512 To 1023 Octets
//			Packets 512 to 1023 octets: The total number of upstream packets (including bad packets)
//			received that were 512..1-023 octets long, excluding framing bits but including FCS. (R)
//			(mandatory) (4-bytes)
//
//		Packets 1024 To 1518 Octets
//			Packets 1024 to 1518 octets: The total number of upstream packets (including bad packets)
//			received that were 1024..1518 octets long, excluding framing bits, but including FCS. (R)
//			(mandatory) (4-bytes)
//
type EthernetFramePerformanceMonitoringHistoryDataUpstream struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	ethernetframeperformancemonitoringhistorydataupstreamBME = &ManagedEntityDefinition{
		Name:    "EthernetFramePerformanceMonitoringHistoryDataUpstream",
		ClassID: 322,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFFFF,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, false, 0),
			1:  ByteField("IntervalEndTime", 0, mapset.NewSetWith(Read), false, false, false, false, 1),
			2:  Uint16Field("ThresholdData12Id", 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 2),
			3:  Uint32Field("DropEvents", 0, mapset.NewSetWith(Read), false, false, false, false, 3),
			4:  Uint32Field("Octets", 0, mapset.NewSetWith(Read), false, false, false, false, 4),
			5:  Uint32Field("Packets", 0, mapset.NewSetWith(Read), false, false, false, false, 5),
			6:  Uint32Field("BroadcastPackets", 0, mapset.NewSetWith(Read), false, false, false, false, 6),
			7:  Uint32Field("MulticastPackets", 0, mapset.NewSetWith(Read), false, false, false, false, 7),
			8:  Uint32Field("CrcErroredPackets", 0, mapset.NewSetWith(Read), false, false, false, false, 8),
			9:  Uint32Field("UndersizePackets", 0, mapset.NewSetWith(Read), false, false, false, false, 9),
			10: Uint32Field("OversizePackets", 0, mapset.NewSetWith(Read), false, false, false, false, 10),
			11: Uint32Field("Packets64Octets", 0, mapset.NewSetWith(Read), false, false, false, false, 11),
			12: Uint32Field("Packets65To127Octets", 0, mapset.NewSetWith(Read), false, false, false, false, 12),
			13: Uint32Field("Packets128To255Octets", 0, mapset.NewSetWith(Read), false, false, false, false, 13),
			14: Uint32Field("Packets256To511Octets", 0, mapset.NewSetWith(Read), false, false, false, false, 14),
			15: Uint32Field("Packets512To1023Octets", 0, mapset.NewSetWith(Read), false, false, false, false, 15),
			16: Uint32Field("Packets1024To1518Octets", 0, mapset.NewSetWith(Read), false, false, false, false, 16),
		},
	}
}

// NewEthernetFramePerformanceMonitoringHistoryDataUpstream (class ID 322 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewEthernetFramePerformanceMonitoringHistoryDataUpstream(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*ethernetframeperformancemonitoringhistorydataupstreamBME, params...)
}
