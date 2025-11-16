package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_createMemberCmd = &cobra.Command{
	Use:   "create-member",
	Short: "Associates an account with an Amazon Macie administrator account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_createMemberCmd).Standalone()

	macie2_createMemberCmd.Flags().String("account", "", "The details of the account to associate with the administrator account.")
	macie2_createMemberCmd.Flags().String("tags", "", "A map of key-value pairs that specifies the tags to associate with the account in Amazon Macie.")
	macie2_createMemberCmd.MarkFlagRequired("account")
	macie2Cmd.AddCommand(macie2_createMemberCmd)
}
