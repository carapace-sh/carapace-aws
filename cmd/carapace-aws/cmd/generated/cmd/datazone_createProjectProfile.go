package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_createProjectProfileCmd = &cobra.Command{
	Use:   "create-project-profile",
	Short: "Creates a project profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_createProjectProfileCmd).Standalone()

	datazone_createProjectProfileCmd.Flags().Bool("allow-custom-project-resource-tags", false, "Specifies whether custom project resource tags are supported.")
	datazone_createProjectProfileCmd.Flags().String("description", "", "A description of a project profile.")
	datazone_createProjectProfileCmd.Flags().String("domain-identifier", "", "A domain ID of the project profile.")
	datazone_createProjectProfileCmd.Flags().String("domain-unit-identifier", "", "A domain unit ID of the project profile.")
	datazone_createProjectProfileCmd.Flags().String("environment-configurations", "", "Environment configurations of the project profile.")
	datazone_createProjectProfileCmd.Flags().String("name", "", "Project profile name.")
	datazone_createProjectProfileCmd.Flags().Bool("no-allow-custom-project-resource-tags", false, "Specifies whether custom project resource tags are supported.")
	datazone_createProjectProfileCmd.Flags().String("project-resource-tags", "", "The resource tags of the project profile.")
	datazone_createProjectProfileCmd.Flags().String("project-resource-tags-description", "", "Field viewable through the UI that provides a project user with the allowed resource tag specifications.")
	datazone_createProjectProfileCmd.Flags().String("status", "", "Project profile status.")
	datazone_createProjectProfileCmd.MarkFlagRequired("domain-identifier")
	datazone_createProjectProfileCmd.MarkFlagRequired("name")
	datazone_createProjectProfileCmd.Flag("no-allow-custom-project-resource-tags").Hidden = true
	datazoneCmd.AddCommand(datazone_createProjectProfileCmd)
}
