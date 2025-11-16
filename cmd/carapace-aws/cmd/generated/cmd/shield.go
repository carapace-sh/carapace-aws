package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shieldCmd = &cobra.Command{
	Use:   "shield",
	Short: "Shield Advanced\n\nThis is the *Shield Advanced API Reference*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shieldCmd).Standalone()

	rootCmd.AddCommand(shieldCmd)
}
