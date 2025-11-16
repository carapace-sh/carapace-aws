package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_createUserProfileCmd = &cobra.Command{
	Use:   "create-user-profile",
	Short: "Creates a user profile in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_createUserProfileCmd).Standalone()

	datazone_createUserProfileCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
	datazone_createUserProfileCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain in which a user profile is created.")
	datazone_createUserProfileCmd.Flags().String("user-identifier", "", "The identifier of the user for which the user profile is created.")
	datazone_createUserProfileCmd.Flags().String("user-type", "", "The user type of the user for which the user profile is created.")
	datazone_createUserProfileCmd.MarkFlagRequired("domain-identifier")
	datazone_createUserProfileCmd.MarkFlagRequired("user-identifier")
	datazoneCmd.AddCommand(datazone_createUserProfileCmd)
}
