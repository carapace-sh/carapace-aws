package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_listMeteredProductsCmd = &cobra.Command{
	Use:   "list-metered-products",
	Short: "Lists metered products.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_listMeteredProductsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_listMeteredProductsCmd).Standalone()

		deadline_listMeteredProductsCmd.Flags().String("license-endpoint-id", "", "The license endpoint ID to include on the list of metered products.")
		deadline_listMeteredProductsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		deadline_listMeteredProductsCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` to start from the beginning.")
		deadline_listMeteredProductsCmd.MarkFlagRequired("license-endpoint-id")
	})
	deadlineCmd.AddCommand(deadline_listMeteredProductsCmd)
}
