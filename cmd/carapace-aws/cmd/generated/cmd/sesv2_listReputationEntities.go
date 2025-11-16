package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_listReputationEntitiesCmd = &cobra.Command{
	Use:   "list-reputation-entities",
	Short: "List reputation entities in your Amazon SES account in the current Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_listReputationEntitiesCmd).Standalone()

	sesv2_listReputationEntitiesCmd.Flags().String("filter", "", "An object that contains filters to apply when listing reputation entities.")
	sesv2_listReputationEntitiesCmd.Flags().String("next-token", "", "A token returned from a previous call to `ListReputationEntities` to indicate the position in the list of reputation entities.")
	sesv2_listReputationEntitiesCmd.Flags().String("page-size", "", "The number of results to show in a single call to `ListReputationEntities`.")
	sesv2Cmd.AddCommand(sesv2_listReputationEntitiesCmd)
}
