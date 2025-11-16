package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acmCmd = &cobra.Command{
	Use:   "acm",
	Short: "Certificate Manager",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acmCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(acmCmd).Standalone()

	})
	rootCmd.AddCommand(acmCmd)
}
