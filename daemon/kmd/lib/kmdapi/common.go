// Copyright (C) 2019-2021 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package kmdapi

import (
	"github.com/algorand/go-algorand/crypto"
	"github.com/algorand/go-algorand/protocol"
)

// SwaggerSpecJSON is autogenerated from swagger.json, and bundled in
// with a script on build through a package init() function.
var SwaggerSpecJSON string

// APIV1MasterDerivationKey is a swagger annotated alias of crypto.MasterDerivationKey
// swagger:strfmt byte
type APIV1MasterDerivationKey = crypto.MasterDerivationKey

// APIV1PrivateKey is a swagger annotated alias of crypto.privateKey
// swagger:strfmt byte
type APIV1PrivateKey = crypto.PrivateKey

// APIV1PublicKey is a swagger annotated alias of crypto.PublicKey
// swagger:strfmt byte
type APIV1PublicKey = crypto.PublicKey

// APIV1Wallet is the API's representation of a wallet
type APIV1Wallet struct {
	ID                    string            `json:"id"`
	Name                  string            `json:"name"`
	DriverName            string            `json:"driver_name"`
	DriverVersion         uint32            `json:"driver_version"`
	SupportsMnemonicUX    bool              `json:"mnemonic_ux"`
	SupportedTransactions []protocol.TxType `json:"supported_txs"`
}

// APIV1WalletHandle includes the wallet the handle corresponds to
// and the number of number of seconds to expiration
type APIV1WalletHandle struct {
	Wallet         APIV1Wallet `json:"wallet"`
	ExpiresSeconds int64       `json:"expires_seconds"`
}