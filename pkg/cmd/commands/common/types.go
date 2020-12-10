// Copyright 2017-2020 The Usacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

import (
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/util"
)

type EditRequest struct {
	HostName string `cli:",category=diskedit,order=10"`
	Password string `cli:",category=diskedit,order=20"`

	IPAddress      string `cli:",category=diskedit,order=30"`
	NetworkMaskLen int    `cli:",category=diskedit,order=31"`
	DefaultRoute   string `cli:",category=diskedit,order=32"`

	DisablePWAuth       bool `cli:",category=diskedit,order=40"`
	EnableDHCP          bool `cli:",category=diskedit,order=50"`
	ChangePartitionUUID bool `cli:",category=diskedit,order=60"`

	SSHKeys            []string   `cli:"ssh-keys,category=diskedit,order=70"`
	SSHKeyIDs          []types.ID `cli:"ssh-key-ids,category=diskedit,order=71"`
	IsSSHKeysEphemeral bool       `cli:"make-ssh-keys-ephemeral,category=diskedit,order=72"`

	NoteIDs          []types.ID              `cli:"note-ids,category=diskedit,order=80" mapconv:"-"`
	NotesData        string                  `cli:"notes,category=diskedit,order=81" mapconv:"-"`
	IsNotesEphemeral bool                    `cli:"make-notes-ephemeral,category=diskedit,order=82"`
	Notes            []*sacloud.DiskEditNote `cli:"-"` // --parametersでファイルからパラメータ指定する場合向け
}

// Customize パラメータ変換処理
func (p *EditRequest) Customize(_ cli.Context) error {
	var notes []*sacloud.DiskEditNote
	if p.NotesData != "" {
		if err := util.MarshalJSONFromPathOrContent(p.NotesData, &notes); err != nil {
			return err
		}
	}

	for _, id := range p.NoteIDs {
		notes = append(notes, &sacloud.DiskEditNote{ID: id})
	}

	p.Notes = append(p.Notes, notes...)
	return nil
}
