package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_associateSigninDelegateGroupsWithAccountCmd = &cobra.Command{
	Use:   "associate-signin-delegate-groups-with-account",
	Short: "Associates the specified sign-in delegate groups with the specified Amazon Chime account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_associateSigninDelegateGroupsWithAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_associateSigninDelegateGroupsWithAccountCmd).Standalone()

		chime_associateSigninDelegateGroupsWithAccountCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
		chime_associateSigninDelegateGroupsWithAccountCmd.Flags().String("signin-delegate-groups", "", "The sign-in delegate groups.")
		chime_associateSigninDelegateGroupsWithAccountCmd.MarkFlagRequired("account-id")
		chime_associateSigninDelegateGroupsWithAccountCmd.MarkFlagRequired("signin-delegate-groups")
	})
	chimeCmd.AddCommand(chime_associateSigninDelegateGroupsWithAccountCmd)
}
