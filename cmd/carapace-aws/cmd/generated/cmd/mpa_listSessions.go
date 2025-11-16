package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mpa_listSessionsCmd = &cobra.Command{
	Use:   "list-sessions",
	Short: "Returns a list of approval sessions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mpa_listSessionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mpa_listSessionsCmd).Standalone()

		mpa_listSessionsCmd.Flags().String("approval-team-arn", "", "Amazon Resource Name (ARN) for the approval team.")
		mpa_listSessionsCmd.Flags().String("filters", "", "An array of `Filter` objects.")
		mpa_listSessionsCmd.Flags().String("max-results", "", "The maximum number of items to return in the response.")
		mpa_listSessionsCmd.Flags().String("next-token", "", "If present, indicates that more output is available than is included in the current response.")
		mpa_listSessionsCmd.MarkFlagRequired("approval-team-arn")
	})
	mpaCmd.AddCommand(mpa_listSessionsCmd)
}
