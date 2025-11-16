package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackage_listChannelsCmd = &cobra.Command{
	Use:   "list-channels",
	Short: "Returns a collection of Channels.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackage_listChannelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackage_listChannelsCmd).Standalone()

		mediapackage_listChannelsCmd.Flags().String("max-results", "", "Upper bound on number of records to return.")
		mediapackage_listChannelsCmd.Flags().String("next-token", "", "A token used to resume pagination from the end of a previous request.")
	})
	mediapackageCmd.AddCommand(mediapackage_listChannelsCmd)
}
