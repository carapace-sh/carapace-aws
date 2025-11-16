package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_stopCalculationExecutionCmd = &cobra.Command{
	Use:   "stop-calculation-execution",
	Short: "Requests the cancellation of a calculation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_stopCalculationExecutionCmd).Standalone()

	athena_stopCalculationExecutionCmd.Flags().String("calculation-execution-id", "", "The calculation execution UUID.")
	athena_stopCalculationExecutionCmd.MarkFlagRequired("calculation-execution-id")
	athenaCmd.AddCommand(athena_stopCalculationExecutionCmd)
}
