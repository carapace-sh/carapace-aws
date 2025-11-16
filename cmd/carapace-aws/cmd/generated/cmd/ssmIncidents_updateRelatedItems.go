package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_updateRelatedItemsCmd = &cobra.Command{
	Use:   "update-related-items",
	Short: "Add or remove related items from the related items tab of an incident record.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_updateRelatedItemsCmd).Standalone()

	ssmIncidents_updateRelatedItemsCmd.Flags().String("client-token", "", "A token that ensures that a client calls the operation only once with the specified details.")
	ssmIncidents_updateRelatedItemsCmd.Flags().String("incident-record-arn", "", "The Amazon Resource Name (ARN) of the incident record that contains the related items that you update.")
	ssmIncidents_updateRelatedItemsCmd.Flags().String("related-items-update", "", "Details about the item that you are add to, or delete from, an incident.")
	ssmIncidents_updateRelatedItemsCmd.MarkFlagRequired("incident-record-arn")
	ssmIncidents_updateRelatedItemsCmd.MarkFlagRequired("related-items-update")
	ssmIncidentsCmd.AddCommand(ssmIncidents_updateRelatedItemsCmd)
}
