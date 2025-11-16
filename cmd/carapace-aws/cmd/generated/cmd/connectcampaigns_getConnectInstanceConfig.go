package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaigns_getConnectInstanceConfigCmd = &cobra.Command{
	Use:   "get-connect-instance-config",
	Short: "Get the specific Connect instance config.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaigns_getConnectInstanceConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcampaigns_getConnectInstanceConfigCmd).Standalone()

		connectcampaigns_getConnectInstanceConfigCmd.Flags().String("connect-instance-id", "", "")
		connectcampaigns_getConnectInstanceConfigCmd.MarkFlagRequired("connect-instance-id")
	})
	connectcampaignsCmd.AddCommand(connectcampaigns_getConnectInstanceConfigCmd)
}
