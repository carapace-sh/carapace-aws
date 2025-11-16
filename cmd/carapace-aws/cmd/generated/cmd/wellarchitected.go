package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitectedCmd = &cobra.Command{
	Use:   "wellarchitected",
	Short: "Well-Architected Tool\n\nThis is the *Well-Architected Tool API Reference*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitectedCmd).Standalone()

	rootCmd.AddCommand(wellarchitectedCmd)
}
