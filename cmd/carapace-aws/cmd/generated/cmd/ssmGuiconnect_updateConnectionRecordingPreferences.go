package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmGuiconnect_updateConnectionRecordingPreferencesCmd = &cobra.Command{
	Use:   "update-connection-recording-preferences",
	Short: "Updates the preferences for recording RDP connections.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmGuiconnect_updateConnectionRecordingPreferencesCmd).Standalone()

	ssmGuiconnect_updateConnectionRecordingPreferencesCmd.Flags().String("client-token", "", "User-provided idempotency token.")
	ssmGuiconnect_updateConnectionRecordingPreferencesCmd.Flags().String("connection-recording-preferences", "", "The set of preferences used for recording RDP connections in the requesting Amazon Web Services account and Amazon Web Services Region.")
	ssmGuiconnect_updateConnectionRecordingPreferencesCmd.MarkFlagRequired("connection-recording-preferences")
	ssmGuiconnectCmd.AddCommand(ssmGuiconnect_updateConnectionRecordingPreferencesCmd)
}
