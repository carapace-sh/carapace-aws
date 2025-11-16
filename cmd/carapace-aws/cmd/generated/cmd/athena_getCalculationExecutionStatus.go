package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_getCalculationExecutionStatusCmd = &cobra.Command{
	Use:   "get-calculation-execution-status",
	Short: "Gets the status of a current calculation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_getCalculationExecutionStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(athena_getCalculationExecutionStatusCmd).Standalone()

		athena_getCalculationExecutionStatusCmd.Flags().String("calculation-execution-id", "", "The calculation execution UUID.")
		athena_getCalculationExecutionStatusCmd.MarkFlagRequired("calculation-execution-id")
	})
	athenaCmd.AddCommand(athena_getCalculationExecutionStatusCmd)
}
