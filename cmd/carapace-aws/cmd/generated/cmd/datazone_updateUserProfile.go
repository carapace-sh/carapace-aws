package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_updateUserProfileCmd = &cobra.Command{
	Use:   "update-user-profile",
	Short: "Updates the specified user profile in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_updateUserProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_updateUserProfileCmd).Standalone()

		datazone_updateUserProfileCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain in which a user profile is updated.")
		datazone_updateUserProfileCmd.Flags().String("status", "", "The status of the user profile that are to be updated.")
		datazone_updateUserProfileCmd.Flags().String("type", "", "The type of the user profile that are to be updated.")
		datazone_updateUserProfileCmd.Flags().String("user-identifier", "", "The identifier of the user whose user profile is to be updated.")
		datazone_updateUserProfileCmd.MarkFlagRequired("domain-identifier")
		datazone_updateUserProfileCmd.MarkFlagRequired("status")
		datazone_updateUserProfileCmd.MarkFlagRequired("user-identifier")
	})
	datazoneCmd.AddCommand(datazone_updateUserProfileCmd)
}
