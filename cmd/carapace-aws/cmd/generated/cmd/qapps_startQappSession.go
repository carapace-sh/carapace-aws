package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_startQappSessionCmd = &cobra.Command{
	Use:   "start-qapp-session",
	Short: "Starts a new session for an Amazon Q App, allowing inputs to be provided and the app to be run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_startQappSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qapps_startQappSessionCmd).Standalone()

		qapps_startQappSessionCmd.Flags().String("app-id", "", "The unique identifier of the Q App to start a session for.")
		qapps_startQappSessionCmd.Flags().String("app-version", "", "The version of the Q App to use for the session.")
		qapps_startQappSessionCmd.Flags().String("initial-values", "", "Optional initial input values to provide for the Q App session.")
		qapps_startQappSessionCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
		qapps_startQappSessionCmd.Flags().String("session-id", "", "The unique identifier of the a Q App session.")
		qapps_startQappSessionCmd.Flags().String("tags", "", "Optional tags to associate with the new Q App session.")
		qapps_startQappSessionCmd.MarkFlagRequired("app-id")
		qapps_startQappSessionCmd.MarkFlagRequired("app-version")
		qapps_startQappSessionCmd.MarkFlagRequired("instance-id")
	})
	qappsCmd.AddCommand(qapps_startQappSessionCmd)
}
