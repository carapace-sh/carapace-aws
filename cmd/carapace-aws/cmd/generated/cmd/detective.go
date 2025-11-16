package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detectiveCmd = &cobra.Command{
	Use:   "detective",
	Short: "Detective uses machine learning and purpose-built visualizations to help you to analyze and investigate security issues across your Amazon Web Services (Amazon Web Services) workloads.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detectiveCmd).Standalone()

	rootCmd.AddCommand(detectiveCmd)
}
