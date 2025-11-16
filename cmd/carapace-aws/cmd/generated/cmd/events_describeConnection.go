package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_describeConnectionCmd = &cobra.Command{
	Use:   "describe-connection",
	Short: "Retrieves details about a connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_describeConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(events_describeConnectionCmd).Standalone()

		events_describeConnectionCmd.Flags().String("name", "", "The name of the connection to retrieve.")
		events_describeConnectionCmd.MarkFlagRequired("name")
	})
	eventsCmd.AddCommand(events_describeConnectionCmd)
}
