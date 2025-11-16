package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iamCmd = &cobra.Command{
	Use:   "iam",
	Short: "Identity and Access Management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iamCmd).Standalone()

	})
	rootCmd.AddCommand(iamCmd)
}
