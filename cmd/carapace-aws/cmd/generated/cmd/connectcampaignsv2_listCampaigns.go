package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_listCampaignsCmd = &cobra.Command{
	Use:   "list-campaigns",
	Short: "Provides summary information about the campaigns under the specified Amazon Connect account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_listCampaignsCmd).Standalone()

	connectcampaignsv2_listCampaignsCmd.Flags().String("filters", "", "")
	connectcampaignsv2_listCampaignsCmd.Flags().String("max-results", "", "")
	connectcampaignsv2_listCampaignsCmd.Flags().String("next-token", "", "")
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_listCampaignsCmd)
}
