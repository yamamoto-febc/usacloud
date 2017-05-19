package command

import (
	"fmt"
)

func VPCRouterDhcpServerUpdate(ctx Context, params *DhcpServerUpdateVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterDhcpServerUpdate is failed: %s", e)
	}

	if !p.HasDHCPServer() {
		return fmt.Errorf("VPCRouter[%d] don't have any DHCP servers", params.Id)
	}

	cnf := p.Settings.Router.FindDHCPServerAt(params.Index)
	if cnf == nil {
		return fmt.Errorf("DHCP server is not found on eth%d", params.Index)
	}

	if ctx.IsSet("range-start") {
		cnf.RangeStart = params.RangeStart
	}
	if ctx.IsSet("range-stop") {
		cnf.RangeStop = params.RangeStop
	}

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterDhcpServerUpdate is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterDhcpServerUpdate is failed: %s", err)
	}

	return nil

}
