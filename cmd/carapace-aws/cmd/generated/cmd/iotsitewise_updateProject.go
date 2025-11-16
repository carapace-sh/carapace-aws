package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_updateProjectCmd = &cobra.Command{
	Use:   "update-project",
	Short: "Updates an IoT SiteWise Monitor project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_updateProjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_updateProjectCmd).Standalone()

		iotsitewise_updateProjectCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
		iotsitewise_updateProjectCmd.Flags().String("project-description", "", "A new description for the project.")
		iotsitewise_updateProjectCmd.Flags().String("project-id", "", "The ID of the project to update.")
		iotsitewise_updateProjectCmd.Flags().String("project-name", "", "A new friendly name for the project.")
		iotsitewise_updateProjectCmd.MarkFlagRequired("project-id")
		iotsitewise_updateProjectCmd.MarkFlagRequired("project-name")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_updateProjectCmd)
}
