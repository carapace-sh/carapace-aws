package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_getEventTypesCmd = &cobra.Command{
	Use:   "get-event-types",
	Short: "Gets all event types or a specific event type if name is provided.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_getEventTypesCmd).Standalone()

	frauddetector_getEventTypesCmd.Flags().String("max-results", "", "The maximum number of objects to return for the request.")
	frauddetector_getEventTypesCmd.Flags().String("name", "", "The name.")
	frauddetector_getEventTypesCmd.Flags().String("next-token", "", "The next token for the subsequent request.")
	frauddetectorCmd.AddCommand(frauddetector_getEventTypesCmd)
}
