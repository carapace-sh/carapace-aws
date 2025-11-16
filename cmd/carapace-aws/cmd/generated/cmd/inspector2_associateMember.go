package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_associateMemberCmd = &cobra.Command{
	Use:   "associate-member",
	Short: "Associates an Amazon Web Services account with an Amazon Inspector delegated administrator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_associateMemberCmd).Standalone()

	inspector2_associateMemberCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the member account to be associated.")
	inspector2_associateMemberCmd.MarkFlagRequired("account-id")
	inspector2Cmd.AddCommand(inspector2_associateMemberCmd)
}
