package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_sendCisSessionHealthCmd = &cobra.Command{
	Use:   "send-cis-session-health",
	Short: "Sends a CIS session health.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_sendCisSessionHealthCmd).Standalone()

	inspector2_sendCisSessionHealthCmd.Flags().String("scan-job-id", "", "A unique identifier for the scan job.")
	inspector2_sendCisSessionHealthCmd.Flags().String("session-token", "", "The unique token that identifies the CIS session.")
	inspector2_sendCisSessionHealthCmd.MarkFlagRequired("scan-job-id")
	inspector2_sendCisSessionHealthCmd.MarkFlagRequired("session-token")
	inspector2Cmd.AddCommand(inspector2_sendCisSessionHealthCmd)
}
