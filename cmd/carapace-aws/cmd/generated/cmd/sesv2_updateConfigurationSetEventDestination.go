package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_updateConfigurationSetEventDestinationCmd = &cobra.Command{
	Use:   "update-configuration-set-event-destination",
	Short: "Update the configuration of an event destination for a configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_updateConfigurationSetEventDestinationCmd).Standalone()

	sesv2_updateConfigurationSetEventDestinationCmd.Flags().String("configuration-set-name", "", "The name of the configuration set that contains the event destination to modify.")
	sesv2_updateConfigurationSetEventDestinationCmd.Flags().String("event-destination", "", "An object that defines the event destination.")
	sesv2_updateConfigurationSetEventDestinationCmd.Flags().String("event-destination-name", "", "The name of the event destination.")
	sesv2_updateConfigurationSetEventDestinationCmd.MarkFlagRequired("configuration-set-name")
	sesv2_updateConfigurationSetEventDestinationCmd.MarkFlagRequired("event-destination")
	sesv2_updateConfigurationSetEventDestinationCmd.MarkFlagRequired("event-destination-name")
	sesv2Cmd.AddCommand(sesv2_updateConfigurationSetEventDestinationCmd)
}
