package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2biCmd = &cobra.Command{
	Use:   "b2bi",
	Short: "This is the *Amazon Web Services B2B Data Interchange API Reference*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2biCmd).Standalone()

	rootCmd.AddCommand(b2biCmd)
}
