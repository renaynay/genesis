/*
	Copyright 2019 whiteblock Inc.
	This file is a part of the genesis.

	Genesis is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	Genesis is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package sysethereum

import (
	"github.com/whiteblock/genesis/protocols/helpers"
	"github.com/whiteblock/genesis/db"
	"github.com/whiteblock/genesis/util"
	"github.com/whiteblock/mustache"
)

type sysethereumConf map[string]interface{}

type sysConf struct {
	Options []string `json:"options"`
	Extras  []string `json:"extras"`

	SenderOptions   []string `json:"senderOptions"`
	ReceiverOptions []string `json:"receiverOptions"`
	MNOptions       []string `json:"mnOptions"`

	SenderExtras   []string `json:"senderExtras"`
	ReceiverExtras []string `json:"receiverExtras"`
	MNExtras       []string `json:"mnExtras"`

	MasterNodeConns int64 `json:"masterNodeConns"`
	NodeConns       int64 `json:"nodeConns"`
	PercOfMNodes    int64 `json:"percentMasternodes"`
	Validators      int64 `json:"validators"`
}

func newConf(data map[string]interface{}) (*sysConf, error) {
	out := new(sysConf)
	return out, helpers.HandleBlockchainConfig(blockchain, data, out)
}

func makeConfig(aconf sysethereumConf, details *db.DeploymentDetails) (string, error) { // TODO where does this get called?
	sysEthConf, err := util.CopyMap(aconf)

	filler := util.ConvertToStringMap(sysEthConf)
	filler["contractsDirectory"] = "/contracts"
	filler["dataDirectory"] = "/data"
	if err != nil {
		return "", util.LogError(err)
	}

	dat, err := helpers.GetBlockchainConfig("sysethereum", 0, "sysethereum.conf.mustache", details)
	if err != nil {
		return "", util.LogError(err)
	}
	return mustache.Render(string(dat), filler)
}


