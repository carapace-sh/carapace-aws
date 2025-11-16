package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var firehose_listDeliveryStreamsCmd = &cobra.Command{
	Use:   "list-delivery-streams",
	Short: "Lists your Firehose streams in alphabetical order of their names.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firehose_listDeliveryStreamsCmd).Standalone()

	firehose_listDeliveryStreamsCmd.Flags().String("delivery-stream-type", "", "The Firehose stream type.")
	firehose_listDeliveryStreamsCmd.Flags().String("exclusive-start-delivery-stream-name", "", "The list of Firehose streams returned by this call to `ListDeliveryStreams` will start with the Firehose stream whose name comes alphabetically immediately after the name you specify in `ExclusiveStartDeliveryStreamName`.")
	firehose_listDeliveryStreamsCmd.Flags().String("limit", "", "The maximum number of Firehose streams to list.")
	firehoseCmd.AddCommand(firehose_listDeliveryStreamsCmd)
}
