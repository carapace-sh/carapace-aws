package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_describeDestinationsCmd = &cobra.Command{
	Use:   "describe-destinations",
	Short: "Lists all your destinations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_describeDestinationsCmd).Standalone()

	logs_describeDestinationsCmd.Flags().String("destination-name-prefix", "", "The prefix to match.")
	logs_describeDestinationsCmd.Flags().String("limit", "", "The maximum number of items returned.")
	logs_describeDestinationsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	logsCmd.AddCommand(logs_describeDestinationsCmd)
}
