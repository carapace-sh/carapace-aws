package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_disassociateUserProficienciesCmd = &cobra.Command{
	Use:   "disassociate-user-proficiencies",
	Short: "Disassociates a set of proficiencies from a user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_disassociateUserProficienciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_disassociateUserProficienciesCmd).Standalone()

		connect_disassociateUserProficienciesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_disassociateUserProficienciesCmd.Flags().String("user-id", "", "The identifier of the user account.")
		connect_disassociateUserProficienciesCmd.Flags().String("user-proficiencies", "", "The proficiencies to disassociate from the user.")
		connect_disassociateUserProficienciesCmd.MarkFlagRequired("instance-id")
		connect_disassociateUserProficienciesCmd.MarkFlagRequired("user-id")
		connect_disassociateUserProficienciesCmd.MarkFlagRequired("user-proficiencies")
	})
	connectCmd.AddCommand(connect_disassociateUserProficienciesCmd)
}
