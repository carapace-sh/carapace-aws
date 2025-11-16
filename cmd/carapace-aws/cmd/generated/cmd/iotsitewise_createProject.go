package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_createProjectCmd = &cobra.Command{
	Use:   "create-project",
	Short: "Creates a project in the specified portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_createProjectCmd).Standalone()

	iotsitewise_createProjectCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
	iotsitewise_createProjectCmd.Flags().String("portal-id", "", "The ID of the portal in which to create the project.")
	iotsitewise_createProjectCmd.Flags().String("project-description", "", "A description for the project.")
	iotsitewise_createProjectCmd.Flags().String("project-name", "", "A friendly name for the project.")
	iotsitewise_createProjectCmd.Flags().String("tags", "", "A list of key-value pairs that contain metadata for the project.")
	iotsitewise_createProjectCmd.MarkFlagRequired("portal-id")
	iotsitewise_createProjectCmd.MarkFlagRequired("project-name")
	iotsitewiseCmd.AddCommand(iotsitewise_createProjectCmd)
}
