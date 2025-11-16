package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_changeResourceRecordSetsCmd = &cobra.Command{
	Use:   "change-resource-record-sets",
	Short: "Creates, changes, or deletes a resource record set, which contains authoritative DNS information for a specified domain name or subdomain name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_changeResourceRecordSetsCmd).Standalone()

	route53_changeResourceRecordSetsCmd.Flags().String("change-batch", "", "A complex type that contains an optional comment and the `Changes` element.")
	route53_changeResourceRecordSetsCmd.Flags().String("hosted-zone-id", "", "The ID of the hosted zone that contains the resource record sets that you want to change.")
	route53_changeResourceRecordSetsCmd.MarkFlagRequired("change-batch")
	route53_changeResourceRecordSetsCmd.MarkFlagRequired("hosted-zone-id")
	route53Cmd.AddCommand(route53_changeResourceRecordSetsCmd)
}
