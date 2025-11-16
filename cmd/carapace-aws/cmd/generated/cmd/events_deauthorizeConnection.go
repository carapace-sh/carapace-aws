package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_deauthorizeConnectionCmd = &cobra.Command{
	Use:   "deauthorize-connection",
	Short: "Removes all authorization parameters from the connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_deauthorizeConnectionCmd).Standalone()

	events_deauthorizeConnectionCmd.Flags().String("name", "", "The name of the connection to remove authorization from.")
	events_deauthorizeConnectionCmd.MarkFlagRequired("name")
	eventsCmd.AddCommand(events_deauthorizeConnectionCmd)
}
