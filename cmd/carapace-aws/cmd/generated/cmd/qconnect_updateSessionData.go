package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_updateSessionDataCmd = &cobra.Command{
	Use:   "update-session-data",
	Short: "Updates the data stored on an Amazon Q in Connect Session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_updateSessionDataCmd).Standalone()

	qconnect_updateSessionDataCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
	qconnect_updateSessionDataCmd.Flags().String("data", "", "The data stored on the Amazon Q in Connect Session.")
	qconnect_updateSessionDataCmd.Flags().String("namespace", "", "The namespace into which the session data is stored.")
	qconnect_updateSessionDataCmd.Flags().String("session-id", "", "The identifier of the session.")
	qconnect_updateSessionDataCmd.MarkFlagRequired("assistant-id")
	qconnect_updateSessionDataCmd.MarkFlagRequired("data")
	qconnect_updateSessionDataCmd.MarkFlagRequired("session-id")
	qconnectCmd.AddCommand(qconnect_updateSessionDataCmd)
}
