package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_listCampaignsCmd = &cobra.Command{
	Use:   "list-campaigns",
	Short: "Lists information about created campaigns.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_listCampaignsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotfleetwise_listCampaignsCmd).Standalone()

		iotfleetwise_listCampaignsCmd.Flags().String("list-response-scope", "", "When you set the `listResponseScope` parameter to `METADATA_ONLY`, the list response includes: campaign name, Amazon Resource Name (ARN), creation time, and last modification time.")
		iotfleetwise_listCampaignsCmd.Flags().String("max-results", "", "The maximum number of items to return, between 1 and 100, inclusive.")
		iotfleetwise_listCampaignsCmd.Flags().String("next-token", "", "A pagination token for the next set of results.")
		iotfleetwise_listCampaignsCmd.Flags().String("status", "", "An optional parameter to filter the results by the status of each created campaign in your account.")
	})
	iotfleetwiseCmd.AddCommand(iotfleetwise_listCampaignsCmd)
}
