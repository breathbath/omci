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
package omci

import (
	"errors"
	"fmt"
	me "github.com/cboling/omci/generated"
	"github.com/deckarep/golang-set"
	"github.com/google/gopacket"
)

var encoderMap map[MessageType]func(*me.ManagedEntity, options) (gopacket.SerializableLayer, error)

func init() {
	encoderMap = make(map[MessageType]func(*me.ManagedEntity, options) (gopacket.SerializableLayer, error))

	encoderMap[CreateRequestType] = CreateRequestFrame
	encoderMap[DeleteRequestType] = DeleteRequestFrame
	encoderMap[SetRequestType] = SetRequestFrame
	encoderMap[GetRequestType] = GetRequestFrame
	encoderMap[GetAllAlarmsRequestType] = GetAllAlarmsRequestFrame
	encoderMap[GetAllAlarmsNextRequestType] = GetAllAlarmsNextRequestFrame
	encoderMap[MibUploadRequestType] = MibUploadRequestFrame
	encoderMap[MibUploadNextRequestType] = MibUploadNextRequestFrame
	encoderMap[MibResetRequestType] = MibResetRequestFrame
	encoderMap[TestRequestType] = TestRequestFrame
	encoderMap[StartSoftwareDownloadRequestType] = StartSoftwareDownloadRequestFrame
	encoderMap[DownloadSectionRequestType] = DownloadSectionRequestFrame
	encoderMap[EndSoftwareDownloadRequestType] = EndSoftwareDownloadRequestFrame
	encoderMap[ActivateSoftwareRequestType] = ActivateSoftwareRequestFrame
	encoderMap[CommitSoftwareRequestType] = CommitSoftwareRequestFrame
	encoderMap[SynchronizeTimeRequestType] = SynchronizeTimeRequestFrame
	encoderMap[RebootRequestType] = RebootRequestFrame
	encoderMap[GetNextRequestType] = GetNextRequestFrame
	encoderMap[GetCurrentDataRequestType] = GetCurrentDataRequestFrame
	encoderMap[SetTableRequestType] = SetTableRequestFrame
	encoderMap[CreateResponseType] = CreateResponseFrame
	encoderMap[DeleteResponseType] = DeleteResponseFrame
	encoderMap[SetResponseType] = SetResponseFrame
	encoderMap[GetResponseType] = GetResponseFrame
	encoderMap[GetAllAlarmsResponseType] = GetAllAlarmsResponseFrame
	encoderMap[GetAllAlarmsNextResponseType] = GetAllAlarmsNextResponseFrame
	encoderMap[MibUploadResponseType] = MibUploadResponseFrame
	encoderMap[MibUploadNextResponseType] = MibUploadNextResponseFrame
	encoderMap[MibResetResponseType] = MibResetResponseFrame
	encoderMap[TestResponseType] = TestResponseFrame
	encoderMap[StartSoftwareDownloadResponseType] = StartSoftwareDownloadResponseFrame
	encoderMap[DownloadSectionResponseType] = DownloadSectionResponseFrame
	encoderMap[EndSoftwareDownloadResponseType] = EndSoftwareDownloadResponseFrame
	encoderMap[ActivateSoftwareResponseType] = ActivateSoftwareResponseFrame
	encoderMap[CommitSoftwareResponseType] = CommitSoftwareResponseFrame
	encoderMap[SynchronizeTimeResponseType] = SynchronizeTimeResponseFrame
	encoderMap[RebootResponseType] = RebootResponseFrame
	encoderMap[GetNextResponseType] = GetNextResponseFrame
	encoderMap[GetCurrentDataResponseType] = GetCurrentDataResponseFrame
	encoderMap[SetTableResponseType] = SetTableResponseFrame
	encoderMap[AlarmNotificationType] = AlarmNotificationFrame
	encoderMap[AttributeValueChangeType] = AttributeValueChangeFrame
	encoderMap[TestResultType] = TestResultFrame
}

type options struct {
	frameFormat       DeviceIdent
	failIfTruncated   bool
	attributeMask     uint16
	result            me.Results // Common for many responses
	attrExecutionMask uint16     // Create Response Only if results == 3 or
	// Set Response only if results == 0
	unsupportedMask uint16 // Set Response only if results == 9
	sequenceNumber  uint16 // For get-next request frames and for
	// frames that return number of commands or
	// length
	transactionID uint16 // OMCI TID
}

var defaultFrameOptions = options{
	frameFormat:       BaselineIdent,
	failIfTruncated:   false,
	attributeMask:     0xFFFF,
	result:            me.Success,
	attrExecutionMask: 0,
	unsupportedMask:   0,
	sequenceNumber:    0,
	transactionID:     0,
}

// A FrameOption sets options such as frame format, etc.
type FrameOption func(*options)

// FrameFormat determines determines the OMCI message format used on the fiber.
// The default value is BaselineIdent
func FrameFormat(ff DeviceIdent) FrameOption {
	return func(o *options) {
		o.frameFormat = ff
	}
}

// FailIfTruncated determines whether a request to encode a frame that does
// not have enough room for all requested options should fail and return an
// error.
//
// If set to 'false', the behaviour depends on the message type/operation
// requested. The table below provides more information:
//
//   Request Type	Behavour
//	 ------------------------------------------------------------------------
//	 CreateRequest  A single CreateRequest struct is always returned as the
//                  CreateRequest message does not have an attributes Mask
//                  field and a Baseline OMCI message is large enough to
//                  support all Set-By-Create attributes.
//
//   GetResponse	If multiple OMCI response frames are needed to return
//					all requested attributes, only the attributes that can
//					fit will be returned and the FailedAttributeMask field
//					set to the attributes that could not be returned
//
//					If this is an ME with an attribute that is a table, the
//					first GetResponse struct will return the size of the
//					attribute and the following GetNextResponse structs will
//					contain the attribute data. The ONU application is
//					responsible for stashing these extra struct(s) away in
//					anticipation of possible GetNext Requests occurring for
//					the attribute.  See the discussion on Table attributes
//					in the GetResponse section of ITU G.988 for more
//					information.
//
// If set to 'true', no struct(s) are returned and an error is provided.
//
// The default value is 'false'
func FailIfTruncated(f bool) FrameOption {
	return func(o *options) {
		o.failIfTruncated = f
	}
}

// attributeMask determines the attributes to encode into the frame.
// The default value is 0xFFFF which specifies all available attributes
// in the frame
func AttributeMask(m uint16) FrameOption {
	return func(o *options) {
		o.attributeMask = m
	}
}

// AttributeExecutionMask is used by the Create and Set Response frames to indicate
// attributes that failed to be created/set.
func AttributeExecutionMask(m uint16) FrameOption {
	return func(o *options) {
		o.attrExecutionMask = m
	}
}

// Result is used to set returned results in responses
// that have that field
func Result(r me.Results) FrameOption {
	return func(o *options) {
		o.result = r
	}
}

// SequenceNumber is used by the GetNext and MibUploadGetNext request frames and for
// frames that return number of commands or length such as Get (table attribute) or
// MibUpload/GetAllAlarms/...
func SequenceNumber(m uint16) FrameOption {
	return func(o *options) {
		o.sequenceNumber = m
	}
}

// TransactionID is to specify the TID in the OMCI header. The default is
// zero which requires the caller to set it to the appropriate value if this
// is not an autonomous ONU notification frame
func TransactionID(tid uint16) FrameOption {
	return func(o *options) {
		o.transactionID = tid
	}
}

// EncodeFrame will encode the Managed Entity specific protocol struct and an
// OMCILayer struct. This struct can be provided to the gopacket.SerializeLayers()
// function to be serialized into a buffer for transmission.
func EncodeFrame(m *me.ManagedEntity, messageType MessageType, opt ...FrameOption) (*OMCI, gopacket.SerializableLayer, error) {
	// Check for message type support
	msgType := me.MsgType(messageType & me.MsgTypeMask)
	meDefinition := m.GetManagedEntityDefinition()

	if !me.SupportsMsgType(meDefinition, msgType) {
		msg := fmt.Sprintf("managed entity %v does not support %v Message-Type",
			meDefinition.GetName(), msgType)
		return nil, nil, errors.New(msg)
	}
	// Decode options
	opts := defaultFrameOptions
	for _, o := range opt {
		o(&opts)
	}
	// Note: Transaction ID should be set before frame serialization
	omci := &OMCI{
		TransactionID:    opts.transactionID,
		MessageType:      messageType,
		DeviceIdentifier: opts.frameFormat,
	}
	var meInfo gopacket.SerializableLayer
	var err error

	if encoder, ok := encoderMap[messageType]; ok {
		meInfo, err = encoder(m, opts)
	} else {
		err = errors.New(fmt.Sprintf("message-type: %v/%#x is not supported", messageType, messageType))
	}
	if err != nil {
		return nil, nil, err
	}
	return omci, meInfo, err
}

// For most all create methods below, error checking for valid masks, attribute
// values, and other fields is left to when the frame is actually serialized.

func checkAttributeMask(m *me.ManagedEntity, mask uint16) (uint16, error) {
	if mask == defaultFrameOptions.attributeMask {
		// Scale back to just what is allowed
		return m.GetAllowedAttributeMask(), nil
	}
	if mask&m.GetManagedEntityDefinition().GetAllowedAttributeMask() != mask {
		return 0, errors.New("invalid attribute mask")
	}
	return mask & m.GetManagedEntityDefinition().GetAllowedAttributeMask(), nil
}

// return the maximum space that can be used by attributes
func maxPacketAvailable(m *me.ManagedEntity, opt options) uint {
	if opt.frameFormat == BaselineIdent {
		// OMCI Header          - 4 octets
		// Class ID/Instance ID - 4 octets
		// Length field			- 4 octets
		// MIC                  - 4 octets
		return MaxBaselineLength - 16
	}
	// OMCI Header          - 4 octets
	// Class ID/Instance ID - 4 octets
	// Length field			- 4 octets
	// MIC                  - 4 octets
	return MaxExtendedLength - 16
}

func calculateAttributeMask(m *me.ManagedEntity, requestedMask uint16) (uint16, error) {
	attrDefs := m.GetAttributeDefinitions()
	var entityIdName string
	if entry, ok := (*attrDefs)[0]; ok {
		entityIdName = entry.GetName()
	} else {
		panic("unexpected error") // All attribute definition maps have an entity ID
	}
	attributeNames := make([]interface{}, 0)
	for attrName := range *m.GetAttributeValueMap() {
		if attrName == entityIdName {
			continue // No mask for EntityID
		}
		attributeNames = append(attributeNames, attrName)
	}
	calculatedMask, err := me.GetAttributeBitmap(*attrDefs, mapset.NewSetWith(attributeNames...))

	if err != nil {
		return 0, err
	}
	return calculatedMask & requestedMask, nil
}

func CreateRequestFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	// NOTE: The OMCI parser does not extract the default values of set-by-create attributes
	//       and are the zero 'default' (or nil) at this time.  For this reason, make sure
	//       you specify all non-zero default values and pass them in appropriate
	meLayer := &CreateRequest{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
		Attributes: *m.GetAttributeValueMap(),
	}
	return meLayer, nil
}

func CreateResponseFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	meLayer := &CreateResponse{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
		Result:                 opt.result,
		AttributeExecutionMask: opt.attributeMask,
	}
	if meLayer.Result == me.ParameterError {
		meLayer.AttributeExecutionMask = opt.attrExecutionMask
	}
	return meLayer, nil
}

func DeleteRequestFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	meLayer := &DeleteRequest{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	return meLayer, nil
}

func DeleteResponseFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	meLayer := &DeleteResponse{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
		Result: opt.result,
	}
	return meLayer, nil
}

func SetRequestFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	mask, err = calculateAttributeMask(m, mask)
	if err != nil {
		return nil, err
	}
	meDefinition := m.GetManagedEntityDefinition()
	attrDefs := *meDefinition.GetAttributeDefinitions()
	attrMap := *m.GetAttributeValueMap()

	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)
	payloadAvailable := int(maxPayload) - 2 // Less attribute mask

	meLayer := &SetRequest{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
		AttributeMask: 0,
		Attributes:    make(me.AttributeValueMap),
	}
	for mask != 0 {
		// Iterate down the attributes (Attribute 0 is the ManagedEntity ID)
		var attrIndex uint
		for attrIndex = 1; attrIndex <= 16; attrIndex++ {
			// Is this attribute requested
			if mask&(1<<(16-attrIndex)) != 0 {
				// Get definitions since we need the name
				attrDef, ok := attrDefs[attrIndex]
				if !ok {
					msg := fmt.Sprintf("Unexpected error, index %v not valued for ME %v",
						attrIndex, meDefinition.GetName())
					return nil, errors.New(msg)
				}
				var attrValue interface{}
				attrValue, ok = attrMap[attrDef.Name]
				if !ok {
					msg := fmt.Sprintf("Unexpected error, attribute %v not provided in ME %v: %v",
						attrDef.GetName(), meDefinition.GetName(), m)
					return nil, errors.New(msg)
				}
				// Is space available?
				if attrDef.Size <= payloadAvailable {
					// Mark bit handled
					mask &= ^(1 << (16 - attrIndex))
					meLayer.AttributeMask |= 1 << (16 - attrIndex)
					meLayer.Attributes[attrDef.Name] = attrValue
					payloadAvailable -= attrDef.Size
				} else {
					// TODO: Should we set truncate?
					msg := fmt.Sprintf("out-of-space. Cannot fit attribute %v into SetRequest message",
						attrDef.GetName())
					return nil, me.NewMessageTruncatedError(msg)
				}
			}
		}
	}
	if err == nil && meLayer.AttributeMask == 0 {
		// TODO: Is a set request with no attributes valid?
		return nil, errors.New("no attributes encoded for SetRequest")
	}
	return meLayer, nil
}

func SetResponseFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	meLayer := &SetResponse{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
		Result: opt.result,
	}
	if meLayer.Result == me.AttributeFailure {
		meLayer.UnsupportedAttributeMask = opt.unsupportedMask
		meLayer.FailedAttributeMask = opt.attrExecutionMask
	}
	return meLayer, nil
}

func GetRequestFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	// Given mask sent in (could be default of 0xFFFF) get what is allowable.
	// This will be all allowed if 0xFFFF is passed in, or a subset if a fixed
	// number of items.
	maxMask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Now scan attributes and reduce mask to only those
	var mask uint16
	mask, err = calculateAttributeMask(m, maxMask)
	if err != nil {
		return nil, err
	}
	if mask == 0 {
		// TODO: Is a Get request with no attributes valid?
		return nil, errors.New("no attributes requested for GetRequest")
	}
	meLayer := &GetRequest{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
		AttributeMask: mask,
	}
	return meLayer, nil
}

func GetResponseFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	if mask == 0 {
		// TODO: Is a Get request with no attributes valid?
		return nil, errors.New("no attributes encoded for Get Response")
	}
	meLayer := &GetResponse{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
		Result:        opt.result,
		AttributeMask: 0,
		Attributes:    make(me.AttributeValueMap),
	}
	if meLayer.Result == me.AttributeFailure {
		meLayer.UnsupportedAttributeMask = opt.unsupportedMask
		meLayer.FailedAttributeMask = opt.attrExecutionMask
	}
	// Encode whatever we can
	if meLayer.Result == me.Success || meLayer.Result == me.AttributeFailure {
		// Encode results
		// Get payload space available
		maxPayload := maxPacketAvailable(m, opt)
		payloadAvailable := int(maxPayload)
		meDefinition := m.GetManagedEntityDefinition()
		attrDefs := *meDefinition.GetAttributeDefinitions()
		attrMap := *m.GetAttributeValueMap()

		for mask != 0 {
			// Iterate down the attributes (Attribute 0 is the ManagedEntity ID)
			var attrIndex uint
			for attrIndex = 1; attrIndex <= 16; attrIndex++ {
				// Is this attribute requested
				if mask&(1<<(16-attrIndex)) != 0 {
					// Get definitions since we need the name
					attrDef, ok := attrDefs[attrIndex]
					if !ok {
						msg := fmt.Sprintf("Unexpected error, index %v not valued for ME %v",
							attrIndex, meDefinition.GetName())
						return nil, errors.New(msg)
					}
					var attrValue interface{}
					attrValue, ok = attrMap[attrDef.Name]
					if !ok {
						msg := fmt.Sprintf("Unexpected error, attribute %v not provided in ME %v: %v",
							attrDef.GetName(), meDefinition.GetName(), m)
						return nil, errors.New(msg)

					}
					// Is space available?
					if attrDef.Size <= payloadAvailable {
						// Mark bit handled
						mask &= ^(1 << (16 - attrIndex))
						meLayer.AttributeMask |= 1 << (16 - attrIndex)
						meLayer.Attributes[attrDef.Name] = attrValue
						payloadAvailable -= attrDef.Size

						// If it is a table, set up our getNextResponses now
						if attrDef.IsTableAttribute() {
						}
					} else if opt.failIfTruncated {
						// TODO: Should we set truncate?
						msg := fmt.Sprintf("out-of-space. Cannot fit attribute %v into GetResponse message",
							attrDef.GetName())
						return nil, me.NewMessageTruncatedError(msg)
					} else {
						// Add to existing 'failed' mask and update result
						meLayer.FailedAttributeMask |= 1 << (16 - attrIndex)
						meLayer.Result = me.AttributeFailure
					}
				}
			}
		}
	}
	return meLayer, nil
}

func GetAllAlarmsRequestFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &GetAllAlarmsRequest{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}

func GetAllAlarmsResponseFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &GetAllAlarmsResponse{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}

func GetAllAlarmsNextRequestFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &GetAllAlarmsNextRequest{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}

func GetAllAlarmsNextResponseFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &GetAllAlarmsNextResponse{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}

func MibUploadRequestFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	// Common for all MEs
	meLayer := &MibUploadRequest{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: 0,
		},
	}
	return meLayer, nil
}

func MibUploadResponseFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	// Common for all MEs
	meLayer := &MibUploadResponse{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: 0,
		},
		NumberOfCommands: opt.sequenceNumber,
	}
	return meLayer, nil
}

func MibUploadNextRequestFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	// Common for all MEs
	meLayer := &MibUploadNextRequest{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: 0,
		},
		CommandSequenceNumber: opt.sequenceNumber,
	}
	return meLayer, nil
}

func MibUploadNextResponseFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &MibUploadNextResponse{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}

func MibResetRequestFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	// Common for all MEs
	meLayer := &MibResetRequest{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	return meLayer, nil
}

func MibResetResponseFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &MibResetResponse{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}

func AlarmNotificationFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &AlarmNotificationMsg{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}

func AttributeValueChangeFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &AttributeValueChangeMsg{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}

func TestRequestFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &TestRequest{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}

func TestResponseFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &TestResponse{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}

func StartSoftwareDownloadRequestFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &StartSoftwareDownloadRequest{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}

func StartSoftwareDownloadResponseFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &StartSoftwareDownloadResponse{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}

func DownloadSectionRequestFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &DownloadSectionRequest{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}

func DownloadSectionResponseFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &DownloadSectionResponse{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}

func EndSoftwareDownloadRequestFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &EndSoftwareDownloadRequest{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}

func EndSoftwareDownloadResponseFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &EndSoftwareDownloadResponse{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}

func ActivateSoftwareRequestFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &ActivateSoftwareRequest{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}

func ActivateSoftwareResponseFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &ActivateSoftwareResponse{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}

func CommitSoftwareRequestFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &CommitSoftwareRequest{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}

func CommitSoftwareResponseFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &CommitSoftwareResponse{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}

func SynchronizeTimeRequestFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &SynchronizeTimeRequest{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}

func SynchronizeTimeResponseFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &SynchronizeTimeResponse{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}

func RebootRequestFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &RebootRequest{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}

func RebootResponseFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &RebootResponse{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}

func GetNextRequestFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// TODO: For GetNext, we may want to make sure that only 1 attribute is being requested
	// Common for all MEs
	meLayer := &GetNextRequest{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
		AttributeMask:  mask,
		SequenceNumber: opt.sequenceNumber,
	}
	return meLayer, nil
}

func GetNextResponseFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &GetNextResponse{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}

func TestResultFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &TestResultMsg{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}

func GetCurrentDataRequestFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &GetCurrentDataRequest{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}

func GetCurrentDataResponseFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &GetCurrentDataResponse{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}

func SetTableRequestFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	if opt.frameFormat != ExtendedIdent {
		return nil, errors.New("SetTable message type only supported with Extended OMCI Messaging")
	}
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &SetTableRequest{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}

func SetTableResponseFrame(m *me.ManagedEntity, opt options) (gopacket.SerializableLayer, error) {
	if opt.frameFormat != ExtendedIdent {
		return nil, errors.New("SetTable message type only supported with Extended OMCI Messaging")
	}
	mask, err := checkAttributeMask(m, opt.attributeMask)
	if err != nil {
		return nil, err
	}
	// Common for all MEs
	meLayer := &SetTableResponse{
		MeBasePacket: MeBasePacket{
			EntityClass:    m.GetClassID(),
			EntityInstance: m.GetEntityID(),
		},
	}
	// Get payload space available
	maxPayload := maxPacketAvailable(m, opt)

	// TODO: Lots of work to do

	fmt.Println(mask, maxPayload)
	return meLayer, errors.New("todo: Not implemented")
}
