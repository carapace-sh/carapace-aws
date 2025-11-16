package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_deleteEnvironmentProfileCmd = &cobra.Command{
	Use:   "delete-environment-profile",
	Short: "Deletes an environment profile in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_deleteEnvironmentProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_deleteEnvironmentProfileCmd).Standalone()

		datazone_deleteEnvironmentProfileCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which the environment profile is deleted.")
		datazone_deleteEnvironmentProfileCmd.Flags().String("identifier", "", "The ID of the environment profile that is deleted.")
		datazone_deleteEnvironmentProfileCmd.MarkFlagRequired("domain-identifier")
		datazone_deleteEnvironmentProfileCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_deleteEnvironmentProfileCmd)
}
