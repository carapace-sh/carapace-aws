package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_getCalculationExecutionCmd = &cobra.Command{
	Use:   "get-calculation-execution",
	Short: "Describes a previously submitted calculation execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_getCalculationExecutionCmd).Standalone()

	athena_getCalculationExecutionCmd.Flags().String("calculation-execution-id", "", "The calculation execution UUID.")
	athena_getCalculationExecutionCmd.MarkFlagRequired("calculation-execution-id")
	athenaCmd.AddCommand(athena_getCalculationExecutionCmd)
}
