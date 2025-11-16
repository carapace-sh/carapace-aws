package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_batchGetPartitionCmd = &cobra.Command{
	Use:   "batch-get-partition",
	Short: "Retrieves partitions in a batch request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_batchGetPartitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_batchGetPartitionCmd).Standalone()

		glue_batchGetPartitionCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog where the partitions in question reside.")
		glue_batchGetPartitionCmd.Flags().String("database-name", "", "The name of the catalog database where the partitions reside.")
		glue_batchGetPartitionCmd.Flags().String("partitions-to-get", "", "A list of partition values identifying the partitions to retrieve.")
		glue_batchGetPartitionCmd.Flags().String("table-name", "", "The name of the partitions' table.")
		glue_batchGetPartitionCmd.MarkFlagRequired("database-name")
		glue_batchGetPartitionCmd.MarkFlagRequired("partitions-to-get")
		glue_batchGetPartitionCmd.MarkFlagRequired("table-name")
	})
	glueCmd.AddCommand(glue_batchGetPartitionCmd)
}
