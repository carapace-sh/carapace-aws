package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_listChannelsCmd = &cobra.Command{
	Use:   "list-channels",
	Short: "Retrieves information about the channels that are associated with the current AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_listChannelsCmd).Standalone()

	mediatailor_listChannelsCmd.Flags().String("max-results", "", "The maximum number of channels that you want MediaTailor to return in response to the current request.")
	mediatailor_listChannelsCmd.Flags().String("next-token", "", "Pagination token returned by the list request when results exceed the maximum allowed.")
	mediatailorCmd.AddCommand(mediatailor_listChannelsCmd)
}
