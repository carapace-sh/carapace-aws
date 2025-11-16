package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_createRotationOverrideCmd = &cobra.Command{
	Use:   "create-rotation-override",
	Short: "Creates an override for a rotation in an on-call schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_createRotationOverrideCmd).Standalone()

	ssmContacts_createRotationOverrideCmd.Flags().String("end-time", "", "The date and time when the override ends.")
	ssmContacts_createRotationOverrideCmd.Flags().String("idempotency-token", "", "A token that ensures that the operation is called only once with the specified details.")
	ssmContacts_createRotationOverrideCmd.Flags().String("new-contact-ids", "", "The Amazon Resource Names (ARNs) of the contacts to replace those in the current on-call rotation with.")
	ssmContacts_createRotationOverrideCmd.Flags().String("rotation-id", "", "The Amazon Resource Name (ARN) of the rotation to create an override for.")
	ssmContacts_createRotationOverrideCmd.Flags().String("start-time", "", "The date and time when the override goes into effect.")
	ssmContacts_createRotationOverrideCmd.MarkFlagRequired("end-time")
	ssmContacts_createRotationOverrideCmd.MarkFlagRequired("new-contact-ids")
	ssmContacts_createRotationOverrideCmd.MarkFlagRequired("rotation-id")
	ssmContacts_createRotationOverrideCmd.MarkFlagRequired("start-time")
	ssmContactsCmd.AddCommand(ssmContacts_createRotationOverrideCmd)
}
