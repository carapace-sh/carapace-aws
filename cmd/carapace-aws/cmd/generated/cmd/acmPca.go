package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acmPcaCmd = &cobra.Command{
	Use:   "acm-pca",
	Short: "This is the *Amazon Web Services Private Certificate Authority API Reference*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acmPcaCmd).Standalone()

	rootCmd.AddCommand(acmPcaCmd)
}
