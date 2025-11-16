package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pollyCmd = &cobra.Command{
	Use:   "polly",
	Short: "Amazon Polly is a web service that makes it easy to synthesize speech from text.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pollyCmd).Standalone()

	rootCmd.AddCommand(pollyCmd)
}
