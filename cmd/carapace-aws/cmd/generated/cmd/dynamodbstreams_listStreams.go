package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodbstreams_listStreamsCmd = &cobra.Command{
	Use:   "list-streams",
	Short: "Returns an array of stream ARNs associated with the current account and endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodbstreams_listStreamsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodbstreams_listStreamsCmd).Standalone()

		dynamodbstreams_listStreamsCmd.Flags().String("exclusive-start-stream-arn", "", "The ARN (Amazon Resource Name) of the first item that this operation will evaluate.")
		dynamodbstreams_listStreamsCmd.Flags().String("limit", "", "The maximum number of streams to return.")
		dynamodbstreams_listStreamsCmd.Flags().String("table-name", "", "If this parameter is provided, then only the streams associated with this table name are returned.")
	})
	dynamodbstreamsCmd.AddCommand(dynamodbstreams_listStreamsCmd)
}
