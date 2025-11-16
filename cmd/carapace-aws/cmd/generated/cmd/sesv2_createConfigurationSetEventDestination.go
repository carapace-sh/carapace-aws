package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_createConfigurationSetEventDestinationCmd = &cobra.Command{
	Use:   "create-configuration-set-event-destination",
	Short: "Create an event destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_createConfigurationSetEventDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_createConfigurationSetEventDestinationCmd).Standalone()

		sesv2_createConfigurationSetEventDestinationCmd.Flags().String("configuration-set-name", "", "The name of the configuration set .")
		sesv2_createConfigurationSetEventDestinationCmd.Flags().String("event-destination", "", "An object that defines the event destination.")
		sesv2_createConfigurationSetEventDestinationCmd.Flags().String("event-destination-name", "", "A name that identifies the event destination within the configuration set.")
		sesv2_createConfigurationSetEventDestinationCmd.MarkFlagRequired("configuration-set-name")
		sesv2_createConfigurationSetEventDestinationCmd.MarkFlagRequired("event-destination")
		sesv2_createConfigurationSetEventDestinationCmd.MarkFlagRequired("event-destination-name")
	})
	sesv2Cmd.AddCommand(sesv2_createConfigurationSetEventDestinationCmd)
}
