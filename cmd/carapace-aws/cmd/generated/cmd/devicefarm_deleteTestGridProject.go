package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_deleteTestGridProjectCmd = &cobra.Command{
	Use:   "delete-test-grid-project",
	Short: "Deletes a Selenium testing project and all content generated under it.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_deleteTestGridProjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_deleteTestGridProjectCmd).Standalone()

		devicefarm_deleteTestGridProjectCmd.Flags().String("project-arn", "", "The ARN of the project to delete, from [CreateTestGridProject]() or [ListTestGridProjects]().")
		devicefarm_deleteTestGridProjectCmd.MarkFlagRequired("project-arn")
	})
	devicefarmCmd.AddCommand(devicefarm_deleteTestGridProjectCmd)
}
