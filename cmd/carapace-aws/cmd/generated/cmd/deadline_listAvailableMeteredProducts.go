package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_listAvailableMeteredProductsCmd = &cobra.Command{
	Use:   "list-available-metered-products",
	Short: "A list of the available metered products.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_listAvailableMeteredProductsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_listAvailableMeteredProductsCmd).Standalone()

		deadline_listAvailableMeteredProductsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		deadline_listAvailableMeteredProductsCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` to start from the beginning.")
	})
	deadlineCmd.AddCommand(deadline_listAvailableMeteredProductsCmd)
}
