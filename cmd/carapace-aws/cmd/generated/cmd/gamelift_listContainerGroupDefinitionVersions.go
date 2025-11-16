package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_listContainerGroupDefinitionVersionsCmd = &cobra.Command{
	Use:   "list-container-group-definition-versions",
	Short: "**This API works with the following fleet types:** Container\n\nRetrieves all versions of a container group definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_listContainerGroupDefinitionVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_listContainerGroupDefinitionVersionsCmd).Standalone()

		gamelift_listContainerGroupDefinitionVersionsCmd.Flags().String("limit", "", "The maximum number of results to return.")
		gamelift_listContainerGroupDefinitionVersionsCmd.Flags().String("name", "", "The unique identifier for the container group definition to retrieve properties for.")
		gamelift_listContainerGroupDefinitionVersionsCmd.Flags().String("next-token", "", "A token that indicates the start of the next sequential page of results.")
		gamelift_listContainerGroupDefinitionVersionsCmd.MarkFlagRequired("name")
	})
	gameliftCmd.AddCommand(gamelift_listContainerGroupDefinitionVersionsCmd)
}
