package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_createAppCmd = &cobra.Command{
	Use:   "create-app",
	Short: "Creates an Resilience Hub application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_createAppCmd).Standalone()

	resiliencehub_createAppCmd.Flags().String("assessment-schedule", "", "Assessment execution schedule with 'Daily' or 'Disabled' values.")
	resiliencehub_createAppCmd.Flags().String("aws-application-arn", "", "Amazon Resource Name (ARN) of Resource Groups group that is integrated with an AppRegistry application.")
	resiliencehub_createAppCmd.Flags().String("client-token", "", "Used for an idempotency token.")
	resiliencehub_createAppCmd.Flags().String("description", "", "The optional description for an app.")
	resiliencehub_createAppCmd.Flags().String("event-subscriptions", "", "The list of events you would like to subscribe and get notification for.")
	resiliencehub_createAppCmd.Flags().String("name", "", "Name of the application.")
	resiliencehub_createAppCmd.Flags().String("permission-model", "", "Defines the roles and credentials that Resilience Hub would use while creating the application, importing its resources, and running an assessment.")
	resiliencehub_createAppCmd.Flags().String("policy-arn", "", "Amazon Resource Name (ARN) of the resiliency policy.")
	resiliencehub_createAppCmd.Flags().String("tags", "", "Tags assigned to the resource.")
	resiliencehub_createAppCmd.MarkFlagRequired("name")
	resiliencehubCmd.AddCommand(resiliencehub_createAppCmd)
}
