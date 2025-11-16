package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_putLifecycleEventHookExecutionStatusCmd = &cobra.Command{
	Use:   "put-lifecycle-event-hook-execution-status",
	Short: "Sets the result of a Lambda validation function.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_putLifecycleEventHookExecutionStatusCmd).Standalone()

	codedeploy_putLifecycleEventHookExecutionStatusCmd.Flags().String("deployment-id", "", "The unique ID of a deployment.")
	codedeploy_putLifecycleEventHookExecutionStatusCmd.Flags().String("lifecycle-event-hook-execution-id", "", "The execution ID of a deployment's lifecycle hook.")
	codedeploy_putLifecycleEventHookExecutionStatusCmd.Flags().String("status", "", "The result of a Lambda function that validates a deployment lifecycle event.")
	codedeployCmd.AddCommand(codedeploy_putLifecycleEventHookExecutionStatusCmd)
}
