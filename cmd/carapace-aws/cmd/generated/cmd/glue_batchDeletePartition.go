package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_batchDeletePartitionCmd = &cobra.Command{
	Use:   "batch-delete-partition",
	Short: "Deletes one or more partitions in a batch operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_batchDeletePartitionCmd).Standalone()

	glue_batchDeletePartitionCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog where the partition to be deleted resides.")
	glue_batchDeletePartitionCmd.Flags().String("database-name", "", "The name of the catalog database in which the table in question resides.")
	glue_batchDeletePartitionCmd.Flags().String("partitions-to-delete", "", "A list of `PartitionInput` structures that define the partitions to be deleted.")
	glue_batchDeletePartitionCmd.Flags().String("table-name", "", "The name of the table that contains the partitions to be deleted.")
	glue_batchDeletePartitionCmd.MarkFlagRequired("database-name")
	glue_batchDeletePartitionCmd.MarkFlagRequired("partitions-to-delete")
	glue_batchDeletePartitionCmd.MarkFlagRequired("table-name")
	glueCmd.AddCommand(glue_batchDeletePartitionCmd)
}
