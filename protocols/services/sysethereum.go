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

package services

import (
	"encoding/json"
	"github.com/whiteblock/genesis/db"
	"github.com/whiteblock/genesis/protocols/helpers"
	"github.com/whiteblock/genesis/ssh"
	"github.com/whiteblock/genesis/testnet"
	"github.com/whiteblock/genesis/util"
	"github.com/whiteblock/mustache"
)

// SysethereumService represents the SysethereumService service
type SysethereumService struct {
	SimpleService
}

type sysethereumConf map[string]interface{}

func newConf(data map[string]interface{}) (sysethereumConf, error) {
	rawDefaults := helpers.DefaultGetDefaultsFn("sysethereum")()
	defaults := map[string]interface{}{}

	err := json.Unmarshal([]byte(rawDefaults), &defaults)
	if err != nil {
		return nil, util.LogError(err)
	}
	finalData := util.MergeStringMaps(defaults, data)
	out := new(sysethereumConf)
	*out = sysethereumConf(finalData)

	return *out, nil
}

// Prepare prepares the sysethereum service
func (p SysethereumService) Prepare(client ssh.Client, tn *testnet.TestNet) error {
	aconf, err := newConf(tn.LDD.Params)
	if err != nil {
		return util.LogError(err)
	}

	err = helpers.CreateConfigs(tn, "/sysethereum.conf", func(node ssh.Node) ([]byte, error) {
		defer tn.BuildState.IncrementBuildProgress()
		conf, err := makeConfig(aconf, &tn.CombinedDetails)
		return []byte(conf), err
	})
	if err != nil {
		return util.LogError(err)
	}
	return nil
}

func makeConfig(aconf sysethereumConf, details *db.DeploymentDetails) (string, error) {

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

//GetCommand gets the command flags
func (p SysethereumService) GetCommand() string {
	return "-Dsysethereum.agents.conf.file=/sysethereum.conf"
}

// RegisterSysethereum exposes a Sysethereum service on the testnet.
func RegisterSysethereum() Service {
	return SysethereumService{
		SimpleService{
			Name:    "ganache",
			Image:   "gcr.io/whiteblock/sysethereum-agents",
			Env:     map[string]string{},
			Ports:   []string{},
			Volumes: []string{},
		},
	}
}
