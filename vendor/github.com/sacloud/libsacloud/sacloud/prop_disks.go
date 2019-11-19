// Copyright 2016-2019 The Libsacloud Authors
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

package sacloud

// propDisks ディスク配列内包型
type propDisks struct {
	Disks []Disk `json:",omitempty"` // ディスク
}

// GetDisks ディスク配列 取得
func (p *propDisks) GetDisks() []Disk {
	return p.Disks
}

// GetDiskIDs ディスクID配列を返す
func (p *propDisks) GetDiskIDs() []int64 {

	ids := []int64{}
	for _, disk := range p.Disks {
		ids = append(ids, disk.ID)
	}
	return ids

}
