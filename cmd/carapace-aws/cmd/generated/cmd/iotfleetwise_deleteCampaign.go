package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_deleteCampaignCmd = &cobra.Command{
	Use:   "delete-campaign",
	Short: "Deletes a data collection campaign.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_deleteCampaignCmd).Standalone()

	iotfleetwise_deleteCampaignCmd.Flags().String("name", "", "The name of the campaign to delete.")
	iotfleetwise_deleteCampaignCmd.MarkFlagRequired("name")
	iotfleetwiseCmd.AddCommand(iotfleetwise_deleteCampaignCmd)
}
