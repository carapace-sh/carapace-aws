package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_listContainerFleetsCmd = &cobra.Command{
	Use:   "list-container-fleets",
	Short: "**This API works with the following fleet types:** Container\n\nRetrieves a collection of container fleet resources in an Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_listContainerFleetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_listContainerFleetsCmd).Standalone()

		gamelift_listContainerFleetsCmd.Flags().String("container-group-definition-name", "", "The container group definition to filter the list on.")
		gamelift_listContainerFleetsCmd.Flags().String("limit", "", "The maximum number of results to return.")
		gamelift_listContainerFleetsCmd.Flags().String("next-token", "", "A token that indicates the start of the next sequential page of results.")
	})
	gameliftCmd.AddCommand(gamelift_listContainerFleetsCmd)
}
