package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gameliftstreams_listApplicationsCmd = &cobra.Command{
	Use:   "list-applications",
	Short: "Retrieves a list of all Amazon GameLift Streams applications that are associated with the Amazon Web Services account in use.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gameliftstreams_listApplicationsCmd).Standalone()

	gameliftstreams_listApplicationsCmd.Flags().String("max-results", "", "The number of results to return.")
	gameliftstreams_listApplicationsCmd.Flags().String("next-token", "", "The token that marks the start of the next set of results.")
	gameliftstreamsCmd.AddCommand(gameliftstreams_listApplicationsCmd)
}
