package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectoryCmd = &cobra.Command{
	Use:   "clouddirectory",
	Short: "Amazon Cloud Directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectoryCmd).Standalone()

	})
	rootCmd.AddCommand(clouddirectoryCmd)
}
