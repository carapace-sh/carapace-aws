package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_listChannelsCmd = &cobra.Command{
	Use:   "list-channels",
	Short: "Retrieves all channels in a specific channel group that are configured in AWS Elemental MediaPackage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_listChannelsCmd).Standalone()

	mediapackagev2_listChannelsCmd.Flags().String("channel-group-name", "", "The name that describes the channel group.")
	mediapackagev2_listChannelsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	mediapackagev2_listChannelsCmd.Flags().String("next-token", "", "The pagination token from the GET list request.")
	mediapackagev2_listChannelsCmd.MarkFlagRequired("channel-group-name")
	mediapackagev2Cmd.AddCommand(mediapackagev2_listChannelsCmd)
}
