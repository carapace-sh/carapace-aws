package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmGuiconnect_getConnectionRecordingPreferencesCmd = &cobra.Command{
	Use:   "get-connection-recording-preferences",
	Short: "Returns the preferences specified for recording RDP connections in the requesting Amazon Web Services account and Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmGuiconnect_getConnectionRecordingPreferencesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmGuiconnect_getConnectionRecordingPreferencesCmd).Standalone()

	})
	ssmGuiconnectCmd.AddCommand(ssmGuiconnect_getConnectionRecordingPreferencesCmd)
}
