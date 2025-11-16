package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_deleteRotationCmd = &cobra.Command{
	Use:   "delete-rotation",
	Short: "Deletes a rotation from the system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_deleteRotationCmd).Standalone()

	ssmContacts_deleteRotationCmd.Flags().String("rotation-id", "", "The Amazon Resource Name (ARN) of the on-call rotation to delete.")
	ssmContacts_deleteRotationCmd.MarkFlagRequired("rotation-id")
	ssmContactsCmd.AddCommand(ssmContacts_deleteRotationCmd)
}
