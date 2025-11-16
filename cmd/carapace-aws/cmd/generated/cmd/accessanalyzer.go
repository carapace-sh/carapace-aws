package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzerCmd = &cobra.Command{
	Use:   "accessanalyzer",
	Short: "Identity and Access Management Access Analyzer helps you to set, verify, and refine your IAM policies by providing a suite of capabilities.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzerCmd).Standalone()

	rootCmd.AddCommand(accessanalyzerCmd)
}
