package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gameliftstreams_listStreamGroupsCmd = &cobra.Command{
	Use:   "list-stream-groups",
	Short: "Retrieves a list of all Amazon GameLift Streams stream groups that are associated with the Amazon Web Services account in use.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gameliftstreams_listStreamGroupsCmd).Standalone()

	gameliftstreams_listStreamGroupsCmd.Flags().String("max-results", "", "The number of results to return.")
	gameliftstreams_listStreamGroupsCmd.Flags().String("next-token", "", "A token that marks the start of the next set of results.")
	gameliftstreamsCmd.AddCommand(gameliftstreams_listStreamGroupsCmd)
}
