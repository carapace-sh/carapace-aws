package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_listStreamsCmd = &cobra.Command{
	Use:   "list-streams",
	Short: "Lists your Kinesis data streams.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_listStreamsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesis_listStreamsCmd).Standalone()

		kinesis_listStreamsCmd.Flags().String("exclusive-start-stream-name", "", "The name of the stream to start the list with.")
		kinesis_listStreamsCmd.Flags().String("limit", "", "The maximum number of streams to list.")
		kinesis_listStreamsCmd.Flags().String("next-token", "", "")
	})
	kinesisCmd.AddCommand(kinesis_listStreamsCmd)
}
