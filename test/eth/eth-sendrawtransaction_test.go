/********************************************************************************
   This file is part of go-web3.
   go-web3 is free software: you can redistribute it and/or modify
   it under the terms of the GNU Lesser General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   go-web3 is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Lesser General Public License for more details.
   You should have received a copy of the GNU Lesser General Public License
   along with go-web3.  If not, see <http://www.gnu.org/licenses/>.
*********************************************************************************/

/**
 * @file eth-sendrawtransaction_test.go
 * @authors:
 *   Junjie Chen <chuckjunjchen@gmail.com>
 * @date 2018
 */
package test

import (
	"testing"

	"github.com/regcostajr/go-web3"
	"github.com/regcostajr/go-web3/complex/types"
	"github.com/regcostajr/go-web3/providers"
	"github.com/regcostajr/go-web3/dto"
	"math/big"
)

func TestSendRawTransaction(t *testing.T) {

	var connection = web3.NewWeb3(providers.NewHTTPProvider("127.0.0.1:8545", 10, false))

	coinbase, err := connection.Eth.GetCoinbase()

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	transaction := new(dto.TransactionParameters)
	transaction.From = coinbase
	transaction.To = coinbase
	transaction.Value = big.NewInt(0).Mul(big.NewInt(500), big.NewInt(1E18))
	transaction.Gas = big.NewInt(40000)
	transaction.GasPrice = big.NewInt(1E9)
	transaction.Data = types.ComplexString("p2p transaction")

	signedTx, err := connection.Eth.SignTransaction(transaction)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	txHash, err := connection.Eth.SendRawTransaction(signedTx.Raw.ToHex())

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(txHash)
}
