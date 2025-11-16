package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_getConfigurationSetEventDestinationsCmd = &cobra.Command{
	Use:   "get-configuration-set-event-destinations",
	Short: "Retrieve a list of event destinations that are associated with a configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_getConfigurationSetEventDestinationsCmd).Standalone()

	sesv2_getConfigurationSetEventDestinationsCmd.Flags().String("configuration-set-name", "", "The name of the configuration set that contains the event destination.")
	sesv2_getConfigurationSetEventDestinationsCmd.MarkFlagRequired("configuration-set-name")
	sesv2Cmd.AddCommand(sesv2_getConfigurationSetEventDestinationsCmd)
}
