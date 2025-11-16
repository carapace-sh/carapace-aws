package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listSubscriptionRequestsCmd = &cobra.Command{
	Use:   "list-subscription-requests",
	Short: "Lists Amazon DataZone subscription requests.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listSubscriptionRequestsCmd).Standalone()

	datazone_listSubscriptionRequestsCmd.Flags().String("approver-project-id", "", "The identifier of the subscription request approver's project.")
	datazone_listSubscriptionRequestsCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain.")
	datazone_listSubscriptionRequestsCmd.Flags().String("max-results", "", "The maximum number of subscription requests to return in a single call to `ListSubscriptionRequests`.")
	datazone_listSubscriptionRequestsCmd.Flags().String("next-token", "", "When the number of subscription requests is greater than the default value for the `MaxResults` parameter, or if you explicitly specify a value for `MaxResults` that is less than the number of subscription requests, the response includes a pagination token named `NextToken`.")
	datazone_listSubscriptionRequestsCmd.Flags().String("owning-group-id", "", "The ID of the owning group.")
	datazone_listSubscriptionRequestsCmd.Flags().String("owning-project-id", "", "The identifier of the project for the subscription requests.")
	datazone_listSubscriptionRequestsCmd.Flags().String("owning-user-id", "", "The ID of the owning user.")
	datazone_listSubscriptionRequestsCmd.Flags().String("sort-by", "", "Specifies the way to sort the results of this action.")
	datazone_listSubscriptionRequestsCmd.Flags().String("sort-order", "", "Specifies the sort order for the results of this action.")
	datazone_listSubscriptionRequestsCmd.Flags().String("status", "", "Specifies the status of the subscription requests.")
	datazone_listSubscriptionRequestsCmd.Flags().String("subscribed-listing-id", "", "The identifier of the subscribed listing.")
	datazone_listSubscriptionRequestsCmd.MarkFlagRequired("domain-identifier")
	datazoneCmd.AddCommand(datazone_listSubscriptionRequestsCmd)
}
