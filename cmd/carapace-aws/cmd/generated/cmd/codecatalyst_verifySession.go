package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_verifySessionCmd = &cobra.Command{
	Use:   "verify-session",
	Short: "Verifies whether the calling user has a valid Amazon CodeCatalyst login and session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_verifySessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecatalyst_verifySessionCmd).Standalone()

	})
	codecatalystCmd.AddCommand(codecatalyst_verifySessionCmd)
}
