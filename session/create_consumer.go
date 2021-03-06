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
	"github.com/mysteriumnetwork/node/identity"
)

// Manager defines methods for session management
type Manager interface {
	Create(consumerID identity.Identity, proposalID int) (Session, error)
}

// createConsumer processes session create requests from communication channel.
type createConsumer struct {
	SessionManager Manager
	PeerID         identity.Identity
}

// GetMessageEndpoint returns endpoint there to receive messages
func (consumer *createConsumer) GetRequestEndpoint() communication.RequestEndpoint {
	return endpointSessionCreate
}

// NewRequest creates struct where request from endpoint will be serialized
func (consumer *createConsumer) NewRequest() (requestPtr interface{}) {
	var request CreateRequest
	return &request
}

// Consume handles requests from endpoint and replies with response
func (consumer *createConsumer) Consume(requestPtr interface{}) (response interface{}, err error) {
	request := requestPtr.(*CreateRequest)

	sessionInstance, err := consumer.SessionManager.Create(consumer.PeerID, request.ProposalId)
	switch err {
	case nil:
		return responseWithSession(sessionInstance), nil
	case ErrorInvalidProposal:
		return responseInvalidProposal, nil
	default:
		return responseInternalError, nil
	}
}

func responseWithSession(sessionInstance Session) CreateResponse {
	serializedConfig, err := json.Marshal(sessionInstance.Config)
	if err != nil {
		// Failed to serialize session
		// TODO Cant expose error to response, some logging should be here
		return responseInternalError
	}

	return CreateResponse{
		Success: true,
		Session: SessionDto{
			ID:     sessionInstance.ID,
			Config: serializedConfig,
		},
	}
}
