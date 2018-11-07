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
 *              https://github.com/cboling/OMCI-parser
 */
package generated

import (
	"../../omci"
)

type ReAniG struct {
	omci.BaseManagedEntity
}

func NewReAniG(params ...ParamData) (omci.IManagedEntity, error) {
	eid := decodeEntityID(params...)
	entity := omci.BaseManagedEntity{
		Name:     "ReAniG",
		ClassID:  313,
		EntityID: eid,
		MessageTypes: []omci.MsgType{
			omci.Set,
			omci.Get,
		},
		AttributeMask: 0,
		Attributes: []omci.IAttribute{
			omci.NewUint16Field("ManagedEntityId", 0, omci.Read),
			omci.NewByteField("AdministrativeState", 0, omci.Read|omci.Write),
			omci.NewByteField("OperationalState", 0, omci.Read),
			omci.NewByteField("Arc", 0, omci.Read|omci.Write),
			omci.NewByteField("ArcInterval", 0, omci.Read|omci.Write),
			omci.NewUint16Field("OpticalSignalLevel", 0, omci.Read),
			omci.NewByteField("LowerOpticalThreshold", 0, omci.Read|omci.Write),
			omci.NewByteField("UpperOpticalThreshold", 0, omci.Read|omci.Write),
			omci.NewUint16Field("TransmitOpticalLevel", 0, omci.Read),
			omci.NewByteField("LowerTransmitPowerThreshold", 0, omci.Read|omci.Write),
			omci.NewByteField("UpperTransmitPowerThreshold", 0, omci.Read|omci.Write),
			omci.NewByteField("UsageMode", 0, omci.Read|omci.Write),
			omci.NewUint32Field("TargetUpstreamFrequency", 0, omci.Read|omci.Write),
			omci.NewUint32Field("TargetDownstreamFrequency", 0, omci.Read|omci.Write),
			omci.NewByteField("UpstreamSignalTransmissionMode", 0, omci.Read|omci.Write),
		},
	}
	entity.ComputeAttributeMask()
	return &ReAniG{entity}, nil
}
