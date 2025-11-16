package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_updateMemberSessionCmd = &cobra.Command{
	Use:   "update-member-session",
	Short: "Enables an Amazon Macie administrator to suspend or re-enable Macie for a member account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_updateMemberSessionCmd).Standalone()

	macie2_updateMemberSessionCmd.Flags().String("id", "", "The unique identifier for the Amazon Macie resource that the request applies to.")
	macie2_updateMemberSessionCmd.Flags().String("status", "", "Specifies the new status for the account.")
	macie2_updateMemberSessionCmd.MarkFlagRequired("id")
	macie2_updateMemberSessionCmd.MarkFlagRequired("status")
	macie2Cmd.AddCommand(macie2_updateMemberSessionCmd)
}
