package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusinessCmd = &cobra.Command{
	Use:   "qbusiness",
	Short: "This is the *Amazon Q Business* API Reference.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusinessCmd).Standalone()

	rootCmd.AddCommand(qbusinessCmd)
}
