package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deletePartitionIndexCmd = &cobra.Command{
	Use:   "delete-partition-index",
	Short: "Deletes a specified partition index from an existing table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deletePartitionIndexCmd).Standalone()

	glue_deletePartitionIndexCmd.Flags().String("catalog-id", "", "The catalog ID where the table resides.")
	glue_deletePartitionIndexCmd.Flags().String("database-name", "", "Specifies the name of a database from which you want to delete a partition index.")
	glue_deletePartitionIndexCmd.Flags().String("index-name", "", "The name of the partition index to be deleted.")
	glue_deletePartitionIndexCmd.Flags().String("table-name", "", "Specifies the name of a table from which you want to delete a partition index.")
	glue_deletePartitionIndexCmd.MarkFlagRequired("database-name")
	glue_deletePartitionIndexCmd.MarkFlagRequired("index-name")
	glue_deletePartitionIndexCmd.MarkFlagRequired("table-name")
	glueCmd.AddCommand(glue_deletePartitionIndexCmd)
}
