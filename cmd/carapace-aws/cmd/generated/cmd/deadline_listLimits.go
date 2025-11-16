package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_listLimitsCmd = &cobra.Command{
	Use:   "list-limits",
	Short: "Gets a list of limits defined in the specified farm.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_listLimitsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_listLimitsCmd).Standalone()

		deadline_listLimitsCmd.Flags().String("farm-id", "", "The unique identifier of the farm that contains the limits.")
		deadline_listLimitsCmd.Flags().String("max-results", "", "The maximum number of limits to return in each page of results.")
		deadline_listLimitsCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` to start from the beginning.")
		deadline_listLimitsCmd.MarkFlagRequired("farm-id")
	})
	deadlineCmd.AddCommand(deadline_listLimitsCmd)
}
