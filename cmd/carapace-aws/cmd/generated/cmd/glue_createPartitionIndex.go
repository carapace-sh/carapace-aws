package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_createPartitionIndexCmd = &cobra.Command{
	Use:   "create-partition-index",
	Short: "Creates a specified partition index in an existing table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_createPartitionIndexCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_createPartitionIndexCmd).Standalone()

		glue_createPartitionIndexCmd.Flags().String("catalog-id", "", "The catalog ID where the table resides.")
		glue_createPartitionIndexCmd.Flags().String("database-name", "", "Specifies the name of a database in which you want to create a partition index.")
		glue_createPartitionIndexCmd.Flags().String("partition-index", "", "Specifies a `PartitionIndex` structure to create a partition index in an existing table.")
		glue_createPartitionIndexCmd.Flags().String("table-name", "", "Specifies the name of a table in which you want to create a partition index.")
		glue_createPartitionIndexCmd.MarkFlagRequired("database-name")
		glue_createPartitionIndexCmd.MarkFlagRequired("partition-index")
		glue_createPartitionIndexCmd.MarkFlagRequired("table-name")
	})
	glueCmd.AddCommand(glue_createPartitionIndexCmd)
}
