package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_putConnectInstanceIntegrationCmd = &cobra.Command{
	Use:   "put-connect-instance-integration",
	Short: "Put or update the integration for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_putConnectInstanceIntegrationCmd).Standalone()

	connectcampaignsv2_putConnectInstanceIntegrationCmd.Flags().String("connect-instance-id", "", "")
	connectcampaignsv2_putConnectInstanceIntegrationCmd.Flags().String("integration-config", "", "")
	connectcampaignsv2_putConnectInstanceIntegrationCmd.MarkFlagRequired("connect-instance-id")
	connectcampaignsv2_putConnectInstanceIntegrationCmd.MarkFlagRequired("integration-config")
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_putConnectInstanceIntegrationCmd)
}
