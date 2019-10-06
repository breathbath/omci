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

const EthernetFrameExtendedPmClassId ClassID = ClassID(334)

var ethernetframeextendedpmBME *ManagedEntityDefinition

// EthernetFrameExtendedPm (class ID #334)
//	This ME collects some of the PM data at a point where an Ethernet flow can be observed. It is
//	based on the Etherstats group of [IETF RFC 2819]. Instances of this ME are created and deleted
//	by the OLT. References to received frames are to be interpreted as the number of frames entering
//	the monitoring point in the direction specified by the control block.
//
//	For a complete discussion of generic PM architecture, refer to clause I.4.
//
//	Relationships
//		An instance of this ME may be associated with an instance of an ME at any Ethernet interface
//		within the ONU. The specific ME is identified in the control block attribute.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. To facilitate
//			discovery, the identification of instances sequentially starting with 1 is encouraged. (R,
//			setbycreate) (mandatory) (2 bytes)
//
//		Interval End Time
//			Interval end time: This attribute identifies the most recently finished 15-min interval. If
//			continuous accumulation is enabled in the control block, this attribute is not used and has the
//			fixed value 0. (R) (mandatory) (1 byte)
//
//		Control Block
//			(R, W, setbycreate) (mandatory) (16 bytes)
//
//		Drop Events
//			Drop events:	The total number of events in which frames were dropped due to a lack of resources.
//			This is not necessarily the number of frames dropped; it is the number of times this event was
//			detected. (R) (mandatory) (4 bytes)
//
//		Octets
//			Octets:	The total number of octets received, including those in bad frames, excluding framing
//			bits, but including FCS. (R) (mandatory) (4 bytes)
//
//		Frames
//			Frames:	The total number of frames received, including bad frames, broadcast frames and
//			multicast frames. (R) (mandatory) (4 bytes)
//
//		Broadcast Frames
//			Broadcast frames: The total number of received good frames directed to the broadcast address.
//			This does not include multicast frames. (R) (mandatory) (4 bytes)
//
//		Multicast Frames
//			Multicast frames: The total number of received good frames directed to a multicast address. This
//			does not include broadcast frames. (R) (mandatory) (4 bytes)
//
//		Crc Errored Frames
//			CRC errored frames: The total number of frames received that had a length (excluding framing
//			bits, but including FCS octets) of between 64 and 1518 octets, inclusive, but had either a bad
//			FCS with an integral number of octets (FCS error) or a bad FCS with a non-integral number of
//			octets (alignment error). (R) (mandatory) (4 bytes)
//
//		Undersize Frames
//			Undersize frames: The total number of frames received that were less than 64 octets long but
//			were otherwise well formed (excluding framing bits, but including FCS octets). (R) (mandatory)
//			(4 bytes)
//
//		Oversize Frames
//			Oversize frames: The total number of frames received that were longer than 1518 octets
//			(excluding framing bits, but including FCS octets) and were otherwise well formed. (R)
//			(mandatory) (4 bytes)
//
//		Frames 64 Octets
//			Frames 64 octets: The total number of received frames (including bad frames) that were 64-octets
//			long, excluding framing bits but including FCS. (R) (mandatory) (4-bytes)
//
//		Frames 65 To 127 Octets
//			Frames 65 to 127 octets: The total number of received frames (including bad frames) that were
//			65..127 octets long, excluding framing bits but including FCS. (R) (mandatory) (4 bytes)
//
//		Frames 128 To 255 Octets
//			Frames 128 to 255 octets: The total number of frames (including bad frames) received that were
//			128..255 octets long, excluding framing bits but including FCS. (R) (mandatory) (4 bytes)
//
//		Frames 256 To 511 Octets
//			Frames 256 to 511 octets: The total number of frames (including bad frames) received that were
//			256..511 octets long, excluding framing bits but including FCS. (R) (mandatory) (4 bytes)
//
//		Frames 512 To 1 023 Octets
//			Frames 512 to 1-023 octets: The total number of frames (including bad frames) received that were
//			512..1-023 octets long, excluding framing bits but including FCS. (R) (mandatory) (4 bytes)
//
//		Frames 1024 To 1518 Octets
//			Frames 1024 to 1518 octets: The total number of frames (including bad frames) received that were
//			1024..1518 octets long, excluding framing bits but including FCS. (R) (mandatory) (4 bytes)
//
type EthernetFrameExtendedPm struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	ethernetframeextendedpmBME = &ManagedEntityDefinition{
		Name:    "EthernetFrameExtendedPm",
		ClassID: 334,
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
			2:  MultiByteField("ControlBlock", 16, nil, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, false, 2),
			3:  Uint32Field("DropEvents", 0, mapset.NewSetWith(Read), false, false, false, false, 3),
			4:  Uint32Field("Octets", 0, mapset.NewSetWith(Read), false, false, false, false, 4),
			5:  Uint32Field("Frames", 0, mapset.NewSetWith(Read), false, false, false, false, 5),
			6:  Uint32Field("BroadcastFrames", 0, mapset.NewSetWith(Read), false, false, false, false, 6),
			7:  Uint32Field("MulticastFrames", 0, mapset.NewSetWith(Read), false, false, false, false, 7),
			8:  Uint32Field("CrcErroredFrames", 0, mapset.NewSetWith(Read), false, false, false, false, 8),
			9:  Uint32Field("UndersizeFrames", 0, mapset.NewSetWith(Read), false, false, false, false, 9),
			10: Uint32Field("OversizeFrames", 0, mapset.NewSetWith(Read), false, false, false, false, 10),
			11: Uint32Field("Frames64Octets", 0, mapset.NewSetWith(Read), false, false, false, false, 11),
			12: Uint32Field("Frames65To127Octets", 0, mapset.NewSetWith(Read), false, false, false, false, 12),
			13: Uint32Field("Frames128To255Octets", 0, mapset.NewSetWith(Read), false, false, false, false, 13),
			14: Uint32Field("Frames256To511Octets", 0, mapset.NewSetWith(Read), false, false, false, false, 14),
			15: Uint32Field("Frames512To1023Octets", 0, mapset.NewSetWith(Read), false, false, false, false, 15),
			16: Uint32Field("Frames1024To1518Octets", 0, mapset.NewSetWith(Read), false, false, false, false, 16),
		},
	}
}

// NewEthernetFrameExtendedPm (class ID 334 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewEthernetFrameExtendedPm(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*ethernetframeextendedpmBME, params...)
}
