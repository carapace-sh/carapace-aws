package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_listSegmentsCmd = &cobra.Command{
	Use:   "list-segments",
	Short: "Returns a list of audience segments that you have created in your account in this Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_listSegmentsCmd).Standalone()

	evidently_listSegmentsCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
	evidently_listSegmentsCmd.Flags().String("next-token", "", "The token to use when requesting the next set of results.")
	evidentlyCmd.AddCommand(evidently_listSegmentsCmd)
}
