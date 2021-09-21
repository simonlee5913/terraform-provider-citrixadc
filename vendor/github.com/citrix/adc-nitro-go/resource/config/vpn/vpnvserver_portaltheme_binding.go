/*
* Copyright (c) 2021 Citrix Systems, Inc.
*
*   Licensed under the Apache License, Version 2.0 (the "License");
*   you may not use this file except in compliance with the License.
*   You may obtain a copy of the License at
*
*       http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*   distributed under the License is distributed on an "AS IS" BASIS,
*   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*   See the License for the specific language governing permissions and
*   limitations under the License.
*/

package vpn

/**
* Binding class showing the portaltheme that can be bound to vpnvserver.
*/
type Vpnvserverportalthemebinding struct {
	/**
	* Name of the portal theme bound to VPN vserver
	*/
	Portaltheme string `json:"portaltheme,omitempty"`
	Acttype uint32 `json:"acttype,omitempty"`
	/**
	* Name of the virtual server.
	*/
	Name string `json:"name,omitempty"`


}