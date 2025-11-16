package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdomCmd = &cobra.Command{
	Use:   "wisdom",
	Short: "Amazon Connect Wisdom delivers agents the information they need to solve customer issues as they're actively speaking with customers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdomCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wisdomCmd).Standalone()

	})
	rootCmd.AddCommand(wisdomCmd)
}
