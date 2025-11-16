package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_listPublishingDestinationsCmd = &cobra.Command{
	Use:   "list-publishing-destinations",
	Short: "Returns a list of publishing destinations associated with the specified `detectorId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_listPublishingDestinationsCmd).Standalone()

	guardduty_listPublishingDestinationsCmd.Flags().String("detector-id", "", "The detector ID for which you want to retrieve the publishing destination.")
	guardduty_listPublishingDestinationsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	guardduty_listPublishingDestinationsCmd.Flags().String("next-token", "", "A token to use for paginating results that are returned in the response.")
	guardduty_listPublishingDestinationsCmd.MarkFlagRequired("detector-id")
	guarddutyCmd.AddCommand(guardduty_listPublishingDestinationsCmd)
}
