/*
 * Copyright (c) 2018 - present.  Boling Consulting Solutions (bcsw.net)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
package omci_test

import (
	. "github.com/cboling/omci"
	. "github.com/cboling/omci/generated"
	"github.com/google/gopacket"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var allMsgTypes = [...]MsgType{
	Create,
	Delete,
	Set,
	Get,
	GetAllAlarms,
	GetAllAlarmsNext,
	MibUpload,
	MibUploadNext,
	MibReset,
	AlarmNotification,
	AttributeValueChange,
	Test,
	StartSoftwareDownload,
	DownloadSection,
	EndSoftwareDownload,
	ActivateSoftware,
	CommitSoftware,
	SynchronizeTime,
	Reboot,
	GetNext,
	TestResult,
	GetCurrentData,
	SetTable}

var allResults = [...]Results{
	Success,
	ProcessingError,
	NotSupported,
	ParameterError,
	UnknownEntity,
	UnknownInstance,
	DeviceBusy,
	InstanceExists}

// MibResetRequestTest tests decode/encode of a MIB Reset Request
func TestMsgTypeStrings(t *testing.T) {
	for _, msg := range allMsgTypes {
		strMsg := msg.String()
		assert.NotEqual(t, len(strMsg), 0)
	}
}

func TestResultsStrings(t *testing.T) {
	for _, code := range allResults {
		strMsg := code.String()
		assert.NotEqual(t, len(strMsg), 0)
	}
}

// TestOmciDecode will test for proper error checking of things that
// are invalid at the OMCI decode layer
func TestOmciDecode(t *testing.T) {
	// TID = 0 on autonomous ONU notifications only.  Get packet back but ErrorLayer()
	// returns non-nil
	tidZeroOnNonNotification := "0000440A010C01000400800003010000" +
		"00000000000000000000000000000000" +
		"000000000000000000000028"

	data, err := stringToPacket(tidZeroOnNonNotification)
	assert.NoError(t, err)
	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)
	assert.NotNil(t, packet.ErrorLayer())

	// Only Baseline and Extended Message types allowed
	invalidMessageType := "000C440F010C01000400800003010000" +
		"00000000000000000000000000000000" +
		"000000000000000000000028"

	data, err = stringToPacket(invalidMessageType)
	assert.NoError(t, err)
	packet = gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)
	assert.NotNil(t, packet.ErrorLayer())

	// Bad baseline message length
	badBaselineMsgLength := "000C440A010C01000400800003010000" +
		"00000000000000000000000000000000" +
		"000000000000000000000029"

	data, err = stringToPacket(badBaselineMsgLength)
	assert.NoError(t, err)
	packet = gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)
	assert.NotNil(t, packet.ErrorLayer())

	// Bad extended message length
	badExtendedMsgLength := "000C440B010C010000290400800003010000" +
		"00000000000000000000000000000000" +
		"00000000000000000000"

	data, err = stringToPacket(badExtendedMsgLength)
	assert.NoError(t, err)
	packet = gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)
	assert.NotNil(t, packet.ErrorLayer())

	// Huge extended message length
	hugeExtendedMsgLength := "000C440B010C010007BD0400800003010000" +
		"00000000000000000000000000000000" +
		"00000000000000000000"

	data, err = stringToPacket(hugeExtendedMsgLength)
	assert.NoError(t, err)
	packet = gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)
	assert.NotNil(t, packet.ErrorLayer())
	// fmt.Println(packet.ErrorLayer())
}

// TestOmciSerialization will test for proper error checking of things that
// are invalid at the OMCI layer
func TestOmciSerialization(t *testing.T) {
	goodMessage := "000C440A010C0100040080000301000000000000000000000000000000000000000000000000000000000028"

	omciLayerDefaults := &OMCI{
		TransactionID: 0x0c,
		MessageType:   CreateRequestType,
		// DeviceIdentifier: BaselineIdent,		// Optional, defaults to Baseline
		// Length:           0x28,				// Optional, defaults to 40 octets
	}
	omciLayerFixed := &OMCI{
		TransactionID:    0x0c,
		MessageType:      CreateRequestType,
		DeviceIdentifier: BaselineIdent,
		Length:           0x28,
	}
	request := &CreateRequest{
		MeBasePacket: MeBasePacket{
			EntityClass:    GemPortNetworkCtpClassId,
			EntityInstance: uint16(0x100),
		},
		Attributes: AttributeValueMap{
			"PortId":                                       0x400,
			"TContPointer":                                 0x8000,
			"Direction":                                    3,
			"TrafficManagementPointerForUpstream":          0x100,
			"TrafficDescriptorProfilePointerForUpstream":   0,
			"PriorityQueuePointerForDownStream":            0,
			"TrafficDescriptorProfilePointerForDownstream": 0,
			"EncryptionKeyRing":                            0,
		},
	}
	// Test serialization back to former string (using defaults in the message parts)
	var options gopacket.SerializeOptions
	options.FixLengths = true

	buffer := gopacket.NewSerializeBuffer()
	err := gopacket.SerializeLayers(buffer, options, omciLayerDefaults, request)
	assert.NoError(t, err)

	outgoingPacket := buffer.Bytes()
	reconstituted := packetToString(outgoingPacket)
	assert.Equal(t, strings.ToLower(goodMessage), reconstituted)

	// Test serialization back to former string (using explicit values in the message parts)
	buffer = gopacket.NewSerializeBuffer()
	err = gopacket.SerializeLayers(buffer, options, omciLayerFixed, request)
	assert.NoError(t, err)

	outgoingPacket = buffer.Bytes()
	reconstituted = packetToString(outgoingPacket)
	assert.Equal(t, strings.ToLower(goodMessage), reconstituted)
}

func TestCreateRequestDecode(t *testing.T) {
	goodMessage := "000C440A010C01000400800003010000" +
		"00000000000000000000000000000000" +
		"000000000000000000000028"
	data, err := stringToPacket(goodMessage)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.TransactionID, uint16(0xc))
	assert.Equal(t, omciMsg.MessageType, CreateRequestType)
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeCreateRequest)
	assert.NotNil(t, msgLayer)

	request, ok2 := msgLayer.(*CreateRequest)
	assert.True(t, ok2)
	assert.Equal(t, request.EntityClass, GemPortNetworkCtpClassId)
	assert.Equal(t, request.EntityInstance, uint16(0x100))

	attributes := request.Attributes
	assert.NotNil(t, attributes)

	// As this is a create request, gather up all set-by-create attributes
	// make sure we got them all, and nothing else
	meDefinition, err := LoadManagedEntityDefinition(request.EntityClass)
	assert.Nil(t, err)

	attrDefs := *meDefinition.GetAttributeDefinitions()

	sbcMask := getSbcMask(meDefinition)
	for index := uint(1); index < uint(len(attrDefs)); index++ {
		attrName := attrDefs[index].GetName()

		if sbcMask&uint16(1<<(uint)(16-index)) != 0 {
			_, ok3 := attributes[attrName]
			assert.True(t, ok3)
		} else {
			_, ok3 := attributes[attrName]
			assert.False(t, ok3)
		}
		//fmt.Printf("Name: %v, Value: %v\n", attrName, attributes[attrName])
	}
	// Test serialization back to former string
	var options gopacket.SerializeOptions
	options.FixLengths = true

	buffer := gopacket.NewSerializeBuffer()
	err = gopacket.SerializeLayers(buffer, options, omciMsg, request)
	assert.NoError(t, err)

	outgoingPacket := buffer.Bytes()
	reconstituted := packetToString(outgoingPacket)
	assert.Equal(t, strings.ToLower(goodMessage), reconstituted)
}

func TestCreateRequestSerialize(t *testing.T) {
	goodMessage := "000C440A010C0100040080000301000000000000000000000000000000000000000000000000000000000028"

	omciLayer := &OMCI{
		TransactionID: 0x0c,
		MessageType:   CreateRequestType,
		// DeviceIdentifier: omci.BaselineIdent,		// Optional, defaults to Baseline
		// Length:           0x28,						// Optional, defaults to 40 octets
	}
	request := &CreateRequest{
		MeBasePacket: MeBasePacket{
			EntityClass:    GemPortNetworkCtpClassId,
			EntityInstance: uint16(0x100),
		},
		Attributes: AttributeValueMap{
			"PortId":                                       0x400,
			"TContPointer":                                 0x8000,
			"Direction":                                    3,
			"TrafficManagementPointerForUpstream":          0x100,
			"TrafficDescriptorProfilePointerForUpstream":   0,
			"PriorityQueuePointerForDownStream":            0,
			"TrafficDescriptorProfilePointerForDownstream": 0,
			"EncryptionKeyRing":                            0,
		},
	}
	// Test serialization back to former string
	var options gopacket.SerializeOptions
	options.FixLengths = true

	buffer := gopacket.NewSerializeBuffer()
	err := gopacket.SerializeLayers(buffer, options, omciLayer, request)
	assert.NoError(t, err)

	outgoingPacket := buffer.Bytes()
	reconstituted := packetToString(outgoingPacket)
	assert.Equal(t, strings.ToLower(goodMessage), reconstituted)
}

func TestCreateResponseDecode(t *testing.T) {
	goodMessage := "0108240a002d0900000000000000000000000000000000000000000000000000000000000000000000000028"
	data, err := stringToPacket(goodMessage)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.MessageType, CreateResponseType)
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeCreateResponse)
	assert.NotNil(t, msgLayer)

	response, ok2 := msgLayer.(*CreateResponse)
	assert.True(t, ok2)
	assert.NotNil(t, response)
}

func TestCreateResponseSerialize(t *testing.T) {

	goodMessage := "0108240a002d0900000000000000000000000000000000000000000000000000000000000000000000000028"

	omciLayer := &OMCI{
		TransactionID: 0x0108,
		MessageType:   CreateResponseType,
		// DeviceIdentifier: omci.BaselineIdent,		// Optional, defaults to Baseline
		// Length:           0x28,						// Optional, defaults to 40 octets
	}
	request := &CreateRequest{
		MeBasePacket: MeBasePacket{
			EntityClass:    MacBridgeServiceProfileClassId,
			EntityInstance: uint16(0x900),
		},
		Attributes: AttributeValueMap{
			"SpanningTreeInd":            0,
			"LearningInd":                0,
			"PortBridgingInd":            0,
			"Priority":                   0,
			"MaxAge":                     0,
			"HelloTime":                  0,
			"ForwardDelay":               0,
			"UnknownMacAddressDiscard":   0,
			"MacLearningDepth":           0,
			"DynamicFilteringAgeingTime": 0,
		},
	}
	// Test serialization back to former string
	var options gopacket.SerializeOptions
	options.FixLengths = true

	buffer := gopacket.NewSerializeBuffer()
	err := gopacket.SerializeLayers(buffer, options, omciLayer, request)
	assert.NoError(t, err)

	outgoingPacket := buffer.Bytes()
	reconstituted := packetToString(outgoingPacket)
	assert.Equal(t, strings.ToLower(goodMessage), reconstituted)
}

func TestDeleteRequestDecode(t *testing.T) {
	goodMessage := "0211460a00ab0202000000000000000000000000000000000000000000000000000000000000000000000028"
	data, err := stringToPacket(goodMessage)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.MessageType, DeleteRequestType)
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeDeleteRequest)

	assert.NotNil(t, msgLayer)

	request, ok2 := msgLayer.(*DeleteRequest)
	assert.True(t, ok2)
	assert.NotNil(t, request)
}

func TestDeleteRequestSerialize(t *testing.T) {
	// TODO:Implement
}

func TestDeleteResponseDecode(t *testing.T) {
	goodMessage := "0211260a00ab0202000000000000000000000000000000000000000000000000000000000000000000000028013437fb"
	data, err := stringToPacket(goodMessage)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.MessageType, DeleteResponseType)
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeDeleteResponse)

	assert.NotNil(t, msgLayer)

	response, ok2 := msgLayer.(*DeleteResponse)
	assert.True(t, ok2)
	assert.NotNil(t, response)
}

func TestDeleteResponseSerialize(t *testing.T) {
	// TODO:Implement
}

func TestSetRequestDecode(t *testing.T) {
	goodMessage := "0107480a01000000020000000000000000000000000000000000000000000000000000000000000000000028"
	data, err := stringToPacket(goodMessage)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.MessageType, SetRequestType)
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeSetRequest)
	assert.NotNil(t, msgLayer)

	request, ok2 := msgLayer.(*SetRequest)
	assert.True(t, ok2)
	assert.NotNil(t, request)
}

func TestSetRequestSerialize(t *testing.T) {
	// TODO:Implement
}

func TestSetResponseDecode(t *testing.T) {
	goodMessage := "0107280a01000000000000000000000000000000000000000000000000000000000000000000000000000028"
	data, err := stringToPacket(goodMessage)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.MessageType, SetResponseType)
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeSetResponse)
	assert.NotNil(t, msgLayer)

	response, ok2 := msgLayer.(*SetResponse)
	assert.True(t, ok2)
	assert.NotNil(t, response)
}

func TestSetResponseSerialize(t *testing.T) {
	// TODO:Implement
}

func TestGetRequestDecode(t *testing.T) {
	goodMessage := "035e490a01070000004400000000000000000000000000000000000000000000000000000000000000000028"
	data, err := stringToPacket(goodMessage)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.MessageType, GetRequestType)
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeGetRequest)
	assert.NotNil(t, msgLayer)

	request, ok2 := msgLayer.(*GetRequest)
	assert.True(t, ok2)
	assert.NotNil(t, request)
}

func TestGetRequestSerialize(t *testing.T) {
	// TODO:Implement
}

func TestGetResponseDecode(t *testing.T) {
	goodMessage := "035e290a01070000000044dbcb05f10000000000000000000000000000000000000000000000000000000028"
	data, err := stringToPacket(goodMessage)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.MessageType, GetResponseType)
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeGetResponse)
	assert.NotNil(t, msgLayer)

	response, ok2 := msgLayer.(*GetResponse)
	assert.True(t, ok2)
	assert.NotNil(t, response)
}

func TestGetResponseSerialize(t *testing.T) {
	// TODO:Implement
}

func TestGetAllAlarmsRequestDecode(t *testing.T) {
	goodMessage := "04454b0a00020000000000000000000000000000000000000000000000000000000000000000000000000028"
	data, err := stringToPacket(goodMessage)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.MessageType, GetAllAlarmsRequestType)
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeGetAllAlarmsRequest)
	assert.NotNil(t, msgLayer)

	request, ok2 := msgLayer.(*GetAllAlarmsRequest)
	assert.True(t, ok2)
	assert.NotNil(t, request)
}

func TestGetAllAlarmsSerialize(t *testing.T) {
	// TODO:Implement
}

func TestGetAllAlarmsResponseDecode(t *testing.T) {
	goodMessage := "04452b0a00020000000000000000000000000000000000000000000000000000000000000000000000000028"
	data, err := stringToPacket(goodMessage)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.MessageType, GetAllAlarmsResponseType)
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeGetAllAlarmsResponse)
	assert.NotNil(t, msgLayer)

	response, ok2 := msgLayer.(*GetAllAlarmsResponse)
	assert.True(t, ok2)
	assert.NotNil(t, response)
}

func TestGetAllAlarmsResponseSerialize(t *testing.T) {
	// TODO:Implement
}

func TestGetAllAlarmsNextRequestDecode(t *testing.T) {
	goodMessage := "02344c0a00020000000000000000000000000000000000000000000000000000000000000000000000000028"

	data, err := stringToPacket(goodMessage)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.MessageType, GetAllAlarmsNextRequestType)
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeGetAllAlarmsNextRequest)
	assert.NotNil(t, msgLayer)

	request, ok2 := msgLayer.(*GetAllAlarmsNextRequest)
	assert.True(t, ok2)
	assert.NotNil(t, request)
}

func TestGetAllAlarmsNextRequestSerialize(t *testing.T) {
	// TODO:Implement
}

func TestGetAllAlarmsNextResponseDecode(t *testing.T) {
	goodMessage := "02342c0a00020000000b01028000000000000000000000000000000000000000000000000000000000000028f040fc87"
	data, err := stringToPacket(goodMessage)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.MessageType, GetAllAlarmsNextResponseType)
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeGetAllAlarmsNextResponse)
	assert.NotNil(t, msgLayer)

	response, ok2 := msgLayer.(*GetAllAlarmsNextResponse)
	assert.True(t, ok2)
	assert.NotNil(t, response)
}

func TestGetAllAlarmsNextResponseSerialize(t *testing.T) {
	// TODO:Implement
}

func TestMibUploadRequestDecode(t *testing.T) {
	goodMessage := "03604d0a00020000000000000000000000000000000000000000000000000000000000000000000000000028"
	data, err := stringToPacket(goodMessage)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.MessageType, MibUploadRequestType)
	assert.Equal(t, omciMsg.Length, uint16(40))
	msgLayer := packet.Layer(LayerTypeMibUploadRequest)

	assert.NotNil(t, msgLayer)

	request, ok2 := msgLayer.(*MibUploadRequest)
	assert.True(t, ok2)
	assert.NotNil(t, request)
}

func TestMibUploadRequestSerialize(t *testing.T) {
	// TODO:Implement
}
func TestMibUploadResponse(t *testing.T) {
	goodMessage := "03602d0a00020000011200000000000000000000000000000000000000000000000000000000000000000028"
	data, err := stringToPacket(goodMessage)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.MessageType, MibUploadResponseType)
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeMibUploadResponse)
	assert.NotNil(t, msgLayer)

	response, ok2 := msgLayer.(*MibUploadResponse)
	assert.True(t, ok2)
	assert.NotNil(t, response)
}

func TestMibUploadResponseSerialize(t *testing.T) {
	// TODO:Implement
}
func TestMibUploadNextRequestDecode(t *testing.T) {
	goodMessage := "02864e0a00020000003a00000000000000000000000000000000000000000000000000000000000000000028"
	data, err := stringToPacket(goodMessage)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.MessageType, MibUploadNextRequestType)
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeMibUploadNextRequest)
	assert.NotNil(t, msgLayer)

	request, ok2 := msgLayer.(*MibUploadNextRequest)
	assert.True(t, ok2)
	assert.NotNil(t, request)
}

func TestMibUploadNextRequestSerialize(t *testing.T) {
	// TODO:Implement
}

func TestMibUploadNextResponse(t *testing.T) {
	goodMessage := "02862e0a0002000001150000fff0000000000000000000010100000000010000000000000000000000000028"
	data, err := stringToPacket(goodMessage)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.MessageType, MibUploadNextResponseType)
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeMibUploadNextResponse)
	assert.NotNil(t, msgLayer)

	response, ok2 := msgLayer.(*MibUploadNextResponse)
	assert.True(t, ok2)
	assert.NotNil(t, response)
}

func TestMibUploadNextResponseSerialize(t *testing.T) {
	// TODO:Implement
}

func TestMibResetRequestDecode(t *testing.T) {
	goodMessage := "00014F0A00020000000000000000000000000000000000000000000000000000000000000000000000000028"
	data, err := stringToPacket(goodMessage)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.MessageType, MibResetRequestType)
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeMibResetRequest)
	assert.NotNil(t, msgLayer)

	request, ok2 := msgLayer.(*MibResetRequest)
	assert.True(t, ok2)
	assert.NotNil(t, request)
}

func TestMibResetRequestSerialize(t *testing.T) {
	goodMessage := "00014F0A00020000000000000000000000000000000000000000000000000000000000000000000000000028"

	omciLayer := &OMCI{
		TransactionID: 0x01,
		MessageType:   MibResetRequestType,
		// DeviceIdentifier: omci.BaselineIdent,		// Optional, defaults to Baseline
		// Length:           0x28,						// Optional, defaults to 40 octets
	}
	request := &MibResetRequest{
		MeBasePacket: MeBasePacket{
			EntityClass: OnuDataClassId,
			// Default Instance ID is 0
		},
	}
	// Test serialization back to former string
	var options gopacket.SerializeOptions
	options.FixLengths = true

	buffer := gopacket.NewSerializeBuffer()
	err := gopacket.SerializeLayers(buffer, options, omciLayer, request)
	assert.NoError(t, err)

	outgoingPacket := buffer.Bytes()
	reconstituted := packetToString(outgoingPacket)
	assert.Equal(t, strings.ToLower(goodMessage), reconstituted)
}

func TestMibResetResponseDecode(t *testing.T) {
	goodMessage := "00012F0A00020000000000000000000000000000000000000000000000000000000000000000000000000028"
	data, err := stringToPacket(goodMessage)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.MessageType, MibResetResponseType)
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeMibResetResponse)

	assert.NotNil(t, msgLayer)

	response, ok2 := msgLayer.(*MibResetResponse)
	assert.True(t, ok2)
	assert.NotNil(t, response)
}

func TestMibResetResponseSerialize(t *testing.T) {
	goodMessage := "00012F0A00020000000000000000000000000000000000000000000000000000000000000000000000000028"

	omciLayer := &OMCI{
		TransactionID: 0x01,
		MessageType:   MibResetResponseType,
		// DeviceIdentifier: omci.BaselineIdent,		// Optional, defaults to Baseline
		// Length:           0x28,						// Optional, defaults to 40 octets
	}
	request := &MibResetResponse{
		MeBasePacket: MeBasePacket{
			EntityClass: OnuDataClassId,
			// Default Instance ID is 0
		},
	}
	// Test serialization back to former string
	var options gopacket.SerializeOptions
	options.FixLengths = true

	buffer := gopacket.NewSerializeBuffer()
	err := gopacket.SerializeLayers(buffer, options, omciLayer, request)
	assert.NoError(t, err)

	outgoingPacket := buffer.Bytes()
	reconstituted := packetToString(outgoingPacket)
	assert.Equal(t, strings.ToLower(goodMessage), reconstituted)
}

// TODO: Create request/response tests for all of the following types
//Test,
//StartSoftwareDownload, reqMsg := "0000530a0007000113000f424001000100000000000000000000000000000000000000000000000000000028"
//DownloadSection, reqMsg := '0000140a00070001083534363836393733323036393733323036313230373436353733373400000000000028'
//EndSoftwareDownload, reqMsg := '0000550a00070001ff92a226000f424001000100000000000000000000000000000000000000000000000028'
//ActivateSoftware, reqMsg := '0000560a00070001000000000000000000000000000000000000000000000000000000000000000000000028'
//CommitSoftware, reqMsg := '0000570a00070001000000000000000000000000000000000000000000000000000000000000000000000028'

func TestSynchronizeTimeRequestDecode(t *testing.T) {
	goodMessage := "0109580a0100000007e20c0001301b0000000000000000000000000000000000000000000000000000000028"
	data, err := stringToPacket(goodMessage)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.MessageType, SynchronizeTimeRequestType)
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeSynchronizeTimeRequest)
	assert.NotNil(t, msgLayer)

	request, ok2 := msgLayer.(*SynchronizeTimeRequest)
	assert.True(t, ok2)
	assert.NotNil(t, request)
}

func TestSynchronizeTimeRequestSerialize(t *testing.T) {
	// TODO:Implement
}

func TestSynchronizeTimeResponseEncode(t *testing.T) {
	goodMessage := "0109380a01000000000000000000000000000000000000000000000000000000000000000000000000000028"
	data, err := stringToPacket(goodMessage)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.MessageType, SynchronizeTimeResponseType)
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeSynchronizeTimeResponse)
	assert.NotNil(t, msgLayer)

	response, ok2 := msgLayer.(*SynchronizeTimeResponse)
	assert.True(t, ok2)
	assert.NotNil(t, response)
}

func TestSynchronizeTimeResponseSerialize(t *testing.T) {
	// TODO:Implement
}

// TODO: Create request/response tests for all of the following types
//Reboot, msgRequest := '0001590a01000000000000000000000000000000000000000000000000000000000000000000000000000028'
//			msgResponse: '023c390a01000000000000000000000000000000000000000000000000000000000000000000000000000028005999e3'
//GetNext,
//GetCurrentData,
//SetTable}

// TODO: Create notification tests for all of the following types
//AlarmNotification,  (TODO: Include alarm bitmap tests as well)

func TestAttributeValueChangeDecode(t *testing.T) {
	goodMessage := "0000110a0007000080004d4c2d33363236000000000000002020202020202020202020202020202000000028"
	data, err := stringToPacket(goodMessage)
	assert.NoError(t, err)

	packet := gopacket.NewPacket(data, LayerTypeOMCI, gopacket.NoCopy)
	assert.NotNil(t, packet)

	omciLayer := packet.Layer(LayerTypeOMCI)
	assert.NotNil(t, packet)

	omciMsg, ok := omciLayer.(*OMCI)
	assert.True(t, ok)
	assert.Equal(t, omciMsg.MessageType, AttributeValueChangeType)
	assert.Equal(t, omciMsg.Length, uint16(40))

	msgLayer := packet.Layer(LayerTypeAttributeValueChange)
	assert.NotNil(t, msgLayer)

	request, ok2 := msgLayer.(*AttributeValueChangeMsg)
	assert.True(t, ok2)
	assert.NotNil(t, request)
}

func TestAttributeValueChangeSerialize(t *testing.T) {
	// TODO:Implement
}

// TODO: Create notification tests for all of the following types
//TestResult,
