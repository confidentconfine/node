/*
 * Copyright (C) 2017 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package session

import (
	"encoding/json"

	"github.com/mysteriumnetwork/node/communication"
)

const endpointSessionCreate = communication.RequestEndpoint("session-create")

var (
	responseInvalidProposal = CreateResponse{Success: false, Message: "Invalid Proposal"}
	responseInternalError   = CreateResponse{Success: false, Message: "Internal Error"}
)

// CreateRequest structure represents message from service consumer to initiate session for given proposal id
type CreateRequest struct {
	ProposalId int `json:"proposal_id"`
}

// CreateResponse structure represents service provider response to given session request from consumer
type CreateResponse struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Session SessionDto `json:"session"`
}

// SessionDto structure represents session information data within session creation response (session id and configuration options for underlaying service type)
type SessionDto struct {
	ID     ID              `json:"id"`
	Config json.RawMessage `json:"config"`
}
