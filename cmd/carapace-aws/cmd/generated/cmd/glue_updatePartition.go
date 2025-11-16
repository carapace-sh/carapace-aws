package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_updatePartitionCmd = &cobra.Command{
	Use:   "update-partition",
	Short: "Updates a partition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_updatePartitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_updatePartitionCmd).Standalone()

		glue_updatePartitionCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog where the partition to be updated resides.")
		glue_updatePartitionCmd.Flags().String("database-name", "", "The name of the catalog database in which the table in question resides.")
		glue_updatePartitionCmd.Flags().String("partition-input", "", "The new partition object to update the partition to.")
		glue_updatePartitionCmd.Flags().String("partition-value-list", "", "List of partition key values that define the partition to update.")
		glue_updatePartitionCmd.Flags().String("table-name", "", "The name of the table in which the partition to be updated is located.")
		glue_updatePartitionCmd.MarkFlagRequired("database-name")
		glue_updatePartitionCmd.MarkFlagRequired("partition-input")
		glue_updatePartitionCmd.MarkFlagRequired("partition-value-list")
		glue_updatePartitionCmd.MarkFlagRequired("table-name")
	})
	glueCmd.AddCommand(glue_updatePartitionCmd)
}
