package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_deleteConnectionCmd = &cobra.Command{
	Use:   "delete-connection",
	Short: "Deletes a connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_deleteConnectionCmd).Standalone()

	events_deleteConnectionCmd.Flags().String("name", "", "The name of the connection to delete.")
	events_deleteConnectionCmd.MarkFlagRequired("name")
	eventsCmd.AddCommand(events_deleteConnectionCmd)
}
