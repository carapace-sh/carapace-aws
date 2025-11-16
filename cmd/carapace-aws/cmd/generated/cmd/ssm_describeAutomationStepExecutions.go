package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describeAutomationStepExecutionsCmd = &cobra.Command{
	Use:   "describe-automation-step-executions",
	Short: "Information about all active and terminated step executions in an Automation workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describeAutomationStepExecutionsCmd).Standalone()

	ssm_describeAutomationStepExecutionsCmd.Flags().String("automation-execution-id", "", "The Automation execution ID for which you want step execution descriptions.")
	ssm_describeAutomationStepExecutionsCmd.Flags().String("filters", "", "One or more filters to limit the number of step executions returned by the request.")
	ssm_describeAutomationStepExecutionsCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	ssm_describeAutomationStepExecutionsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	ssm_describeAutomationStepExecutionsCmd.Flags().Bool("no-reverse-order", false, "Indicates whether to list step executions in reverse order by start time.")
	ssm_describeAutomationStepExecutionsCmd.Flags().Bool("reverse-order", false, "Indicates whether to list step executions in reverse order by start time.")
	ssm_describeAutomationStepExecutionsCmd.MarkFlagRequired("automation-execution-id")
	ssm_describeAutomationStepExecutionsCmd.Flag("no-reverse-order").Hidden = true
	ssmCmd.AddCommand(ssm_describeAutomationStepExecutionsCmd)
}
