package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Retrieves the tags associated with the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_listTagsForResourceCmd).Standalone()

	networkFirewall_listTagsForResourceCmd.Flags().String("max-results", "", "The maximum number of objects that you want Network Firewall to return for this request.")
	networkFirewall_listTagsForResourceCmd.Flags().String("next-token", "", "When you request a list of objects with a `MaxResults` setting, if the number of objects that are still available for retrieval exceeds the maximum you requested, Network Firewall returns a `NextToken` value in the response.")
	networkFirewall_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	networkFirewall_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	networkFirewallCmd.AddCommand(networkFirewall_listTagsForResourceCmd)
}
