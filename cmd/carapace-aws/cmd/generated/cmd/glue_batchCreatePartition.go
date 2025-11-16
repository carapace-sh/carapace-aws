package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_batchCreatePartitionCmd = &cobra.Command{
	Use:   "batch-create-partition",
	Short: "Creates one or more partitions in a batch operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_batchCreatePartitionCmd).Standalone()

	glue_batchCreatePartitionCmd.Flags().String("catalog-id", "", "The ID of the catalog in which the partition is to be created.")
	glue_batchCreatePartitionCmd.Flags().String("database-name", "", "The name of the metadata database in which the partition is to be created.")
	glue_batchCreatePartitionCmd.Flags().String("partition-input-list", "", "A list of `PartitionInput` structures that define the partitions to be created.")
	glue_batchCreatePartitionCmd.Flags().String("table-name", "", "The name of the metadata table in which the partition is to be created.")
	glue_batchCreatePartitionCmd.MarkFlagRequired("database-name")
	glue_batchCreatePartitionCmd.MarkFlagRequired("partition-input-list")
	glue_batchCreatePartitionCmd.MarkFlagRequired("table-name")
	glueCmd.AddCommand(glue_batchCreatePartitionCmd)
}
