package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationAutoscaling_describeScheduledActionsCmd = &cobra.Command{
	Use:   "describe-scheduled-actions",
	Short: "Describes the Application Auto Scaling scheduled actions for the specified service namespace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationAutoscaling_describeScheduledActionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationAutoscaling_describeScheduledActionsCmd).Standalone()

		applicationAutoscaling_describeScheduledActionsCmd.Flags().String("max-results", "", "The maximum number of scheduled action results.")
		applicationAutoscaling_describeScheduledActionsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		applicationAutoscaling_describeScheduledActionsCmd.Flags().String("resource-id", "", "The identifier of the resource associated with the scheduled action.")
		applicationAutoscaling_describeScheduledActionsCmd.Flags().String("scalable-dimension", "", "The scalable dimension.")
		applicationAutoscaling_describeScheduledActionsCmd.Flags().String("scheduled-action-names", "", "The names of the scheduled actions to describe.")
		applicationAutoscaling_describeScheduledActionsCmd.Flags().String("service-namespace", "", "The namespace of the Amazon Web Services service that provides the resource.")
		applicationAutoscaling_describeScheduledActionsCmd.MarkFlagRequired("service-namespace")
	})
	applicationAutoscalingCmd.AddCommand(applicationAutoscaling_describeScheduledActionsCmd)
}
