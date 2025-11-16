package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listSubscriptionsCmd = &cobra.Command{
	Use:   "list-subscriptions",
	Short: "Lists subscriptions in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listSubscriptionsCmd).Standalone()

	datazone_listSubscriptionsCmd.Flags().String("approver-project-id", "", "The identifier of the project for the subscription's approver.")
	datazone_listSubscriptionsCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain.")
	datazone_listSubscriptionsCmd.Flags().String("max-results", "", "The maximum number of subscriptions to return in a single call to `ListSubscriptions`.")
	datazone_listSubscriptionsCmd.Flags().String("next-token", "", "When the number of subscriptions is greater than the default value for the `MaxResults` parameter, or if you explicitly specify a value for `MaxResults` that is less than the number of subscriptions, the response includes a pagination token named `NextToken`.")
	datazone_listSubscriptionsCmd.Flags().String("owning-group-id", "", "The ID of the owning group.")
	datazone_listSubscriptionsCmd.Flags().String("owning-project-id", "", "The identifier of the owning project.")
	datazone_listSubscriptionsCmd.Flags().String("owning-user-id", "", "The ID of the owning user.")
	datazone_listSubscriptionsCmd.Flags().String("sort-by", "", "Specifies the way in which the results of this action are to be sorted.")
	datazone_listSubscriptionsCmd.Flags().String("sort-order", "", "Specifies the sort order for the results of this action.")
	datazone_listSubscriptionsCmd.Flags().String("status", "", "The status of the subscriptions that you want to list.")
	datazone_listSubscriptionsCmd.Flags().String("subscribed-listing-id", "", "The identifier of the subscribed listing for the subscriptions that you want to list.")
	datazone_listSubscriptionsCmd.Flags().String("subscription-request-identifier", "", "The identifier of the subscription request for the subscriptions that you want to list.")
	datazone_listSubscriptionsCmd.MarkFlagRequired("domain-identifier")
	datazoneCmd.AddCommand(datazone_listSubscriptionsCmd)
}
