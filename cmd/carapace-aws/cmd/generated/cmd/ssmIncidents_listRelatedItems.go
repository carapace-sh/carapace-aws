package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_listRelatedItemsCmd = &cobra.Command{
	Use:   "list-related-items",
	Short: "List all related items for an incident record.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_listRelatedItemsCmd).Standalone()

	ssmIncidents_listRelatedItemsCmd.Flags().String("incident-record-arn", "", "The Amazon Resource Name (ARN) of the incident record containing the listed related items.")
	ssmIncidents_listRelatedItemsCmd.Flags().String("max-results", "", "The maximum number of related items per page.")
	ssmIncidents_listRelatedItemsCmd.Flags().String("next-token", "", "The pagination token for the next set of items to return.")
	ssmIncidents_listRelatedItemsCmd.MarkFlagRequired("incident-record-arn")
	ssmIncidentsCmd.AddCommand(ssmIncidents_listRelatedItemsCmd)
}
