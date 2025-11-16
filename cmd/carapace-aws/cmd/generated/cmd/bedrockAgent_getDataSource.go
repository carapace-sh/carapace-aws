package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_getDataSourceCmd = &cobra.Command{
	Use:   "get-data-source",
	Short: "Gets information about a data source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_getDataSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_getDataSourceCmd).Standalone()

		bedrockAgent_getDataSourceCmd.Flags().String("data-source-id", "", "The unique identifier of the data source.")
		bedrockAgent_getDataSourceCmd.Flags().String("knowledge-base-id", "", "The unique identifier of the knowledge base for the data source.")
		bedrockAgent_getDataSourceCmd.MarkFlagRequired("data-source-id")
		bedrockAgent_getDataSourceCmd.MarkFlagRequired("knowledge-base-id")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_getDataSourceCmd)
}
