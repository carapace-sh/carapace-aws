package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_listRotationOverridesCmd = &cobra.Command{
	Use:   "list-rotation-overrides",
	Short: "Retrieves a list of overrides currently specified for an on-call rotation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_listRotationOverridesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmContacts_listRotationOverridesCmd).Standalone()

		ssmContacts_listRotationOverridesCmd.Flags().String("end-time", "", "The date and time for the end of a time range for listing overrides.")
		ssmContacts_listRotationOverridesCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
		ssmContacts_listRotationOverridesCmd.Flags().String("next-token", "", "A token to start the list.")
		ssmContacts_listRotationOverridesCmd.Flags().String("rotation-id", "", "The Amazon Resource Name (ARN) of the rotation to retrieve information about.")
		ssmContacts_listRotationOverridesCmd.Flags().String("start-time", "", "The date and time for the beginning of a time range for listing overrides.")
		ssmContacts_listRotationOverridesCmd.MarkFlagRequired("end-time")
		ssmContacts_listRotationOverridesCmd.MarkFlagRequired("rotation-id")
		ssmContacts_listRotationOverridesCmd.MarkFlagRequired("start-time")
	})
	ssmContactsCmd.AddCommand(ssmContacts_listRotationOverridesCmd)
}
