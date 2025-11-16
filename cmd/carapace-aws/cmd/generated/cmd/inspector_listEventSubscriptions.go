package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_listEventSubscriptionsCmd = &cobra.Command{
	Use:   "list-event-subscriptions",
	Short: "Lists all the event subscriptions for the assessment template that is specified by the ARN of the assessment template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_listEventSubscriptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector_listEventSubscriptionsCmd).Standalone()

		inspector_listEventSubscriptionsCmd.Flags().String("max-results", "", "You can use this parameter to indicate the maximum number of items you want in the response.")
		inspector_listEventSubscriptionsCmd.Flags().String("next-token", "", "You can use this parameter when paginating results.")
		inspector_listEventSubscriptionsCmd.Flags().String("resource-arn", "", "The ARN of the assessment template for which you want to list the existing event subscriptions.")
	})
	inspectorCmd.AddCommand(inspector_listEventSubscriptionsCmd)
}
