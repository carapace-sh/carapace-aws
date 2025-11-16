package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationAutoscaling_deleteScheduledActionCmd = &cobra.Command{
	Use:   "delete-scheduled-action",
	Short: "Deletes the specified scheduled action for an Application Auto Scaling scalable target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationAutoscaling_deleteScheduledActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationAutoscaling_deleteScheduledActionCmd).Standalone()

		applicationAutoscaling_deleteScheduledActionCmd.Flags().String("resource-id", "", "The identifier of the resource associated with the scheduled action.")
		applicationAutoscaling_deleteScheduledActionCmd.Flags().String("scalable-dimension", "", "The scalable dimension.")
		applicationAutoscaling_deleteScheduledActionCmd.Flags().String("scheduled-action-name", "", "The name of the scheduled action.")
		applicationAutoscaling_deleteScheduledActionCmd.Flags().String("service-namespace", "", "The namespace of the Amazon Web Services service that provides the resource.")
		applicationAutoscaling_deleteScheduledActionCmd.MarkFlagRequired("resource-id")
		applicationAutoscaling_deleteScheduledActionCmd.MarkFlagRequired("scalable-dimension")
		applicationAutoscaling_deleteScheduledActionCmd.MarkFlagRequired("scheduled-action-name")
		applicationAutoscaling_deleteScheduledActionCmd.MarkFlagRequired("service-namespace")
	})
	applicationAutoscalingCmd.AddCommand(applicationAutoscaling_deleteScheduledActionCmd)
}
