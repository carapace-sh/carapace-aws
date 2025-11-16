package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_listMembersCmd = &cobra.Command{
	Use:   "list-members",
	Short: "Lists details about all member accounts for the current GuardDuty administrator account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_listMembersCmd).Standalone()

	guardduty_listMembersCmd.Flags().String("detector-id", "", "The unique ID of the detector that is associated with the member.")
	guardduty_listMembersCmd.Flags().String("max-results", "", "You can use this parameter to indicate the maximum number of items you want in the response.")
	guardduty_listMembersCmd.Flags().String("next-token", "", "You can use this parameter when paginating results.")
	guardduty_listMembersCmd.Flags().String("only-associated", "", "Specifies whether to only return associated members or to return all members (including members who haven't been invited yet or have been disassociated).")
	guardduty_listMembersCmd.MarkFlagRequired("detector-id")
	guarddutyCmd.AddCommand(guardduty_listMembersCmd)
}
