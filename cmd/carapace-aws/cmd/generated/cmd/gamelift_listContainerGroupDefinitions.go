package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_listContainerGroupDefinitionsCmd = &cobra.Command{
	Use:   "list-container-group-definitions",
	Short: "**This API works with the following fleet types:** Container\n\nRetrieves container group definitions for the Amazon Web Services account and Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_listContainerGroupDefinitionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_listContainerGroupDefinitionsCmd).Standalone()

		gamelift_listContainerGroupDefinitionsCmd.Flags().String("container-group-type", "", "The type of container group to retrieve.")
		gamelift_listContainerGroupDefinitionsCmd.Flags().String("limit", "", "The maximum number of results to return.")
		gamelift_listContainerGroupDefinitionsCmd.Flags().String("next-token", "", "A token that indicates the start of the next sequential page of results.")
	})
	gameliftCmd.AddCommand(gamelift_listContainerGroupDefinitionsCmd)
}
