package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qappsCmd = &cobra.Command{
	Use:   "qapps",
	Short: "The Amazon Q Apps feature capability within Amazon Q Business allows web experience users to create lightweight, purpose-built AI apps to fulfill specific tasks from within their web experience.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qappsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qappsCmd).Standalone()

	})
	rootCmd.AddCommand(qappsCmd)
}
