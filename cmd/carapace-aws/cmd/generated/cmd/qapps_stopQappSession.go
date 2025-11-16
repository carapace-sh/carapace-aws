package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_stopQappSessionCmd = &cobra.Command{
	Use:   "stop-qapp-session",
	Short: "Stops an active session for an Amazon Q App.This deletes all data related to the session and makes it invalid for future uses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_stopQappSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qapps_stopQappSessionCmd).Standalone()

		qapps_stopQappSessionCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
		qapps_stopQappSessionCmd.Flags().String("session-id", "", "The unique identifier of the Q App session to stop.")
		qapps_stopQappSessionCmd.MarkFlagRequired("instance-id")
		qapps_stopQappSessionCmd.MarkFlagRequired("session-id")
	})
	qappsCmd.AddCommand(qapps_stopQappSessionCmd)
}
