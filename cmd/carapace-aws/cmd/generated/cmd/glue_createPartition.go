package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_createPartitionCmd = &cobra.Command{
	Use:   "create-partition",
	Short: "Creates a new partition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_createPartitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_createPartitionCmd).Standalone()

		glue_createPartitionCmd.Flags().String("catalog-id", "", "The Amazon Web Services account ID of the catalog in which the partition is to be created.")
		glue_createPartitionCmd.Flags().String("database-name", "", "The name of the metadata database in which the partition is to be created.")
		glue_createPartitionCmd.Flags().String("partition-input", "", "A `PartitionInput` structure defining the partition to be created.")
		glue_createPartitionCmd.Flags().String("table-name", "", "The name of the metadata table in which the partition is to be created.")
		glue_createPartitionCmd.MarkFlagRequired("database-name")
		glue_createPartitionCmd.MarkFlagRequired("partition-input")
		glue_createPartitionCmd.MarkFlagRequired("table-name")
	})
	glueCmd.AddCommand(glue_createPartitionCmd)
}
