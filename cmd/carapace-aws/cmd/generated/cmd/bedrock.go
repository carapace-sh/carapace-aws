package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockCmd = &cobra.Command{
	Use:   "bedrock",
	Short: "Describes the API operations for creating, managing, fine-turning, and evaluating Amazon Bedrock models.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockCmd).Standalone()

	})
	rootCmd.AddCommand(bedrockCmd)
}
