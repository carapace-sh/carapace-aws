package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_getInventoryCmd = &cobra.Command{
	Use:   "get-inventory",
	Short: "Query inventory information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_getInventoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_getInventoryCmd).Standalone()

		ssm_getInventoryCmd.Flags().String("aggregators", "", "Returns counts of inventory types based on one or more expressions.")
		ssm_getInventoryCmd.Flags().String("filters", "", "One or more filters.")
		ssm_getInventoryCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
		ssm_getInventoryCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		ssm_getInventoryCmd.Flags().String("result-attributes", "", "The list of inventory item types to return.")
	})
	ssmCmd.AddCommand(ssm_getInventoryCmd)
}
