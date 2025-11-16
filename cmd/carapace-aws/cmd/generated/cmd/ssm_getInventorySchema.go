package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_getInventorySchemaCmd = &cobra.Command{
	Use:   "get-inventory-schema",
	Short: "Return a list of inventory type names for the account, or return a list of attribute names for a specific Inventory item type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_getInventorySchemaCmd).Standalone()

	ssm_getInventorySchemaCmd.Flags().String("aggregator", "", "Returns inventory schemas that support aggregation.")
	ssm_getInventorySchemaCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	ssm_getInventorySchemaCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	ssm_getInventorySchemaCmd.Flags().String("sub-type", "", "Returns the sub-type schema for a specified inventory type.")
	ssm_getInventorySchemaCmd.Flags().String("type-name", "", "The type of inventory item to return.")
	ssmCmd.AddCommand(ssm_getInventorySchemaCmd)
}
