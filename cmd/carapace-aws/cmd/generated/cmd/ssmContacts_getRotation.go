package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_getRotationCmd = &cobra.Command{
	Use:   "get-rotation",
	Short: "Retrieves information about an on-call rotation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_getRotationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmContacts_getRotationCmd).Standalone()

		ssmContacts_getRotationCmd.Flags().String("rotation-id", "", "The Amazon Resource Name (ARN) of the on-call rotation to retrieve information about.")
		ssmContacts_getRotationCmd.MarkFlagRequired("rotation-id")
	})
	ssmContactsCmd.AddCommand(ssmContacts_getRotationCmd)
}
