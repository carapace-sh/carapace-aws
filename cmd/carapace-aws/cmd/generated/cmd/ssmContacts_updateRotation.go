package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_updateRotationCmd = &cobra.Command{
	Use:   "update-rotation",
	Short: "Updates the information specified for an on-call rotation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_updateRotationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmContacts_updateRotationCmd).Standalone()

		ssmContacts_updateRotationCmd.Flags().String("contact-ids", "", "The Amazon Resource Names (ARNs) of the contacts to include in the updated rotation.")
		ssmContacts_updateRotationCmd.Flags().String("recurrence", "", "Information about how long the updated rotation lasts before restarting at the beginning of the shift order.")
		ssmContacts_updateRotationCmd.Flags().String("rotation-id", "", "The Amazon Resource Name (ARN) of the rotation to update.")
		ssmContacts_updateRotationCmd.Flags().String("start-time", "", "The date and time the rotation goes into effect.")
		ssmContacts_updateRotationCmd.Flags().String("time-zone-id", "", "The time zone to base the updated rotation’s activity on, in Internet Assigned Numbers Authority (IANA) format.")
		ssmContacts_updateRotationCmd.MarkFlagRequired("recurrence")
		ssmContacts_updateRotationCmd.MarkFlagRequired("rotation-id")
	})
	ssmContactsCmd.AddCommand(ssmContacts_updateRotationCmd)
}
