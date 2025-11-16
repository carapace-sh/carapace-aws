package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_deleteDestinationCmd = &cobra.Command{
	Use:   "delete-destination",
	Short: "Deletes the specified destination, and eventually disables all the subscription filters that publish to it.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_deleteDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_deleteDestinationCmd).Standalone()

		logs_deleteDestinationCmd.Flags().String("destination-name", "", "The name of the destination.")
		logs_deleteDestinationCmd.MarkFlagRequired("destination-name")
	})
	logsCmd.AddCommand(logs_deleteDestinationCmd)
}
