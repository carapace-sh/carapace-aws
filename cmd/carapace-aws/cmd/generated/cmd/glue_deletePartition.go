package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deletePartitionCmd = &cobra.Command{
	Use:   "delete-partition",
	Short: "Deletes a specified partition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deletePartitionCmd).Standalone()

	glue_deletePartitionCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog where the partition to be deleted resides.")
	glue_deletePartitionCmd.Flags().String("database-name", "", "The name of the catalog database in which the table in question resides.")
	glue_deletePartitionCmd.Flags().String("partition-values", "", "The values that define the partition.")
	glue_deletePartitionCmd.Flags().String("table-name", "", "The name of the table that contains the partition to be deleted.")
	glue_deletePartitionCmd.MarkFlagRequired("database-name")
	glue_deletePartitionCmd.MarkFlagRequired("partition-values")
	glue_deletePartitionCmd.MarkFlagRequired("table-name")
	glueCmd.AddCommand(glue_deletePartitionCmd)
}
