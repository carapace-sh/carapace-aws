package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitectedCmd = &cobra.Command{
	Use:   "wellarchitected",
	Short: "Well-Architected Tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitectedCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitectedCmd).Standalone()

	})
	rootCmd.AddCommand(wellarchitectedCmd)
}
