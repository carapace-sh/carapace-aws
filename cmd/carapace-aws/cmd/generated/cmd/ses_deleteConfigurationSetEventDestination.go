package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_deleteConfigurationSetEventDestinationCmd = &cobra.Command{
	Use:   "delete-configuration-set-event-destination",
	Short: "Deletes a configuration set event destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_deleteConfigurationSetEventDestinationCmd).Standalone()

	ses_deleteConfigurationSetEventDestinationCmd.Flags().String("configuration-set-name", "", "The name of the configuration set from which to delete the event destination.")
	ses_deleteConfigurationSetEventDestinationCmd.Flags().String("event-destination-name", "", "The name of the event destination to delete.")
	ses_deleteConfigurationSetEventDestinationCmd.MarkFlagRequired("configuration-set-name")
	ses_deleteConfigurationSetEventDestinationCmd.MarkFlagRequired("event-destination-name")
	sesCmd.AddCommand(ses_deleteConfigurationSetEventDestinationCmd)
}
