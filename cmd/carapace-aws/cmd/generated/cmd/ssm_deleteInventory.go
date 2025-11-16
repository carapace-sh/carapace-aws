package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_deleteInventoryCmd = &cobra.Command{
	Use:   "delete-inventory",
	Short: "Delete a custom inventory type or the data associated with a custom Inventory type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_deleteInventoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_deleteInventoryCmd).Standalone()

		ssm_deleteInventoryCmd.Flags().String("client-token", "", "User-provided idempotency token.")
		ssm_deleteInventoryCmd.Flags().String("dry-run", "", "Use this option to view a summary of the deletion request without deleting any data or the data type.")
		ssm_deleteInventoryCmd.Flags().String("schema-delete-option", "", "Use the `SchemaDeleteOption` to delete a custom inventory type (schema).")
		ssm_deleteInventoryCmd.Flags().String("type-name", "", "The name of the custom inventory type for which you want to delete either all previously collected data or the inventory type itself.")
		ssm_deleteInventoryCmd.MarkFlagRequired("type-name")
	})
	ssmCmd.AddCommand(ssm_deleteInventoryCmd)
}
