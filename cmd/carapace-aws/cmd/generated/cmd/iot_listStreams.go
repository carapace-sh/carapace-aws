package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listStreamsCmd = &cobra.Command{
	Use:   "list-streams",
	Short: "Lists all of the streams in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listStreamsCmd).Standalone()

	iot_listStreamsCmd.Flags().String("ascending-order", "", "Set to true to return the list of streams in ascending order.")
	iot_listStreamsCmd.Flags().String("max-results", "", "The maximum number of results to return at a time.")
	iot_listStreamsCmd.Flags().String("next-token", "", "A token used to get the next set of results.")
	iotCmd.AddCommand(iot_listStreamsCmd)
}
