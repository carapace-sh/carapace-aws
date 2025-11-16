package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarConnectionsCmd = &cobra.Command{
	Use:   "codestar-connections",
	Short: "AWS CodeStar Connections",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarConnectionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codestarConnectionsCmd).Standalone()

	})
	rootCmd.AddCommand(codestarConnectionsCmd)
}
