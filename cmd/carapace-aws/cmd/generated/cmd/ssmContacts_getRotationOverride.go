package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_getRotationOverrideCmd = &cobra.Command{
	Use:   "get-rotation-override",
	Short: "Retrieves information about an override to an on-call rotation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_getRotationOverrideCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmContacts_getRotationOverrideCmd).Standalone()

		ssmContacts_getRotationOverrideCmd.Flags().String("rotation-id", "", "The Amazon Resource Name (ARN) of the overridden rotation to retrieve information about.")
		ssmContacts_getRotationOverrideCmd.Flags().String("rotation-override-id", "", "The Amazon Resource Name (ARN) of the on-call rotation override to retrieve information about.")
		ssmContacts_getRotationOverrideCmd.MarkFlagRequired("rotation-id")
		ssmContacts_getRotationOverrideCmd.MarkFlagRequired("rotation-override-id")
	})
	ssmContactsCmd.AddCommand(ssmContacts_getRotationOverrideCmd)
}
