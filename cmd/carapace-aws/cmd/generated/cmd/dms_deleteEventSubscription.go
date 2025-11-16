package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_deleteEventSubscriptionCmd = &cobra.Command{
	Use:   "delete-event-subscription",
	Short: "Deletes an DMS event subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_deleteEventSubscriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_deleteEventSubscriptionCmd).Standalone()

		dms_deleteEventSubscriptionCmd.Flags().String("subscription-name", "", "The name of the DMS event notification subscription to be deleted.")
		dms_deleteEventSubscriptionCmd.MarkFlagRequired("subscription-name")
	})
	dmsCmd.AddCommand(dms_deleteEventSubscriptionCmd)
}
