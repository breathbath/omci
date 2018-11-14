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

import (
	"errors"
	"fmt"
	"github.com/google/gopacket"
	"math/bits"
)

// MsgType represents a OMCI message-type
type MsgType byte

// MsgType represents the status field in a OMCI Response frame
type Results byte

// AttributeAccess represents the access allowed to an Attribute.  Some MEs
// are instantiated by the ONU autonomously. Others are instantiated on
// explicit request of the OLT via a create command, and a few ME types may
// be instantiated in either way, depending on the ONU architecture or
// circumstances.
//
//Attributes of an ME that is auto-instantiated by the ONU can be read (R),
// write (W), or read, write (R, W). On the other hand, attributes of a ME
// that is instantiated by the OLT can be either (R), (W), (R, W),
// (R, set by create) or (R, W, set by create).
type AttributeAccess byte

const (
	// AK (Bit 6), indicates whether this message is an AK to an action request.
	// If a message is an AK, this bit is set to 1. If the message is not a
	// response to a command, this bit is set to 0. In messages sent by the OLT,
	// this bit is always 0.
	AK byte = 0x20

	// AR (Bit 7), acknowledge request, indicates whether the message requires an
	// AK. An AK is a response to an action request, not a link layer handshake.
	// If an AK is expected, this bit is set to 1. If no AK is expected, this bit
	// is 0. In messages sent by the ONU, this bit is always 0
	AR byte = 0x40

	// MsgTypeMask provides a mask to get the base message type
	MsgTypeMask = 0x1F
)

const (
	// Message Types
	_                             = iota
	Create                MsgType = 4
	Delete                MsgType = 6
	Set                   MsgType = 8
	Get                   MsgType = 9
	GetAllAlarms          MsgType = 11
	GetAllAlarmsNext      MsgType = 12
	MibUpload             MsgType = 13
	MibUploadNext         MsgType = 14
	MibReset              MsgType = 15
	AlarmNotification     MsgType = 16
	AttributeValueChange  MsgType = 17
	Test                  MsgType = 18
	StartSoftwareDownload MsgType = 19
	DownloadSection       MsgType = 20
	EndSoftwareDownload   MsgType = 21
	ActivateSoftware      MsgType = 22
	CommitSoftware        MsgType = 23
	SynchronizeTime       MsgType = 24
	Reboot                MsgType = 25
	GetNext               MsgType = 26
	TestResult            MsgType = 27
	GetCurrentData        MsgType = 28
	SetTable              MsgType = 29 // Defined in Extended Message Set Only
)

func (mt MsgType) String() string {
	switch mt {
	default:
		return "Unknown"
	case Create:
		return "Create"
	case Delete:
		return "Delete"
	case Set:
		return "Set"
	case Get:
		return "Get"
	case GetAllAlarms:
		return "Get All Alarms"
	case GetAllAlarmsNext:
		return "Get All Alarms Next"
	case MibUpload:
		return "MIB Upload"
	case MibUploadNext:
		return "MIB Upload Next"
	case MibReset:
		return "MIB Reset"
	case AlarmNotification:
		return "Alarm Notification"
	case AttributeValueChange:
		return "Attribute Value Change"
	case Test:
		return "Test"
	case StartSoftwareDownload:
		return "Start Software Download"
	case DownloadSection:
		return "Download Section"
	case EndSoftwareDownload:
		return "EndSoftware Download"
	case ActivateSoftware:
		return "Activate Software"
	case CommitSoftware:
		return "Commit Software"
	case SynchronizeTime:
		return "Synchronize Time"
	case Reboot:
		return "Reboot"
	case GetNext:
		return "Get Next"
	case TestResult:
		return "Test Result"
	case GetCurrentData:
		return "Get Current Data"
	case SetTable:
		return "Set Table"
	}
}

var allNotificationTypes = [...]MsgType{
	AlarmNotification,
	AttributeValueChange,
	TestResult,
}

// SupportsMsgType returns true if the managed entity supports the desired
// Message Type / action
func SupportsMsgType(entity IManagedEntityDefinition, msgType MsgType) bool {
	for _, msg := range entity.GetMessageTypes() {
		if msgType == msg {
			return true
		}
	}
	return false
}

func IsAutonomousNotification(mt MsgType) bool {
	for _, m := range allNotificationTypes {
		if mt == m {
			return true
		}
	}
	return false
}

const (
	// Response status codes
	_                        = iota
	Success          Results = 0 // command processed successfully
	ProcessingError  Results = 1 // command processing error
	NotSupported     Results = 2 // command not supported
	ParameterError   Results = 3 // parameter error
	UnknownEntity    Results = 4 // unknown managed entity
	UnknownInstance  Results = 5 // unknown managed entity instance
	DeviceBusy       Results = 6 // device busy
	InstanceExists   Results = 7 // instance exists
	AttributeFailure Results = 9 // Attribute(s) failed or unknown
)

func (rc Results) String() string {
	switch rc {
	default:
		return "Unknown"
	case Success:
		return "Success"
	case ProcessingError:
		return "Processing Error"
	case NotSupported:
		return "Not Supported"
	case ParameterError:
		return "Parameter Error"
	case UnknownEntity:
		return "Unknown Entity"
	case UnknownInstance:
		return "Unknown Instance"
	case DeviceBusy:
		return "Device Busy"
	case InstanceExists:
		return "Instance Exists"
	case AttributeFailure:
		return "Attribute Failure"
	}
}

const (
	// Access allowed on a Managed Entity attribute
	Read AttributeAccess = 1 << iota
	Write
	SetByCreate
)

func (access AttributeAccess) String() string {
	switch access {
	default:
		return "Unknown"
	case Read:
		return "Read"
	case Write:
		return "Write"
	case SetByCreate:
		return "SetByCreate"
	case Read | Write:
		return "Read/Write"
	case Read | SetByCreate:
		return "Read/SetByCreate"
	case Write | SetByCreate:
		return "Write/SetByCreate"
	case Read | Write | SetByCreate:
		return "Read/Write/SetByCreate"
	}
}

// SupportsAttributeAccess returns true if the managed entity attribute
// supports the desired access
func SupportsAttributeAccess(attr *AttributeDefinition, acc AttributeAccess) bool {
	return attr.GetAccess()&acc == acc
}

type IManagedEntityDefinition interface {
	GetName() string
	GetClassID() uint16
	GetEntityID() uint16
	GetMessageTypes() []MsgType
	GetAllowedAttributeMask() uint16
	GetAttributeDefinitions() AttributeDefinitionMap

	DecodeAttributes(uint16, []byte, gopacket.PacketBuilder) (AttributeValueMap, error)
	SerializeAttributes(AttributeValueMap, uint16, gopacket.SerializeBuffer) error
}

type BaseManagedEntityDefinition struct {
	Name                 string
	ClassID              uint16
	EntityID             uint16       // TODO: Move to be inside attributes
	MessageTypes         []MsgType
	AllowedAttributeMask uint16
	AttributeDefinitions AttributeDefinitionMap
}

func (bme *BaseManagedEntityDefinition) String() string {
	return fmt.Sprintf("Definition: %v: CID: %v (%#x), EID: %v (%#x), Attributes: %v",
		bme.Name, bme.ClassID, bme.ClassID, bme.EntityID, bme.EntityID,
		bme.AttributeDefinitions)
}

func (bme *BaseManagedEntityDefinition) GetName() string {
	return bme.Name
}
func (bme *BaseManagedEntityDefinition) GetClassID() uint16 {
	return bme.ClassID
}
func (bme *BaseManagedEntityDefinition) GetEntityID() uint16 {
	return bme.EntityID
}
func (bme *BaseManagedEntityDefinition) GetMessageTypes() []MsgType {
	return bme.MessageTypes
}
func (bme *BaseManagedEntityDefinition) GetAllowedAttributeMask() uint16 {
	return bme.AllowedAttributeMask
}
func (bme *BaseManagedEntityDefinition) GetAttributeDefinitions() AttributeDefinitionMap {
	return bme.AttributeDefinitions
}

func (bme *BaseManagedEntityDefinition) computeAttributeMask() {
	for index := range bme.AttributeDefinitions {
		if index == 0 {
			continue	// Skip Entity ID
		}
		bme.AllowedAttributeMask |= 1 << (15 - uint(index - 1))
	}
}

func (bme* BaseManagedEntityDefinition) DecodeAttributes(mask uint16, data []byte, p gopacket.PacketBuilder) (AttributeValueMap, error) {
	if (mask | bme.GetAllowedAttributeMask()) != bme.GetAllowedAttributeMask() {
		// TODO: Provide custom error code so a response 'result' can properly be coded
		return nil, errors.New("unsupported attribute mask")
	}
	keyList := GetAttributeDefinitionMapKeys(bme.AttributeDefinitions)

	attrMap := make(AttributeValueMap, bits.OnesCount16(mask))
	for _, index := range keyList {
		if index == 0 {
			continue	// Skip Entity ID
		}
		attrDef := bme.AttributeDefinitions[index]

		if mask & (1 << (15 - uint(index - 1))) != 0 {
			value, err := attrDef.Decode(data, p)
			if err != nil {
				return nil, err
			}
			attrMap[index] = &AttributeValue{
				Name: attrDef.GetName(),
				Value: value,
			}
			data = data[attrDef.GetSize():]
		}
	}
	return attrMap, nil
}

func (bme* BaseManagedEntityDefinition) SerializeAttributes(attr AttributeValueMap, mask uint16, b gopacket.SerializeBuffer) error {
	if (mask | bme.GetAllowedAttributeMask()) != bme.GetAllowedAttributeMask() {
		// TODO: Provide custom error code so a response 'result' can properly be coded
		return errors.New("unsupported attribute mask")
	}
	keyList := GetAttributeDefinitionMapKeys(bme.AttributeDefinitions)

	for _, index := range keyList {
		if index == 0 {
			continue	// Skip Entity ID
		}
		attrDef := bme.AttributeDefinitions[index]

		if mask & (1 << (15 - uint(index - 1))) != 0 {
			attribute, ok := attr[index]
			if !ok {
				// TODO: Custom error, or more detail?
				return errors.New("attribute not found")
			}
			err := attrDef.SerializeTo(attribute.Value, b)
			if err != nil {
				return nil
			}
		}
	}
	return nil
}
