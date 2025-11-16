package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaigns_listCampaignsCmd = &cobra.Command{
	Use:   "list-campaigns",
	Short: "Provides summary information about the campaigns under the specified Amazon Connect account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaigns_listCampaignsCmd).Standalone()

	connectcampaigns_listCampaignsCmd.Flags().String("filters", "", "")
	connectcampaigns_listCampaignsCmd.Flags().String("max-results", "", "")
	connectcampaigns_listCampaignsCmd.Flags().String("next-token", "", "")
	connectcampaignsCmd.AddCommand(connectcampaigns_listCampaignsCmd)
}
