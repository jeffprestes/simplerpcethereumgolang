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
 * @file transaction.go
 * @authors:
 *   Reginaldo Costa <regcostajr@gmail.com>
 * @date 2017
 */

package dto

import (
	"github.com/regcostajr/go-web3/complex/types"
)

// TransactionParameters GO transaction to make more easy controll the parameters
type TransactionParameters struct {
	From     string
	To       string
	Gas      types.ComplexIntParameter
	GasPrice types.ComplexIntParameter
	Value    types.ComplexIntParameter
	Data     types.ComplexString
}

// RequestTransactionParameters JSON
type RequestTransactionParameters struct {
	From     string `json:"from"`
	To       string `json:"to,omitempty"`
	Gas      string `json:"gas,omitempty"`
	GasPrice string `json:"gasPrice,omitempty"`
	Value    string `json:"value,omitempty"`
	Data     string `json:"data,omitempty"`
}

// Transform the GO transactions parameters to json style
func (params *TransactionParameters) Transform() *RequestTransactionParameters {
	request := new(RequestTransactionParameters)
	request.From = params.From
	if params.To != "" {
		request.To = params.To
	}
	if params.Gas != 0 {
		request.Gas = params.Gas.ToHex()
	}
	if params.GasPrice != 0 {
		request.GasPrice = params.GasPrice.ToHex()
	}
	if params.Value != 0 {
		request.Value = params.Value.ToHex()
	}
	if params.Data != "" {
		request.Data = params.Data.ToHex()
	}
	return request
}

type TransactionResponse struct {
	Hash             string                   `json:"hash"`
	Nonce            int                      `json:"nonce"`
	BlockHash        string                   `json:"blockHash"`
	BlockNumber      int64                    `json:"blockNumber"`
	TransactionIndex int64                    `json:"transactionIndex"`
	From             string                   `json:"from"`
	To               string                   `json:"to"`
	Value            types.ComplexIntResponse `json:"value"`
	GasPrice         types.ComplexIntResponse `json:"gasPrice,omitempty"`
	Gas              types.ComplexIntResponse `json:"gas,omitempty"`
	Data             types.ComplexString      `json:"data,omitempty"`
}

type TransactionReceipt struct {
	TransactionHash   string   `json:"transactionHash"`
	TransactionIndex  int64    `json:"transactionIndex"`
	BlockHash         string   `json:"blockHash"`
	BlockNumber       int64    `json:"blockNumber"`
	CumulativeGasUsed int64    `json:"cumulativeGasUsed"`
	GasUsed           int64    `json:"gasUsed"`
	ContractAddress   string   `json:"contractAddress"`
	Logs              []string `json:"logs"`
}
