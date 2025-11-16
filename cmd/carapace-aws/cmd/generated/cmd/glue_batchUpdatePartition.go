package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_batchUpdatePartitionCmd = &cobra.Command{
	Use:   "batch-update-partition",
	Short: "Updates one or more partitions in a batch operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_batchUpdatePartitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_batchUpdatePartitionCmd).Standalone()

		glue_batchUpdatePartitionCmd.Flags().String("catalog-id", "", "The ID of the catalog in which the partition is to be updated.")
		glue_batchUpdatePartitionCmd.Flags().String("database-name", "", "The name of the metadata database in which the partition is to be updated.")
		glue_batchUpdatePartitionCmd.Flags().String("entries", "", "A list of up to 100 `BatchUpdatePartitionRequestEntry` objects to update.")
		glue_batchUpdatePartitionCmd.Flags().String("table-name", "", "The name of the metadata table in which the partition is to be updated.")
		glue_batchUpdatePartitionCmd.MarkFlagRequired("database-name")
		glue_batchUpdatePartitionCmd.MarkFlagRequired("entries")
		glue_batchUpdatePartitionCmd.MarkFlagRequired("table-name")
	})
	glueCmd.AddCommand(glue_batchUpdatePartitionCmd)
}
