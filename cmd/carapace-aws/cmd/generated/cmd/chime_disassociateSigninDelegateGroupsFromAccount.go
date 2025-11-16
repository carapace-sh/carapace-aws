package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_disassociateSigninDelegateGroupsFromAccountCmd = &cobra.Command{
	Use:   "disassociate-signin-delegate-groups-from-account",
	Short: "Disassociates the specified sign-in delegate groups from the specified Amazon Chime account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_disassociateSigninDelegateGroupsFromAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_disassociateSigninDelegateGroupsFromAccountCmd).Standalone()

		chime_disassociateSigninDelegateGroupsFromAccountCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
		chime_disassociateSigninDelegateGroupsFromAccountCmd.Flags().String("group-names", "", "The sign-in delegate group names.")
		chime_disassociateSigninDelegateGroupsFromAccountCmd.MarkFlagRequired("account-id")
		chime_disassociateSigninDelegateGroupsFromAccountCmd.MarkFlagRequired("group-names")
	})
	chimeCmd.AddCommand(chime_disassociateSigninDelegateGroupsFromAccountCmd)
}
