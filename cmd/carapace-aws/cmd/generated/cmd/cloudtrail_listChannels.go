package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_listChannelsCmd = &cobra.Command{
	Use:   "list-channels",
	Short: "Lists the channels in the current account, and their source names.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_listChannelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudtrail_listChannelsCmd).Standalone()

		cloudtrail_listChannelsCmd.Flags().String("max-results", "", "The maximum number of CloudTrail channels to display on a single page.")
		cloudtrail_listChannelsCmd.Flags().String("next-token", "", "The token to use to get the next page of results after a previous API call.")
	})
	cloudtrailCmd.AddCommand(cloudtrail_listChannelsCmd)
}
