package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_createRotationCmd = &cobra.Command{
	Use:   "create-rotation",
	Short: "Creates a rotation in an on-call schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_createRotationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmContacts_createRotationCmd).Standalone()

		ssmContacts_createRotationCmd.Flags().String("contact-ids", "", "The Amazon Resource Names (ARNs) of the contacts to add to the rotation.")
		ssmContacts_createRotationCmd.Flags().String("idempotency-token", "", "A token that ensures that the operation is called only once with the specified details.")
		ssmContacts_createRotationCmd.Flags().String("name", "", "The name of the rotation.")
		ssmContacts_createRotationCmd.Flags().String("recurrence", "", "Information about the rule that specifies when a shift's team members rotate.")
		ssmContacts_createRotationCmd.Flags().String("start-time", "", "The date and time that the rotation goes into effect.")
		ssmContacts_createRotationCmd.Flags().String("tags", "", "Optional metadata to assign to the rotation.")
		ssmContacts_createRotationCmd.Flags().String("time-zone-id", "", "The time zone to base the rotation’s activity on in Internet Assigned Numbers Authority (IANA) format.")
		ssmContacts_createRotationCmd.MarkFlagRequired("contact-ids")
		ssmContacts_createRotationCmd.MarkFlagRequired("name")
		ssmContacts_createRotationCmd.MarkFlagRequired("recurrence")
		ssmContacts_createRotationCmd.MarkFlagRequired("time-zone-id")
	})
	ssmContactsCmd.AddCommand(ssmContacts_createRotationCmd)
}
