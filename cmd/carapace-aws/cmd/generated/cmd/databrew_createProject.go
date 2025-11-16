package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_createProjectCmd = &cobra.Command{
	Use:   "create-project",
	Short: "Creates a new DataBrew project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_createProjectCmd).Standalone()

	databrew_createProjectCmd.Flags().String("dataset-name", "", "The name of an existing dataset to associate this project with.")
	databrew_createProjectCmd.Flags().String("name", "", "A unique name for the new project.")
	databrew_createProjectCmd.Flags().String("recipe-name", "", "The name of an existing recipe to associate with the project.")
	databrew_createProjectCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role to be assumed for this request.")
	databrew_createProjectCmd.Flags().String("sample", "", "")
	databrew_createProjectCmd.Flags().String("tags", "", "Metadata tags to apply to this project.")
	databrew_createProjectCmd.MarkFlagRequired("dataset-name")
	databrew_createProjectCmd.MarkFlagRequired("name")
	databrew_createProjectCmd.MarkFlagRequired("recipe-name")
	databrew_createProjectCmd.MarkFlagRequired("role-arn")
	databrewCmd.AddCommand(databrew_createProjectCmd)
}
