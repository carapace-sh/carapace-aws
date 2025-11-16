package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateUserProficienciesCmd = &cobra.Command{
	Use:   "update-user-proficiencies",
	Short: "Updates the properties associated with the proficiencies of a user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateUserProficienciesCmd).Standalone()

	connect_updateUserProficienciesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_updateUserProficienciesCmd.Flags().String("user-id", "", "The identifier of the user account.")
	connect_updateUserProficienciesCmd.Flags().String("user-proficiencies", "", "The proficiencies to be updated for the user.")
	connect_updateUserProficienciesCmd.MarkFlagRequired("instance-id")
	connect_updateUserProficienciesCmd.MarkFlagRequired("user-id")
	connect_updateUserProficienciesCmd.MarkFlagRequired("user-proficiencies")
	connectCmd.AddCommand(connect_updateUserProficienciesCmd)
}
