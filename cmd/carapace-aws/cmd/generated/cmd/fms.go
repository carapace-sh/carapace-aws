package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fmsCmd = &cobra.Command{
	Use:   "fms",
	Short: "This is the *Firewall Manager API Reference*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fmsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fmsCmd).Standalone()

	})
	rootCmd.AddCommand(fmsCmd)
}
