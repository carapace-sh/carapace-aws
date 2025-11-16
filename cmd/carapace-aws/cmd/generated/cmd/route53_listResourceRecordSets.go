package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_listResourceRecordSetsCmd = &cobra.Command{
	Use:   "list-resource-record-sets",
	Short: "Lists the resource record sets in a specified hosted zone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_listResourceRecordSetsCmd).Standalone()

	route53_listResourceRecordSetsCmd.Flags().String("hosted-zone-id", "", "The ID of the hosted zone that contains the resource record sets that you want to list.")
	route53_listResourceRecordSetsCmd.Flags().String("max-items", "", "(Optional) The maximum number of resource records sets to include in the response body for this request.")
	route53_listResourceRecordSetsCmd.Flags().String("start-record-identifier", "", "*Resource record sets that have a routing policy other than simple:* If results were truncated for a given DNS name and type, specify the value of `NextRecordIdentifier` from the previous response to get the next resource record set that has the current DNS name and type.")
	route53_listResourceRecordSetsCmd.Flags().String("start-record-name", "", "The first name in the lexicographic ordering of resource record sets that you want to list.")
	route53_listResourceRecordSetsCmd.Flags().String("start-record-type", "", "The type of resource record set to begin the record listing from.")
	route53_listResourceRecordSetsCmd.MarkFlagRequired("hosted-zone-id")
	route53Cmd.AddCommand(route53_listResourceRecordSetsCmd)
}
