package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_createEnvironmentProfileCmd = &cobra.Command{
	Use:   "create-environment-profile",
	Short: "Creates an Amazon DataZone environment profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_createEnvironmentProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_createEnvironmentProfileCmd).Standalone()

		datazone_createEnvironmentProfileCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account in which the Amazon DataZone environment is created.")
		datazone_createEnvironmentProfileCmd.Flags().String("aws-account-region", "", "The Amazon Web Services region in which this environment profile is created.")
		datazone_createEnvironmentProfileCmd.Flags().String("description", "", "The description of this Amazon DataZone environment profile.")
		datazone_createEnvironmentProfileCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which this environment profile is created.")
		datazone_createEnvironmentProfileCmd.Flags().String("environment-blueprint-identifier", "", "The ID of the blueprint with which this environment profile is created.")
		datazone_createEnvironmentProfileCmd.Flags().String("name", "", "The name of this Amazon DataZone environment profile.")
		datazone_createEnvironmentProfileCmd.Flags().String("project-identifier", "", "The identifier of the project in which to create the environment profile.")
		datazone_createEnvironmentProfileCmd.Flags().String("user-parameters", "", "The user parameters of this Amazon DataZone environment profile.")
		datazone_createEnvironmentProfileCmd.MarkFlagRequired("domain-identifier")
		datazone_createEnvironmentProfileCmd.MarkFlagRequired("environment-blueprint-identifier")
		datazone_createEnvironmentProfileCmd.MarkFlagRequired("name")
		datazone_createEnvironmentProfileCmd.MarkFlagRequired("project-identifier")
	})
	datazoneCmd.AddCommand(datazone_createEnvironmentProfileCmd)
}
