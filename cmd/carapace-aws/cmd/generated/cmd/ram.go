package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ramCmd = &cobra.Command{
	Use:   "ram",
	Short: "This is the *Resource Access Manager API Reference*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ramCmd).Standalone()

	rootCmd.AddCommand(ramCmd)
}
