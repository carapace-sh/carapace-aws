package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaigns_deleteConnectInstanceConfigCmd = &cobra.Command{
	Use:   "delete-connect-instance-config",
	Short: "Deletes a connect instance config from the specified AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaigns_deleteConnectInstanceConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcampaigns_deleteConnectInstanceConfigCmd).Standalone()

		connectcampaigns_deleteConnectInstanceConfigCmd.Flags().String("connect-instance-id", "", "")
		connectcampaigns_deleteConnectInstanceConfigCmd.MarkFlagRequired("connect-instance-id")
	})
	connectcampaignsCmd.AddCommand(connectcampaigns_deleteConnectInstanceConfigCmd)
}
