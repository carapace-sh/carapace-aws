package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_batchGetTableOptimizerCmd = &cobra.Command{
	Use:   "batch-get-table-optimizer",
	Short: "Returns the configuration for the specified table optimizers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_batchGetTableOptimizerCmd).Standalone()

	glue_batchGetTableOptimizerCmd.Flags().String("entries", "", "A list of `BatchGetTableOptimizerEntry` objects specifying the table optimizers to retrieve.")
	glue_batchGetTableOptimizerCmd.MarkFlagRequired("entries")
	glueCmd.AddCommand(glue_batchGetTableOptimizerCmd)
}
