package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repostspace_listChannelsCmd = &cobra.Command{
	Use:   "list-channels",
	Short: "Returns the list of channel within a private re:Post with some information about each channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repostspace_listChannelsCmd).Standalone()

	repostspace_listChannelsCmd.Flags().String("max-results", "", "The maximum number of channels to include in the results.")
	repostspace_listChannelsCmd.Flags().String("next-token", "", "The token for the next set of channel to return.")
	repostspace_listChannelsCmd.Flags().String("space-id", "", "The unique ID of the private re:Post.")
	repostspace_listChannelsCmd.MarkFlagRequired("space-id")
	repostspaceCmd.AddCommand(repostspace_listChannelsCmd)
}
