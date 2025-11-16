package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_createConfigurationSetEventDestinationCmd = &cobra.Command{
	Use:   "create-configuration-set-event-destination",
	Short: "Creates a configuration set event destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_createConfigurationSetEventDestinationCmd).Standalone()

	ses_createConfigurationSetEventDestinationCmd.Flags().String("configuration-set-name", "", "The name of the configuration set that the event destination should be associated with.")
	ses_createConfigurationSetEventDestinationCmd.Flags().String("event-destination", "", "An object that describes the Amazon Web Services service that email sending event where information is published.")
	ses_createConfigurationSetEventDestinationCmd.MarkFlagRequired("configuration-set-name")
	ses_createConfigurationSetEventDestinationCmd.MarkFlagRequired("event-destination")
	sesCmd.AddCommand(ses_createConfigurationSetEventDestinationCmd)
}
