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

type SipAgentConfigData struct {
	omci.BaseManagedEntity
}

func NewSipAgentConfigData(params ...ParamData) (omci.IManagedEntity, error) {
	eid := decodeEntityID(params...)
	entity := omci.BaseManagedEntity{
		Name:     "SipAgentConfigData",
		ClassID:  150,
		EntityID: eid,
		MessageTypes: []omci.MsgType{
			omci.Set,
			omci.Get,
			omci.Create,
			omci.Delete,
		},
		AttributeMask: 0,
		Attributes: []omci.IAttribute{
			omci.NewUint16Field("ManagedEntityId", 0, omci.Read|omci.SetByCreate),
			omci.NewUint16Field("ProxyServerAddressPointer", 0, omci.Read|omci.Write|omci.SetByCreate),
			omci.NewUint16Field("OutboundProxyAddressPointer", 0, omci.Read|omci.Write|omci.SetByCreate),
			omci.NewUint32Field("PrimarySipDns", 0, omci.Read|omci.Write|omci.SetByCreate),
			omci.NewUint32Field("SecondarySipDns", 0, omci.Read|omci.Write|omci.SetByCreate),
			omci.NewUint16Field("TcpUdpPointer", 0, omci.Read|omci.Write),
			omci.NewUint32Field("SipRegExpTime", 0, omci.Read|omci.Write),
			omci.NewUint32Field("SipReregHeadStartTime", 0, omci.Read|omci.Write),
			omci.NewUint16Field("HostPartUri", 0, omci.Read|omci.Write|omci.SetByCreate),
			omci.NewByteField("SipStatus", 0, omci.Read),
			omci.NewUint16Field("SipRegistrar", 0, omci.Read|omci.Write|omci.SetByCreate),
			omci.NewUint32Field("Softswitch", 0, omci.Read|omci.Write|omci.SetByCreate),
			omci.NewUnknownField("SipResponseTable", 0, omci.Read|omci.Write),
			omci.NewByteField("SipOptionTransmitControl", 0, omci.Read|omci.Write|omci.SetByCreate),
			omci.NewByteField("SipUriFormat", 0, omci.Read|omci.Write|omci.SetByCreate),
			omci.NewUint16Field("RedundantSipAgentPointer", 0, omci.Read|omci.Write|omci.SetByCreate),
		},
	}
	entity.ComputeAttributeMask()
	return &SipAgentConfigData{entity}, nil
}
