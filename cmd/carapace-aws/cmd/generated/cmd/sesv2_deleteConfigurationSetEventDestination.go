package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_deleteConfigurationSetEventDestinationCmd = &cobra.Command{
	Use:   "delete-configuration-set-event-destination",
	Short: "Delete an event destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_deleteConfigurationSetEventDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_deleteConfigurationSetEventDestinationCmd).Standalone()

		sesv2_deleteConfigurationSetEventDestinationCmd.Flags().String("configuration-set-name", "", "The name of the configuration set that contains the event destination to delete.")
		sesv2_deleteConfigurationSetEventDestinationCmd.Flags().String("event-destination-name", "", "The name of the event destination to delete.")
		sesv2_deleteConfigurationSetEventDestinationCmd.MarkFlagRequired("configuration-set-name")
		sesv2_deleteConfigurationSetEventDestinationCmd.MarkFlagRequired("event-destination-name")
	})
	sesv2Cmd.AddCommand(sesv2_deleteConfigurationSetEventDestinationCmd)
}
