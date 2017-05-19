package command

import (
	"fmt"
)

func VPCRouterDhcpServerAdd(ctx Context, params *DhcpServerAddVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterDhcpServerAdd is failed: %s", e)
	}

	if !p.HasSetting() {
		p.InitVPCRouterSetting()
	}

	p.Settings.Router.AddDHCPServer(
		params.Index,
		params.RangeStart,
		params.RangeStop,
	)

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterDhcpServerAdd is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterDhcpServerAdd is failed: %s", err)
	}
	return nil
}
