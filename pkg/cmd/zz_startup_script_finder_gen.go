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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-resource-finder'; DO NOT EDIT

package cmd

import (
	"fmt"
	"strings"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
	"github.com/sacloud/usacloud/pkg/util"
)

func findStartupScriptReadTargets(ctx cli.Context, param *params.ReadStartupScriptParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().Note

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.Notes {
			if util.HasTags(&v, param.Selector) {
				ids = append(ids, v.GetID())
			}
		}
		if len(ids) == 0 {
			return ids, fmt.Errorf("finding resource id is failed: not found with search param [tags=%s]", param.Selector)
		}
	} else {
		for _, arg := range args {
			for _, a := range strings.Split(arg, "\n") {
				idOrName := a
				if id := sacloud.StringID(idOrName); !id.IsEmpty() {
					ids = append(ids, id)
				} else {
					apiClient.Reset()
					apiClient.SetFilterBy("Name", idOrName)
					res, err := apiClient.Find()
					if err != nil {
						return ids, fmt.Errorf("finding resource id is failed: %s", err)
					}
					if res.Count == 0 {
						return ids, fmt.Errorf("finding resource id is failed: not found with search param [%q]", idOrName)
					}
					for _, v := range res.Notes {
						if len(param.Selector) == 0 || util.HasTags(&v, param.Selector) {
							ids = append(ids, v.GetID())
						}
					}
				}
			}

		}

	}

	ids = util.UniqIDs(ids)
	if len(ids) == 0 {
		return ids, fmt.Errorf("finding resource is is failed: not found")
	}
	if len(ids) != 1 {
		return ids, fmt.Errorf("could not run with multiple targets: %v", ids)
	}

	return ids, nil
}

func findStartupScriptUpdateTargets(ctx cli.Context, param *params.UpdateStartupScriptParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().Note

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.Notes {
			if util.HasTags(&v, param.Selector) {
				ids = append(ids, v.GetID())
			}
		}
		if len(ids) == 0 {
			return ids, fmt.Errorf("finding resource id is failed: not found with search param [tags=%s]", param.Selector)
		}
	} else {
		for _, arg := range args {
			for _, a := range strings.Split(arg, "\n") {
				idOrName := a
				if id := sacloud.StringID(idOrName); !id.IsEmpty() {
					ids = append(ids, id)
				} else {
					apiClient.Reset()
					apiClient.SetFilterBy("Name", idOrName)
					res, err := apiClient.Find()
					if err != nil {
						return ids, fmt.Errorf("finding resource id is failed: %s", err)
					}
					if res.Count == 0 {
						return ids, fmt.Errorf("finding resource id is failed: not found with search param [%q]", idOrName)
					}
					for _, v := range res.Notes {
						if len(param.Selector) == 0 || util.HasTags(&v, param.Selector) {
							ids = append(ids, v.GetID())
						}
					}
				}
			}

		}

	}

	ids = util.UniqIDs(ids)
	if len(ids) == 0 {
		return ids, fmt.Errorf("finding resource is is failed: not found")
	}

	return ids, nil
}

func findStartupScriptDeleteTargets(ctx cli.Context, param *params.DeleteStartupScriptParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().Note

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.Notes {
			if util.HasTags(&v, param.Selector) {
				ids = append(ids, v.GetID())
			}
		}
		if len(ids) == 0 {
			return ids, fmt.Errorf("finding resource id is failed: not found with search param [tags=%s]", param.Selector)
		}
	} else {
		for _, arg := range args {
			for _, a := range strings.Split(arg, "\n") {
				idOrName := a
				if id := sacloud.StringID(idOrName); !id.IsEmpty() {
					ids = append(ids, id)
				} else {
					apiClient.Reset()
					apiClient.SetFilterBy("Name", idOrName)
					res, err := apiClient.Find()
					if err != nil {
						return ids, fmt.Errorf("finding resource id is failed: %s", err)
					}
					if res.Count == 0 {
						return ids, fmt.Errorf("finding resource id is failed: not found with search param [%q]", idOrName)
					}
					for _, v := range res.Notes {
						if len(param.Selector) == 0 || util.HasTags(&v, param.Selector) {
							ids = append(ids, v.GetID())
						}
					}
				}
			}

		}

	}

	ids = util.UniqIDs(ids)
	if len(ids) == 0 {
		return ids, fmt.Errorf("finding resource is is failed: not found")
	}

	return ids, nil
}
