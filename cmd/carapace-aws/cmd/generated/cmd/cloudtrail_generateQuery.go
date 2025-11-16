package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_generateQueryCmd = &cobra.Command{
	Use:   "generate-query",
	Short: "Generates a query from a natural language prompt.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_generateQueryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudtrail_generateQueryCmd).Standalone()

		cloudtrail_generateQueryCmd.Flags().String("event-data-stores", "", "The ARN (or ID suffix of the ARN) of the event data store that you want to query.")
		cloudtrail_generateQueryCmd.Flags().String("prompt", "", "The prompt that you want to use to generate the query.")
		cloudtrail_generateQueryCmd.MarkFlagRequired("event-data-stores")
		cloudtrail_generateQueryCmd.MarkFlagRequired("prompt")
	})
	cloudtrailCmd.AddCommand(cloudtrail_generateQueryCmd)
}
