package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listSubscriptionGrantsCmd = &cobra.Command{
	Use:   "list-subscription-grants",
	Short: "Lists subscription grants.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listSubscriptionGrantsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_listSubscriptionGrantsCmd).Standalone()

		datazone_listSubscriptionGrantsCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain.")
		datazone_listSubscriptionGrantsCmd.Flags().String("environment-id", "", "The identifier of the Amazon DataZone environment.")
		datazone_listSubscriptionGrantsCmd.Flags().String("max-results", "", "The maximum number of subscription grants to return in a single call to `ListSubscriptionGrants`.")
		datazone_listSubscriptionGrantsCmd.Flags().String("next-token", "", "When the number of subscription grants is greater than the default value for the `MaxResults` parameter, or if you explicitly specify a value for `MaxResults` that is less than the number of subscription grants, the response includes a pagination token named `NextToken`.")
		datazone_listSubscriptionGrantsCmd.Flags().String("owning-group-id", "", "The ID of the owning group.")
		datazone_listSubscriptionGrantsCmd.Flags().String("owning-project-id", "", "The ID of the owning project of the subscription grants.")
		datazone_listSubscriptionGrantsCmd.Flags().String("owning-user-id", "", "The ID of the owning user.")
		datazone_listSubscriptionGrantsCmd.Flags().String("sort-by", "", "Specifies the way of sorting the results of this action.")
		datazone_listSubscriptionGrantsCmd.Flags().String("sort-order", "", "Specifies the sort order of this action.")
		datazone_listSubscriptionGrantsCmd.Flags().String("subscribed-listing-id", "", "The identifier of the subscribed listing.")
		datazone_listSubscriptionGrantsCmd.Flags().String("subscription-id", "", "The identifier of the subscription.")
		datazone_listSubscriptionGrantsCmd.Flags().String("subscription-target-id", "", "The identifier of the subscription target.")
		datazone_listSubscriptionGrantsCmd.MarkFlagRequired("domain-identifier")
	})
	datazoneCmd.AddCommand(datazone_listSubscriptionGrantsCmd)
}
