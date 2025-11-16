package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcZonalShift_updateAutoshiftObserverNotificationStatusCmd = &cobra.Command{
	Use:   "update-autoshift-observer-notification-status",
	Short: "Update the status of autoshift observer notification.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcZonalShift_updateAutoshiftObserverNotificationStatusCmd).Standalone()

	arcZonalShift_updateAutoshiftObserverNotificationStatusCmd.Flags().String("status", "", "The status to set for autoshift observer notification.")
	arcZonalShift_updateAutoshiftObserverNotificationStatusCmd.MarkFlagRequired("status")
	arcZonalShiftCmd.AddCommand(arcZonalShift_updateAutoshiftObserverNotificationStatusCmd)
}
