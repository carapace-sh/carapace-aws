package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_updateAppCmd = &cobra.Command{
	Use:   "update-app",
	Short: "Updates an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_updateAppCmd).Standalone()

	resiliencehub_updateAppCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
	resiliencehub_updateAppCmd.Flags().String("assessment-schedule", "", "Assessment execution schedule with 'Daily' or 'Disabled' values.")
	resiliencehub_updateAppCmd.Flags().String("clear-resiliency-policy-arn", "", "Specifies if the resiliency policy ARN should be cleared.")
	resiliencehub_updateAppCmd.Flags().String("description", "", "The optional description for an app.")
	resiliencehub_updateAppCmd.Flags().String("event-subscriptions", "", "The list of events you would like to subscribe and get notification for.")
	resiliencehub_updateAppCmd.Flags().String("permission-model", "", "Defines the roles and credentials that Resilience Hub would use while creating an application, importing its resources, and running an assessment.")
	resiliencehub_updateAppCmd.Flags().String("policy-arn", "", "Amazon Resource Name (ARN) of the resiliency policy.")
	resiliencehub_updateAppCmd.MarkFlagRequired("app-arn")
	resiliencehubCmd.AddCommand(resiliencehub_updateAppCmd)
}
