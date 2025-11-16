package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_getTableObjectsCmd = &cobra.Command{
	Use:   "get-table-objects",
	Short: "Returns the set of Amazon S3 objects that make up the specified governed table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_getTableObjectsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lakeformation_getTableObjectsCmd).Standalone()

		lakeformation_getTableObjectsCmd.Flags().String("catalog-id", "", "The catalog containing the governed table.")
		lakeformation_getTableObjectsCmd.Flags().String("database-name", "", "The database containing the governed table.")
		lakeformation_getTableObjectsCmd.Flags().String("max-results", "", "Specifies how many values to return in a page.")
		lakeformation_getTableObjectsCmd.Flags().String("next-token", "", "A continuation token if this is not the first call to retrieve these objects.")
		lakeformation_getTableObjectsCmd.Flags().String("partition-predicate", "", "A predicate to filter the objects returned based on the partition keys defined in the governed table.")
		lakeformation_getTableObjectsCmd.Flags().String("query-as-of-time", "", "The time as of when to read the governed table contents.")
		lakeformation_getTableObjectsCmd.Flags().String("table-name", "", "The governed table for which to retrieve objects.")
		lakeformation_getTableObjectsCmd.Flags().String("transaction-id", "", "The transaction ID at which to read the governed table contents.")
		lakeformation_getTableObjectsCmd.MarkFlagRequired("database-name")
		lakeformation_getTableObjectsCmd.MarkFlagRequired("table-name")
	})
	lakeformationCmd.AddCommand(lakeformation_getTableObjectsCmd)
}
