package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_deleteObjectsOnCancelCmd = &cobra.Command{
	Use:   "delete-objects-on-cancel",
	Short: "For a specific governed table, provides a list of Amazon S3 objects that will be written during the current transaction and that can be automatically deleted if the transaction is canceled.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_deleteObjectsOnCancelCmd).Standalone()

	lakeformation_deleteObjectsOnCancelCmd.Flags().String("catalog-id", "", "The Glue data catalog that contains the governed table.")
	lakeformation_deleteObjectsOnCancelCmd.Flags().String("database-name", "", "The database that contains the governed table.")
	lakeformation_deleteObjectsOnCancelCmd.Flags().String("objects", "", "A list of VirtualObject structures, which indicates the Amazon S3 objects to be deleted if the transaction cancels.")
	lakeformation_deleteObjectsOnCancelCmd.Flags().String("table-name", "", "The name of the governed table.")
	lakeformation_deleteObjectsOnCancelCmd.Flags().String("transaction-id", "", "ID of the transaction that the writes occur in.")
	lakeformation_deleteObjectsOnCancelCmd.MarkFlagRequired("database-name")
	lakeformation_deleteObjectsOnCancelCmd.MarkFlagRequired("objects")
	lakeformation_deleteObjectsOnCancelCmd.MarkFlagRequired("table-name")
	lakeformation_deleteObjectsOnCancelCmd.MarkFlagRequired("transaction-id")
	lakeformationCmd.AddCommand(lakeformation_deleteObjectsOnCancelCmd)
}
