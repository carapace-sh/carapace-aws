package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_listPreviewRotationShiftsCmd = &cobra.Command{
	Use:   "list-preview-rotation-shifts",
	Short: "Returns a list of shifts based on rotation configuration parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_listPreviewRotationShiftsCmd).Standalone()

	ssmContacts_listPreviewRotationShiftsCmd.Flags().String("end-time", "", "The date and time a rotation shift would end.")
	ssmContacts_listPreviewRotationShiftsCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	ssmContacts_listPreviewRotationShiftsCmd.Flags().String("members", "", "The contacts that would be assigned to a rotation.")
	ssmContacts_listPreviewRotationShiftsCmd.Flags().String("next-token", "", "A token to start the list.")
	ssmContacts_listPreviewRotationShiftsCmd.Flags().String("overrides", "", "Information about changes that would be made in a rotation override.")
	ssmContacts_listPreviewRotationShiftsCmd.Flags().String("recurrence", "", "Information about how long a rotation would last before restarting at the beginning of the shift order.")
	ssmContacts_listPreviewRotationShiftsCmd.Flags().String("rotation-start-time", "", "The date and time a rotation would begin.")
	ssmContacts_listPreviewRotationShiftsCmd.Flags().String("start-time", "", "Used to filter the range of calculated shifts before sending the response back to the user.")
	ssmContacts_listPreviewRotationShiftsCmd.Flags().String("time-zone-id", "", "The time zone the rotation’s activity would be based on, in Internet Assigned Numbers Authority (IANA) format.")
	ssmContacts_listPreviewRotationShiftsCmd.MarkFlagRequired("end-time")
	ssmContacts_listPreviewRotationShiftsCmd.MarkFlagRequired("members")
	ssmContacts_listPreviewRotationShiftsCmd.MarkFlagRequired("recurrence")
	ssmContacts_listPreviewRotationShiftsCmd.MarkFlagRequired("time-zone-id")
	ssmContactsCmd.AddCommand(ssmContacts_listPreviewRotationShiftsCmd)
}
