package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_updateConfigurationSetEventDestinationCmd = &cobra.Command{
	Use:   "update-configuration-set-event-destination",
	Short: "Updates the event destination of a configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_updateConfigurationSetEventDestinationCmd).Standalone()

	ses_updateConfigurationSetEventDestinationCmd.Flags().String("configuration-set-name", "", "The name of the configuration set that contains the event destination.")
	ses_updateConfigurationSetEventDestinationCmd.Flags().String("event-destination", "", "The event destination object.")
	ses_updateConfigurationSetEventDestinationCmd.MarkFlagRequired("configuration-set-name")
	ses_updateConfigurationSetEventDestinationCmd.MarkFlagRequired("event-destination")
	sesCmd.AddCommand(ses_updateConfigurationSetEventDestinationCmd)
}
