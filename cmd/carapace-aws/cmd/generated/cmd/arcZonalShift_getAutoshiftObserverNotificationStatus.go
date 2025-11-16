package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcZonalShift_getAutoshiftObserverNotificationStatusCmd = &cobra.Command{
	Use:   "get-autoshift-observer-notification-status",
	Short: "Returns the status of the autoshift observer notification.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcZonalShift_getAutoshiftObserverNotificationStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(arcZonalShift_getAutoshiftObserverNotificationStatusCmd).Standalone()

	})
	arcZonalShiftCmd.AddCommand(arcZonalShift_getAutoshiftObserverNotificationStatusCmd)
}
