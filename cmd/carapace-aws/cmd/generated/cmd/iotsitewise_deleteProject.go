package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_deleteProjectCmd = &cobra.Command{
	Use:   "delete-project",
	Short: "Deletes a project from IoT SiteWise Monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_deleteProjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_deleteProjectCmd).Standalone()

		iotsitewise_deleteProjectCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
		iotsitewise_deleteProjectCmd.Flags().String("project-id", "", "The ID of the project.")
		iotsitewise_deleteProjectCmd.MarkFlagRequired("project-id")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_deleteProjectCmd)
}
