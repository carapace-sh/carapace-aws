package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_associateUserProficienciesCmd = &cobra.Command{
	Use:   "associate-user-proficiencies",
	Short: "Associates a set of proficiencies with a user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_associateUserProficienciesCmd).Standalone()

	connect_associateUserProficienciesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_associateUserProficienciesCmd.Flags().String("user-id", "", "The identifier of the user account.")
	connect_associateUserProficienciesCmd.Flags().String("user-proficiencies", "", "The proficiencies to associate with the user.")
	connect_associateUserProficienciesCmd.MarkFlagRequired("instance-id")
	connect_associateUserProficienciesCmd.MarkFlagRequired("user-id")
	connect_associateUserProficienciesCmd.MarkFlagRequired("user-proficiencies")
	connectCmd.AddCommand(connect_associateUserProficienciesCmd)
}
