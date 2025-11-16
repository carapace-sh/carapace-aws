package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_deleteMemberCmd = &cobra.Command{
	Use:   "delete-member",
	Short: "Deletes the association between an Amazon Macie administrator account and an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_deleteMemberCmd).Standalone()

	macie2_deleteMemberCmd.Flags().String("id", "", "The unique identifier for the Amazon Macie resource that the request applies to.")
	macie2_deleteMemberCmd.MarkFlagRequired("id")
	macie2Cmd.AddCommand(macie2_deleteMemberCmd)
}
