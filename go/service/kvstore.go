// Copyright 2019 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

// RPC handlers for kvstore operations

package service

import (
	"github.com/keybase/client/go/chat/globals"
	"github.com/keybase/client/go/libkb"
	keybase1 "github.com/keybase/client/go/protocol/keybase1"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	"golang.org/x/net/context"
)

type KVStoreHandler struct {
	*BaseHandler
	globals.Contextified
	connID  libkb.ConnectionID
	service *Service
}

var _ keybase1.KvstoreInterface = (*KVStoreHandler)(nil)

func NewKVStoreHandler(xp rpc.Transporter, id libkb.ConnectionID, g *globals.Context, service *Service) *KVStoreHandler {
	return &KVStoreHandler{
		BaseHandler:  NewBaseHandler(g.ExternalG(), xp),
		Contextified: globals.NewContextified(g),
		connID:       id,
		service:      service,
	}
}

func (h *KVStoreHandler) assertLoggedIn(ctx context.Context) error {
	loggedIn := h.G().ExternalG().ActiveDevice.Valid()
	if !loggedIn {
		return libkb.LoginRequiredError{}
	}
	return nil
}

func (h *KVStoreHandler) GetKVEntry(ctx context.Context, arg keybase1.GetKVEntryArg) (res keybase1.KVEntry, err error) {
	ctx = libkb.WithLogTag(ctx, "KV")
	if err := h.assertLoggedIn(ctx); err != nil {
		return res, err
	}
	return res, nil
	// arg2 := keybase1.TeamCreateWithSettingsArg{
	// 	SessionID:   arg.SessionID,
	// 	Name:        arg.Name,
	// 	JoinSubteam: arg.JoinSubteam,
	// }
	// return h.TeamCreateWithSettings(ctx, arg2)
}

func (h *KVStoreHandler) PutKVEntry(ctx context.Context, arg keybase1.PutKVEntryArg) (res keybase1.Revision, err error) {
	ctx = libkb.WithLogTag(ctx, "KV")
	if err := h.assertLoggedIn(ctx); err != nil {
		return res, err
	}
	return res, nil
}
