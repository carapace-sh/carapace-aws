package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_getTestGridProjectCmd = &cobra.Command{
	Use:   "get-test-grid-project",
	Short: "Retrieves information about a Selenium testing project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_getTestGridProjectCmd).Standalone()

	devicefarm_getTestGridProjectCmd.Flags().String("project-arn", "", "The ARN of the Selenium testing project, from either [CreateTestGridProject]() or [ListTestGridProjects]().")
	devicefarm_getTestGridProjectCmd.MarkFlagRequired("project-arn")
	devicefarmCmd.AddCommand(devicefarm_getTestGridProjectCmd)
}
