package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_deleteDataSourceCmd = &cobra.Command{
	Use:   "delete-data-source",
	Short: "Deletes a data source from a knowledge base.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_deleteDataSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_deleteDataSourceCmd).Standalone()

		bedrockAgent_deleteDataSourceCmd.Flags().String("data-source-id", "", "The unique identifier of the data source to delete.")
		bedrockAgent_deleteDataSourceCmd.Flags().String("knowledge-base-id", "", "The unique identifier of the knowledge base from which to delete the data source.")
		bedrockAgent_deleteDataSourceCmd.MarkFlagRequired("data-source-id")
		bedrockAgent_deleteDataSourceCmd.MarkFlagRequired("knowledge-base-id")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_deleteDataSourceCmd)
}
