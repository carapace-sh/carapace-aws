package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_getCalculationExecutionCodeCmd = &cobra.Command{
	Use:   "get-calculation-execution-code",
	Short: "Retrieves the unencrypted code that was executed for the calculation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_getCalculationExecutionCodeCmd).Standalone()

	athena_getCalculationExecutionCodeCmd.Flags().String("calculation-execution-id", "", "The calculation execution UUID.")
	athena_getCalculationExecutionCodeCmd.MarkFlagRequired("calculation-execution-id")
	athenaCmd.AddCommand(athena_getCalculationExecutionCodeCmd)
}
