package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_deleteRotationOverrideCmd = &cobra.Command{
	Use:   "delete-rotation-override",
	Short: "Deletes an existing override for an on-call rotation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_deleteRotationOverrideCmd).Standalone()

	ssmContacts_deleteRotationOverrideCmd.Flags().String("rotation-id", "", "The Amazon Resource Name (ARN) of the rotation that was overridden.")
	ssmContacts_deleteRotationOverrideCmd.Flags().String("rotation-override-id", "", "The Amazon Resource Name (ARN) of the on-call rotation override to delete.")
	ssmContacts_deleteRotationOverrideCmd.MarkFlagRequired("rotation-id")
	ssmContacts_deleteRotationOverrideCmd.MarkFlagRequired("rotation-override-id")
	ssmContactsCmd.AddCommand(ssmContacts_deleteRotationOverrideCmd)
}
