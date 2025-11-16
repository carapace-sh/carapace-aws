package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getPartitionIndexesCmd = &cobra.Command{
	Use:   "get-partition-indexes",
	Short: "Retrieves the partition indexes associated with a table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getPartitionIndexesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getPartitionIndexesCmd).Standalone()

		glue_getPartitionIndexesCmd.Flags().String("catalog-id", "", "The catalog ID where the table resides.")
		glue_getPartitionIndexesCmd.Flags().String("database-name", "", "Specifies the name of a database from which you want to retrieve partition indexes.")
		glue_getPartitionIndexesCmd.Flags().String("next-token", "", "A continuation token, included if this is a continuation call.")
		glue_getPartitionIndexesCmd.Flags().String("table-name", "", "Specifies the name of a table for which you want to retrieve the partition indexes.")
		glue_getPartitionIndexesCmd.MarkFlagRequired("database-name")
		glue_getPartitionIndexesCmd.MarkFlagRequired("table-name")
	})
	glueCmd.AddCommand(glue_getPartitionIndexesCmd)
}
