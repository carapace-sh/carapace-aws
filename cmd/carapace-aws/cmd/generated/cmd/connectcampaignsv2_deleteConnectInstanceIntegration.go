package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_deleteConnectInstanceIntegrationCmd = &cobra.Command{
	Use:   "delete-connect-instance-integration",
	Short: "Delete the integration for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_deleteConnectInstanceIntegrationCmd).Standalone()

	connectcampaignsv2_deleteConnectInstanceIntegrationCmd.Flags().String("connect-instance-id", "", "")
	connectcampaignsv2_deleteConnectInstanceIntegrationCmd.Flags().String("integration-identifier", "", "")
	connectcampaignsv2_deleteConnectInstanceIntegrationCmd.MarkFlagRequired("connect-instance-id")
	connectcampaignsv2_deleteConnectInstanceIntegrationCmd.MarkFlagRequired("integration-identifier")
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_deleteConnectInstanceIntegrationCmd)
}
