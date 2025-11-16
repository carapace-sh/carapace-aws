package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommitCmd = &cobra.Command{
	Use:   "codecommit",
	Short: "CodeCommit\n\nThis is the *CodeCommit API Reference*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommitCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommitCmd).Standalone()

	})
	rootCmd.AddCommand(codecommitCmd)
}
