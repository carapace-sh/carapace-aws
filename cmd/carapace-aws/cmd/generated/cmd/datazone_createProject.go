package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_createProjectCmd = &cobra.Command{
	Use:   "create-project",
	Short: "Creates an Amazon DataZone project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_createProjectCmd).Standalone()

	datazone_createProjectCmd.Flags().String("description", "", "The description of the Amazon DataZone project.")
	datazone_createProjectCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which this project is created.")
	datazone_createProjectCmd.Flags().String("domain-unit-id", "", "The ID of the domain unit.")
	datazone_createProjectCmd.Flags().String("glossary-terms", "", "The glossary terms that can be used in this Amazon DataZone project.")
	datazone_createProjectCmd.Flags().String("name", "", "The name of the Amazon DataZone project.")
	datazone_createProjectCmd.Flags().String("project-profile-id", "", "The ID of the project profile.")
	datazone_createProjectCmd.Flags().String("resource-tags", "", "The resource tags of the project.")
	datazone_createProjectCmd.Flags().String("user-parameters", "", "The user parameters of the project.")
	datazone_createProjectCmd.MarkFlagRequired("domain-identifier")
	datazone_createProjectCmd.MarkFlagRequired("name")
	datazoneCmd.AddCommand(datazone_createProjectCmd)
}
