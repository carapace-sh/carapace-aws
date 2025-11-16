package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmDataExports_getExecutionCmd = &cobra.Command{
	Use:   "get-execution",
	Short: "Exports data based on the source data update.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmDataExports_getExecutionCmd).Standalone()

	bcmDataExports_getExecutionCmd.Flags().String("execution-id", "", "The ID for this specific execution.")
	bcmDataExports_getExecutionCmd.Flags().String("export-arn", "", "The Amazon Resource Name (ARN) of the Export object that generated this specific execution.")
	bcmDataExports_getExecutionCmd.MarkFlagRequired("execution-id")
	bcmDataExports_getExecutionCmd.MarkFlagRequired("export-arn")
	bcmDataExportsCmd.AddCommand(bcmDataExports_getExecutionCmd)
}
