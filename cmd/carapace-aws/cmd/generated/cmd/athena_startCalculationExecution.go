package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_startCalculationExecutionCmd = &cobra.Command{
	Use:   "start-calculation-execution",
	Short: "Submits calculations for execution within a session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_startCalculationExecutionCmd).Standalone()

	athena_startCalculationExecutionCmd.Flags().String("calculation-configuration", "", "Contains configuration information for the calculation.")
	athena_startCalculationExecutionCmd.Flags().String("client-request-token", "", "A unique case-sensitive string used to ensure the request to create the calculation is idempotent (executes only once).")
	athena_startCalculationExecutionCmd.Flags().String("code-block", "", "A string that contains the code of the calculation.")
	athena_startCalculationExecutionCmd.Flags().String("description", "", "A description of the calculation.")
	athena_startCalculationExecutionCmd.Flags().String("session-id", "", "The session ID.")
	athena_startCalculationExecutionCmd.MarkFlagRequired("session-id")
	athenaCmd.AddCommand(athena_startCalculationExecutionCmd)
}
