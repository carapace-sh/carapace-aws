package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_deleteConnectInstanceConfigCmd = &cobra.Command{
	Use:   "delete-connect-instance-config",
	Short: "Deletes a connect instance config from the specified AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_deleteConnectInstanceConfigCmd).Standalone()

	connectcampaignsv2_deleteConnectInstanceConfigCmd.Flags().String("campaign-deletion-policy", "", "")
	connectcampaignsv2_deleteConnectInstanceConfigCmd.Flags().String("connect-instance-id", "", "")
	connectcampaignsv2_deleteConnectInstanceConfigCmd.MarkFlagRequired("connect-instance-id")
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_deleteConnectInstanceConfigCmd)
}
