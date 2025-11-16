package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_createProjectCmd = &cobra.Command{
	Use:   "create-project",
	Short: "Creates a project, which is the logical object in Evidently that can contain features, launches, and experiments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_createProjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(evidently_createProjectCmd).Standalone()

		evidently_createProjectCmd.Flags().String("app-config-resource", "", "Use this parameter if the project will use *client-side evaluation powered by AppConfig*.")
		evidently_createProjectCmd.Flags().String("data-delivery", "", "A structure that contains information about where Evidently is to store evaluation events for longer term storage, if you choose to do so.")
		evidently_createProjectCmd.Flags().String("description", "", "An optional description of the project.")
		evidently_createProjectCmd.Flags().String("name", "", "The name for the project.")
		evidently_createProjectCmd.Flags().String("tags", "", "Assigns one or more tags (key-value pairs) to the project.")
		evidently_createProjectCmd.MarkFlagRequired("name")
	})
	evidentlyCmd.AddCommand(evidently_createProjectCmd)
}
