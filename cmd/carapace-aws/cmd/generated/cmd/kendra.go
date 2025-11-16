package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendraCmd = &cobra.Command{
	Use:   "kendra",
	Short: "Amazon Kendra is a service for indexing large document sets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendraCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendraCmd).Standalone()

	})
	rootCmd.AddCommand(kendraCmd)
}
