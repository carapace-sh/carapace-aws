package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_listEventSourceMappingsCmd = &cobra.Command{
	Use:   "list-event-source-mappings",
	Short: "Lists event source mappings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_listEventSourceMappingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambda_listEventSourceMappingsCmd).Standalone()

		lambda_listEventSourceMappingsCmd.Flags().String("event-source-arn", "", "The Amazon Resource Name (ARN) of the event source.")
		lambda_listEventSourceMappingsCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
		lambda_listEventSourceMappingsCmd.Flags().String("marker", "", "A pagination token returned by a previous call.")
		lambda_listEventSourceMappingsCmd.Flags().String("max-items", "", "The maximum number of event source mappings to return.")
	})
	lambdaCmd.AddCommand(lambda_listEventSourceMappingsCmd)
}
