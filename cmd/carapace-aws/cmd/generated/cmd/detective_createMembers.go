package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detective_createMembersCmd = &cobra.Command{
	Use:   "create-members",
	Short: "`CreateMembers` is used to send invitations to accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detective_createMembersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(detective_createMembersCmd).Standalone()

		detective_createMembersCmd.Flags().String("accounts", "", "The list of Amazon Web Services accounts to invite or to enable.")
		detective_createMembersCmd.Flags().Bool("disable-email-notification", false, "if set to `true`, then the invited accounts do not receive email notifications.")
		detective_createMembersCmd.Flags().String("graph-arn", "", "The ARN of the behavior graph.")
		detective_createMembersCmd.Flags().String("message", "", "Customized message text to include in the invitation email message to the invited member accounts.")
		detective_createMembersCmd.Flags().Bool("no-disable-email-notification", false, "if set to `true`, then the invited accounts do not receive email notifications.")
		detective_createMembersCmd.MarkFlagRequired("accounts")
		detective_createMembersCmd.MarkFlagRequired("graph-arn")
		detective_createMembersCmd.Flag("no-disable-email-notification").Hidden = true
	})
	detectiveCmd.AddCommand(detective_createMembersCmd)
}
