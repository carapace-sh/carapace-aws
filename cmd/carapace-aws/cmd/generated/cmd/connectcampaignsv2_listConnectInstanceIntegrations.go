package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_listConnectInstanceIntegrationsCmd = &cobra.Command{
	Use:   "list-connect-instance-integrations",
	Short: "Provides summary information about the integration under the specified Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_listConnectInstanceIntegrationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcampaignsv2_listConnectInstanceIntegrationsCmd).Standalone()

		connectcampaignsv2_listConnectInstanceIntegrationsCmd.Flags().String("connect-instance-id", "", "")
		connectcampaignsv2_listConnectInstanceIntegrationsCmd.Flags().String("max-results", "", "")
		connectcampaignsv2_listConnectInstanceIntegrationsCmd.Flags().String("next-token", "", "")
		connectcampaignsv2_listConnectInstanceIntegrationsCmd.MarkFlagRequired("connect-instance-id")
	})
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_listConnectInstanceIntegrationsCmd)
}
