package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgh_listProgressUpdateStreamsCmd = &cobra.Command{
	Use:   "list-progress-update-streams",
	Short: "Lists progress update streams associated with the user account making this call.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgh_listProgressUpdateStreamsCmd).Standalone()

	mgh_listProgressUpdateStreamsCmd.Flags().String("max-results", "", "Filter to limit the maximum number of results to list per page.")
	mgh_listProgressUpdateStreamsCmd.Flags().String("next-token", "", "If a `NextToken` was returned by a previous call, there are more results available.")
	mghCmd.AddCommand(mgh_listProgressUpdateStreamsCmd)
}
