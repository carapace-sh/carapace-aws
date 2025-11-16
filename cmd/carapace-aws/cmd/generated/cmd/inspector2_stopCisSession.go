package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_stopCisSessionCmd = &cobra.Command{
	Use:   "stop-cis-session",
	Short: "Stops a CIS session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_stopCisSessionCmd).Standalone()

	inspector2_stopCisSessionCmd.Flags().String("message", "", "The stop CIS session message.")
	inspector2_stopCisSessionCmd.Flags().String("scan-job-id", "", "A unique identifier for the scan job.")
	inspector2_stopCisSessionCmd.Flags().String("session-token", "", "The unique token that identifies the CIS session.")
	inspector2_stopCisSessionCmd.MarkFlagRequired("message")
	inspector2_stopCisSessionCmd.MarkFlagRequired("scan-job-id")
	inspector2_stopCisSessionCmd.MarkFlagRequired("session-token")
	inspector2Cmd.AddCommand(inspector2_stopCisSessionCmd)
}
