package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_startCisSessionCmd = &cobra.Command{
	Use:   "start-cis-session",
	Short: "Starts a CIS session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_startCisSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_startCisSessionCmd).Standalone()

		inspector2_startCisSessionCmd.Flags().String("message", "", "The start CIS session message.")
		inspector2_startCisSessionCmd.Flags().String("scan-job-id", "", "A unique identifier for the scan job.")
		inspector2_startCisSessionCmd.MarkFlagRequired("message")
		inspector2_startCisSessionCmd.MarkFlagRequired("scan-job-id")
	})
	inspector2Cmd.AddCommand(inspector2_startCisSessionCmd)
}
