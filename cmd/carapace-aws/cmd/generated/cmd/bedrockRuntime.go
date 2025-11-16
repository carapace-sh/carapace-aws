package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockRuntimeCmd = &cobra.Command{
	Use:   "bedrock-runtime",
	Short: "Describes the API operations for running inference using Amazon Bedrock models.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockRuntimeCmd).Standalone()

	rootCmd.AddCommand(bedrockRuntimeCmd)
}
