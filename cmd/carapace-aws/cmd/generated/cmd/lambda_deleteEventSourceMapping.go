package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_deleteEventSourceMappingCmd = &cobra.Command{
	Use:   "delete-event-source-mapping",
	Short: "Deletes an [event source mapping](https://docs.aws.amazon.com/lambda/latest/dg/intro-invocation-modes.html). You can get the identifier of a mapping from the output of [ListEventSourceMappings]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_deleteEventSourceMappingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambda_deleteEventSourceMappingCmd).Standalone()

		lambda_deleteEventSourceMappingCmd.Flags().String("uuid", "", "The identifier of the event source mapping.")
		lambda_deleteEventSourceMappingCmd.MarkFlagRequired("uuid")
	})
	lambdaCmd.AddCommand(lambda_deleteEventSourceMappingCmd)
}
