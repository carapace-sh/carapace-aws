package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describeAutomationExecutionsCmd = &cobra.Command{
	Use:   "describe-automation-executions",
	Short: "Provides details about all active and terminated Automation executions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describeAutomationExecutionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_describeAutomationExecutionsCmd).Standalone()

		ssm_describeAutomationExecutionsCmd.Flags().String("filters", "", "Filters used to limit the scope of executions that are requested.")
		ssm_describeAutomationExecutionsCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
		ssm_describeAutomationExecutionsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	})
	ssmCmd.AddCommand(ssm_describeAutomationExecutionsCmd)
}
