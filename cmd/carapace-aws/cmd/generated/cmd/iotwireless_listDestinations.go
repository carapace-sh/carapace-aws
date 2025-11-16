package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_listDestinationsCmd = &cobra.Command{
	Use:   "list-destinations",
	Short: "Lists the destinations registered to your AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_listDestinationsCmd).Standalone()

	iotwireless_listDestinationsCmd.Flags().String("max-results", "", "The maximum number of results to return in this operation.")
	iotwireless_listDestinationsCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
	iotwirelessCmd.AddCommand(iotwireless_listDestinationsCmd)
}
