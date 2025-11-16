package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getPartitionsCmd = &cobra.Command{
	Use:   "get-partitions",
	Short: "Retrieves information about the partitions in a table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getPartitionsCmd).Standalone()

	glue_getPartitionsCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog where the partitions in question reside.")
	glue_getPartitionsCmd.Flags().String("database-name", "", "The name of the catalog database where the partitions reside.")
	glue_getPartitionsCmd.Flags().String("exclude-column-schema", "", "When true, specifies not returning the partition column schema.")
	glue_getPartitionsCmd.Flags().String("expression", "", "An expression that filters the partitions to be returned.")
	glue_getPartitionsCmd.Flags().String("max-results", "", "The maximum number of partitions to return in a single response.")
	glue_getPartitionsCmd.Flags().String("next-token", "", "A continuation token, if this is not the first call to retrieve these partitions.")
	glue_getPartitionsCmd.Flags().String("query-as-of-time", "", "The time as of when to read the partition contents.")
	glue_getPartitionsCmd.Flags().String("segment", "", "The segment of the table's partitions to scan in this request.")
	glue_getPartitionsCmd.Flags().String("table-name", "", "The name of the partitions' table.")
	glue_getPartitionsCmd.Flags().String("transaction-id", "", "The transaction ID at which to read the partition contents.")
	glue_getPartitionsCmd.MarkFlagRequired("database-name")
	glue_getPartitionsCmd.MarkFlagRequired("table-name")
	glueCmd.AddCommand(glue_getPartitionsCmd)
}
