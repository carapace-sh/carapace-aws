package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_updateTableObjectsCmd = &cobra.Command{
	Use:   "update-table-objects",
	Short: "Updates the manifest of Amazon S3 objects that make up the specified governed table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_updateTableObjectsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lakeformation_updateTableObjectsCmd).Standalone()

		lakeformation_updateTableObjectsCmd.Flags().String("catalog-id", "", "The catalog containing the governed table to update.")
		lakeformation_updateTableObjectsCmd.Flags().String("database-name", "", "The database containing the governed table to update.")
		lakeformation_updateTableObjectsCmd.Flags().String("table-name", "", "The governed table to update.")
		lakeformation_updateTableObjectsCmd.Flags().String("transaction-id", "", "The transaction at which to do the write.")
		lakeformation_updateTableObjectsCmd.Flags().String("write-operations", "", "A list of `WriteOperation` objects that define an object to add to or delete from the manifest for a governed table.")
		lakeformation_updateTableObjectsCmd.MarkFlagRequired("database-name")
		lakeformation_updateTableObjectsCmd.MarkFlagRequired("table-name")
		lakeformation_updateTableObjectsCmd.MarkFlagRequired("write-operations")
	})
	lakeformationCmd.AddCommand(lakeformation_updateTableObjectsCmd)
}
