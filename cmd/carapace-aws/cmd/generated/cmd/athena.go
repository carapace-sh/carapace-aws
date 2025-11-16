package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athenaCmd = &cobra.Command{
	Use:   "athena",
	Short: "Amazon Athena is an interactive query service that lets you use standard SQL to analyze data directly in Amazon S3.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athenaCmd).Standalone()

	rootCmd.AddCommand(athenaCmd)
}
