package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var trustedadvisorCmd = &cobra.Command{
	Use:   "trustedadvisor",
	Short: "TrustedAdvisor Public API",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trustedadvisorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(trustedadvisorCmd).Standalone()

	})
	rootCmd.AddCommand(trustedadvisorCmd)
}
