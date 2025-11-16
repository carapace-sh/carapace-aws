package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmGuiconnect_deleteConnectionRecordingPreferencesCmd = &cobra.Command{
	Use:   "delete-connection-recording-preferences",
	Short: "Deletes the preferences for recording RDP connections.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmGuiconnect_deleteConnectionRecordingPreferencesCmd).Standalone()

	ssmGuiconnect_deleteConnectionRecordingPreferencesCmd.Flags().String("client-token", "", "User-provided idempotency token.")
	ssmGuiconnectCmd.AddCommand(ssmGuiconnect_deleteConnectionRecordingPreferencesCmd)
}
