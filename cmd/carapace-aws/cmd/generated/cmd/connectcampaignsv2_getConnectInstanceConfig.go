package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_getConnectInstanceConfigCmd = &cobra.Command{
	Use:   "get-connect-instance-config",
	Short: "Get the specific Connect instance config.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_getConnectInstanceConfigCmd).Standalone()

	connectcampaignsv2_getConnectInstanceConfigCmd.Flags().String("connect-instance-id", "", "")
	connectcampaignsv2_getConnectInstanceConfigCmd.MarkFlagRequired("connect-instance-id")
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_getConnectInstanceConfigCmd)
}
