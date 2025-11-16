package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_updateSubscriptionsToEventBridgeCmd = &cobra.Command{
	Use:   "update-subscriptions-to-event-bridge",
	Short: "Migrates 10 active and enabled Amazon SNS subscriptions at a time and converts them to corresponding Amazon EventBridge rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_updateSubscriptionsToEventBridgeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_updateSubscriptionsToEventBridgeCmd).Standalone()

		dms_updateSubscriptionsToEventBridgeCmd.Flags().String("force-move", "", "When set to true, this operation migrates DMS subscriptions for Amazon SNS notifications no matter what your replication instance version is.")
	})
	dmsCmd.AddCommand(dms_updateSubscriptionsToEventBridgeCmd)
}
