// Copyright 2016-2020 The Libsacloud Authors
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

package disk

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud"
)

func (s *Service) Apply(req *ApplyRequest) (*sacloud.Disk, *sacloud.SSHKeyGenerated, error) {
	return s.ApplyWithContext(context.Background(), req)
}

func (s *Service) ApplyWithContext(ctx context.Context, req *ApplyRequest) (*sacloud.Disk, *sacloud.SSHKeyGenerated, error) {
	if err := req.Validate(); err != nil {
		return nil, nil, err
	}

	builder, err := req.Builder(s.caller)
	if err != nil {
		return nil, nil, err
	}

	// create
	if req.ID.IsEmpty() {
		res, err := builder.Build(ctx, req.Zone, req.ServerID)
		if err != nil {
			return nil, nil, err
		}
		diskOp := sacloud.NewDiskOp(s.caller)
		disk, err := diskOp.Read(ctx, req.Zone, res.DiskID)
		if err != nil {
			return nil, nil, err
		}
		return disk, res.GeneratedSSHKey, nil
	}

	// update
	res, err := builder.Update(ctx, req.Zone)
	if err != nil {
		return nil, nil, err
	}
	return res.Disk, nil, nil
}