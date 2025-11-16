package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_listChannelGroupsCmd = &cobra.Command{
	Use:   "list-channel-groups",
	Short: "Retrieves all channel groups that are configured in Elemental MediaPackage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_listChannelGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackagev2_listChannelGroupsCmd).Standalone()

		mediapackagev2_listChannelGroupsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		mediapackagev2_listChannelGroupsCmd.Flags().String("next-token", "", "The pagination token from the GET list request.")
	})
	mediapackagev2Cmd.AddCommand(mediapackagev2_listChannelGroupsCmd)
}
