package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_listConfiguredAudienceModelAssociationsCmd = &cobra.Command{
	Use:   "list-configured-audience-model-associations",
	Short: "Lists information about requested configured audience model associations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_listConfiguredAudienceModelAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_listConfiguredAudienceModelAssociationsCmd).Standalone()

		cleanrooms_listConfiguredAudienceModelAssociationsCmd.Flags().String("max-results", "", "The maximum number of results that are returned for an API request call.")
		cleanrooms_listConfiguredAudienceModelAssociationsCmd.Flags().String("membership-identifier", "", "A unique identifier for a membership that contains the configured audience model associations that you want to retrieve.")
		cleanrooms_listConfiguredAudienceModelAssociationsCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
		cleanrooms_listConfiguredAudienceModelAssociationsCmd.MarkFlagRequired("membership-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_listConfiguredAudienceModelAssociationsCmd)
}
