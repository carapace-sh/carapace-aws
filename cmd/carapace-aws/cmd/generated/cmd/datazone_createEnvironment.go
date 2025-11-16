package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_createEnvironmentCmd = &cobra.Command{
	Use:   "create-environment",
	Short: "Create an Amazon DataZone environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_createEnvironmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_createEnvironmentCmd).Standalone()

		datazone_createEnvironmentCmd.Flags().String("deployment-order", "", "The deployment order of the environment.")
		datazone_createEnvironmentCmd.Flags().String("description", "", "The description of the Amazon DataZone environment.")
		datazone_createEnvironmentCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain in which the environment is created.")
		datazone_createEnvironmentCmd.Flags().String("environment-account-identifier", "", "The ID of the account in which the environment is being created.")
		datazone_createEnvironmentCmd.Flags().String("environment-account-region", "", "The region of the account in which the environment is being created.")
		datazone_createEnvironmentCmd.Flags().String("environment-blueprint-identifier", "", "The ID of the blueprint with which the environment is being created.")
		datazone_createEnvironmentCmd.Flags().String("environment-configuration-id", "", "The configuration ID of the environment.")
		datazone_createEnvironmentCmd.Flags().String("environment-profile-identifier", "", "The identifier of the environment profile that is used to create this Amazon DataZone environment.")
		datazone_createEnvironmentCmd.Flags().String("glossary-terms", "", "The glossary terms that can be used in this Amazon DataZone environment.")
		datazone_createEnvironmentCmd.Flags().String("name", "", "The name of the Amazon DataZone environment.")
		datazone_createEnvironmentCmd.Flags().String("project-identifier", "", "The identifier of the Amazon DataZone project in which this environment is created.")
		datazone_createEnvironmentCmd.Flags().String("user-parameters", "", "The user parameters of this Amazon DataZone environment.")
		datazone_createEnvironmentCmd.MarkFlagRequired("domain-identifier")
		datazone_createEnvironmentCmd.MarkFlagRequired("name")
		datazone_createEnvironmentCmd.MarkFlagRequired("project-identifier")
	})
	datazoneCmd.AddCommand(datazone_createEnvironmentCmd)
}
