package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_getQappSessionCmd = &cobra.Command{
	Use:   "get-qapp-session",
	Short: "Retrieves the current state and results for an active session of an Amazon Q App.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_getQappSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qapps_getQappSessionCmd).Standalone()

		qapps_getQappSessionCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
		qapps_getQappSessionCmd.Flags().String("session-id", "", "The unique identifier of the Q App session to retrieve.")
		qapps_getQappSessionCmd.MarkFlagRequired("instance-id")
		qapps_getQappSessionCmd.MarkFlagRequired("session-id")
	})
	qappsCmd.AddCommand(qapps_getQappSessionCmd)
}
