package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_updateProjectCmd = &cobra.Command{
	Use:   "update-project",
	Short: "Updates the specified project in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_updateProjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_updateProjectCmd).Standalone()

		datazone_updateProjectCmd.Flags().String("description", "", "The description to be updated as part of the `UpdateProject` action.")
		datazone_updateProjectCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain where a project is being updated.")
		datazone_updateProjectCmd.Flags().String("domain-unit-id", "", "The ID of the domain unit.")
		datazone_updateProjectCmd.Flags().String("environment-deployment-details", "", "The environment deployment details of the project.")
		datazone_updateProjectCmd.Flags().String("glossary-terms", "", "The glossary terms to be updated as part of the `UpdateProject` action.")
		datazone_updateProjectCmd.Flags().String("identifier", "", "The identifier of the project that is to be updated.")
		datazone_updateProjectCmd.Flags().String("name", "", "The name to be updated as part of the `UpdateProject` action.")
		datazone_updateProjectCmd.Flags().String("project-profile-version", "", "The project profile version to which the project should be updated.")
		datazone_updateProjectCmd.Flags().String("resource-tags", "", "The resource tags of the project.")
		datazone_updateProjectCmd.Flags().String("user-parameters", "", "The user parameters of the project.")
		datazone_updateProjectCmd.MarkFlagRequired("domain-identifier")
		datazone_updateProjectCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_updateProjectCmd)
}
