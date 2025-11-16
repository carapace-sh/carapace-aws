package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetectorCmd = &cobra.Command{
	Use:   "frauddetector",
	Short: "This is the Amazon Fraud Detector API Reference.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetectorCmd).Standalone()

	rootCmd.AddCommand(frauddetectorCmd)
}
