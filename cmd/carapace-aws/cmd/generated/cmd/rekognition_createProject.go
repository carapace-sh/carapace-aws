package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_createProjectCmd = &cobra.Command{
	Use:   "create-project",
	Short: "Creates a new Amazon Rekognition project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_createProjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_createProjectCmd).Standalone()

		rekognition_createProjectCmd.Flags().String("auto-update", "", "Specifies whether automatic retraining should be attempted for the versions of the project.")
		rekognition_createProjectCmd.Flags().String("feature", "", "Specifies feature that is being customized.")
		rekognition_createProjectCmd.Flags().String("project-name", "", "The name of the project to create.")
		rekognition_createProjectCmd.Flags().String("tags", "", "A set of tags (key-value pairs) that you want to attach to the project.")
		rekognition_createProjectCmd.MarkFlagRequired("project-name")
	})
	rekognitionCmd.AddCommand(rekognition_createProjectCmd)
}
