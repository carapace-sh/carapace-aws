package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_updateProjectProfileCmd = &cobra.Command{
	Use:   "update-project-profile",
	Short: "Updates a project profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_updateProjectProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_updateProjectProfileCmd).Standalone()

		datazone_updateProjectProfileCmd.Flags().Bool("allow-custom-project-resource-tags", false, "Specifies whether custom project resource tags are supported.")
		datazone_updateProjectProfileCmd.Flags().String("description", "", "The description of a project profile.")
		datazone_updateProjectProfileCmd.Flags().String("domain-identifier", "", "The ID of the domain where a project profile is to be updated.")
		datazone_updateProjectProfileCmd.Flags().String("domain-unit-identifier", "", "The ID of the domain unit where a project profile is to be updated.")
		datazone_updateProjectProfileCmd.Flags().String("environment-configurations", "", "The environment configurations of a project profile.")
		datazone_updateProjectProfileCmd.Flags().String("identifier", "", "The ID of a project profile that is to be updated.")
		datazone_updateProjectProfileCmd.Flags().String("name", "", "The name of a project profile.")
		datazone_updateProjectProfileCmd.Flags().Bool("no-allow-custom-project-resource-tags", false, "Specifies whether custom project resource tags are supported.")
		datazone_updateProjectProfileCmd.Flags().String("project-resource-tags", "", "The resource tags of the project profile.")
		datazone_updateProjectProfileCmd.Flags().String("project-resource-tags-description", "", "Field viewable through the UI that provides a project user with the allowed resource tag specifications.")
		datazone_updateProjectProfileCmd.Flags().String("status", "", "The status of a project profile.")
		datazone_updateProjectProfileCmd.MarkFlagRequired("domain-identifier")
		datazone_updateProjectProfileCmd.MarkFlagRequired("identifier")
		datazone_updateProjectProfileCmd.Flag("no-allow-custom-project-resource-tags").Hidden = true
	})
	datazoneCmd.AddCommand(datazone_updateProjectProfileCmd)
}
