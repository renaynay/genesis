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
	"github.com/whiteblock/genesis/protocols/registrar"
	"github.com/whiteblock/genesis/util"
	"github.com/whiteblock/genesis/testnet"
	"github.com/whiteblock/genesis/protocols/services"
	"github.com/whiteblock/renaynay/genesis/protocols/helpers"
)

var conf = util.GetConfig()

const (
	blockchain = "syscoin"
)

func init() {
	registrar.RegisterBuild(blockchain, build)
	registrar.RegisterAddNodes(blockchain, add)
	registrar.RegisterServices(blockchain, GetServices) // TODO do we need services?
	registrar.RegisterDefaults(blockchain, helpers.DefaultGetDefaultsFn(blockchain))
	registrar.RegisterParams(blockchain, helpers.DefaultGetParamsFn(blockchain))

	registrar.RegisterBlockchainSideCars(blockchain, func(tn *testnet.TestNet) []string {
		sconf, err := newConf(tn.LDD.Extras)
	})
}

func build(tn *testnet.TestNet) error {

	return nil
}

func add(tn *testnet.TestNet) error {
	return nil
}

func GetServices() []services.Service {

	return nil // TODO
}