package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aiopsCmd = &cobra.Command{
	Use:   "aiops",
	Short: "The CloudWatch investigations feature is a generative AI-powered assistant that can help you respond to incidents in your system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aiopsCmd).Standalone()

	rootCmd.AddCommand(aiopsCmd)
}
