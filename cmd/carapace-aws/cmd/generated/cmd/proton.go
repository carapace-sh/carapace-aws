package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var protonCmd = &cobra.Command{
	Use:   "proton",
	Short: "This is the Proton Service API Reference.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(protonCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(protonCmd).Standalone()

	})
	rootCmd.AddCommand(protonCmd)
}
