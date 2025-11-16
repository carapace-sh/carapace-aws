package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmQuicksetupCmd = &cobra.Command{
	Use:   "ssm-quicksetup",
	Short: "Quick Setup helps you quickly configure frequently used services and features with recommended best practices.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmQuicksetupCmd).Standalone()

	rootCmd.AddCommand(ssmQuicksetupCmd)
}
