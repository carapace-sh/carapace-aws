package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_deleteContainerGroupDefinitionCmd = &cobra.Command{
	Use:   "delete-container-group-definition",
	Short: "**This API works with the following fleet types:** Container\n\n**Request options:**\n\nDeletes a container group definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_deleteContainerGroupDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_deleteContainerGroupDefinitionCmd).Standalone()

		gamelift_deleteContainerGroupDefinitionCmd.Flags().String("name", "", "The unique identifier for the container group definition to delete.")
		gamelift_deleteContainerGroupDefinitionCmd.Flags().String("version-count-to-retain", "", "The number of most recent versions to keep while deleting all older versions.")
		gamelift_deleteContainerGroupDefinitionCmd.Flags().String("version-number", "", "The specific version to delete.")
		gamelift_deleteContainerGroupDefinitionCmd.MarkFlagRequired("name")
	})
	gameliftCmd.AddCommand(gamelift_deleteContainerGroupDefinitionCmd)
}
