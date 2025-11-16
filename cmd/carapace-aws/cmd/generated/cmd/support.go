package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supportCmd = &cobra.Command{
	Use:   "support",
	Short: "Amazon Web Services Support\n\nThe *Amazon Web Services Support API Reference* is intended for programmers who need detailed information about the Amazon Web Services Support operations and data types.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supportCmd).Standalone()

	rootCmd.AddCommand(supportCmd)
}
