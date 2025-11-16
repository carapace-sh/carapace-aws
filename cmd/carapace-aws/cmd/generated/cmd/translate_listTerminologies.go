package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var translate_listTerminologiesCmd = &cobra.Command{
	Use:   "list-terminologies",
	Short: "Provides a list of custom terminologies associated with your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(translate_listTerminologiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(translate_listTerminologiesCmd).Standalone()

		translate_listTerminologiesCmd.Flags().String("max-results", "", "The maximum number of custom terminologies returned per list request.")
		translate_listTerminologiesCmd.Flags().String("next-token", "", "If the result of the request to ListTerminologies was truncated, include the NextToken to fetch the next group of custom terminologies.")
	})
	translateCmd.AddCommand(translate_listTerminologiesCmd)
}
