package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supportCmd = &cobra.Command{
	Use:   "support",
	Short: "Amazon Web Services Support",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(supportCmd).Standalone()

	})
	rootCmd.AddCommand(supportCmd)
}
