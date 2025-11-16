package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_updateEnvironmentProfileCmd = &cobra.Command{
	Use:   "update-environment-profile",
	Short: "Updates the specified environment profile in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_updateEnvironmentProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_updateEnvironmentProfileCmd).Standalone()

		datazone_updateEnvironmentProfileCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account in which a specified environment profile is to be udpated.")
		datazone_updateEnvironmentProfileCmd.Flags().String("aws-account-region", "", "The Amazon Web Services Region in which a specified environment profile is to be updated.")
		datazone_updateEnvironmentProfileCmd.Flags().String("description", "", "The description to be updated as part of the `UpdateEnvironmentProfile` action.")
		datazone_updateEnvironmentProfileCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain in which an environment profile is to be updated.")
		datazone_updateEnvironmentProfileCmd.Flags().String("identifier", "", "The identifier of the environment profile that is to be updated.")
		datazone_updateEnvironmentProfileCmd.Flags().String("name", "", "The name to be updated as part of the `UpdateEnvironmentProfile` action.")
		datazone_updateEnvironmentProfileCmd.Flags().String("user-parameters", "", "The user parameters to be updated as part of the `UpdateEnvironmentProfile` action.")
		datazone_updateEnvironmentProfileCmd.MarkFlagRequired("domain-identifier")
		datazone_updateEnvironmentProfileCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_updateEnvironmentProfileCmd)
}
