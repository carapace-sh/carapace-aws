package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getPartitionCmd = &cobra.Command{
	Use:   "get-partition",
	Short: "Retrieves information about a specified partition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getPartitionCmd).Standalone()

	glue_getPartitionCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog where the partition in question resides.")
	glue_getPartitionCmd.Flags().String("database-name", "", "The name of the catalog database where the partition resides.")
	glue_getPartitionCmd.Flags().String("partition-values", "", "The values that define the partition.")
	glue_getPartitionCmd.Flags().String("table-name", "", "The name of the partition's table.")
	glue_getPartitionCmd.MarkFlagRequired("database-name")
	glue_getPartitionCmd.MarkFlagRequired("partition-values")
	glue_getPartitionCmd.MarkFlagRequired("table-name")
	glueCmd.AddCommand(glue_getPartitionCmd)
}
