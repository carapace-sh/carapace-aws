package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_updateQappSessionCmd = &cobra.Command{
	Use:   "update-qapp-session",
	Short: "Updates the session for a given Q App `sessionId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_updateQappSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qapps_updateQappSessionCmd).Standalone()

		qapps_updateQappSessionCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
		qapps_updateQappSessionCmd.Flags().String("session-id", "", "The unique identifier of the Q App session to provide input for.")
		qapps_updateQappSessionCmd.Flags().String("values", "", "The input values to provide for the current state of the Q App session.")
		qapps_updateQappSessionCmd.MarkFlagRequired("instance-id")
		qapps_updateQappSessionCmd.MarkFlagRequired("session-id")
	})
	qappsCmd.AddCommand(qapps_updateQappSessionCmd)
}
