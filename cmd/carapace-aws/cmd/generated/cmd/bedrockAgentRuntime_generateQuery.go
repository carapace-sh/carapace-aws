package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_generateQueryCmd = &cobra.Command{
	Use:   "generate-query",
	Short: "Generates an SQL query from a natural language query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_generateQueryCmd).Standalone()

	bedrockAgentRuntime_generateQueryCmd.Flags().String("query-generation-input", "", "Specifies information about a natural language query to transform into SQL.")
	bedrockAgentRuntime_generateQueryCmd.Flags().String("transformation-configuration", "", "Specifies configurations for transforming the natural language query into SQL.")
	bedrockAgentRuntime_generateQueryCmd.MarkFlagRequired("query-generation-input")
	bedrockAgentRuntime_generateQueryCmd.MarkFlagRequired("transformation-configuration")
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_generateQueryCmd)
}
