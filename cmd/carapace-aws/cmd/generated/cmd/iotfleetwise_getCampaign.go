package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_getCampaignCmd = &cobra.Command{
	Use:   "get-campaign",
	Short: "Retrieves information about a campaign.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_getCampaignCmd).Standalone()

	iotfleetwise_getCampaignCmd.Flags().String("name", "", "The name of the campaign to retrieve information about.")
	iotfleetwise_getCampaignCmd.MarkFlagRequired("name")
	iotfleetwiseCmd.AddCommand(iotfleetwise_getCampaignCmd)
}
