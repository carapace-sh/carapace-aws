package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_updateCampaignCmd = &cobra.Command{
	Use:   "update-campaign",
	Short: "Updates a campaign.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_updateCampaignCmd).Standalone()

	iotfleetwise_updateCampaignCmd.Flags().String("action", "", "Specifies how to update a campaign.")
	iotfleetwise_updateCampaignCmd.Flags().String("data-extra-dimensions", "", "A list of vehicle attributes to associate with a signal.")
	iotfleetwise_updateCampaignCmd.Flags().String("description", "", "The description of the campaign.")
	iotfleetwise_updateCampaignCmd.Flags().String("name", "", "The name of the campaign to update.")
	iotfleetwise_updateCampaignCmd.MarkFlagRequired("action")
	iotfleetwise_updateCampaignCmd.MarkFlagRequired("name")
	iotfleetwiseCmd.AddCommand(iotfleetwise_updateCampaignCmd)
}
