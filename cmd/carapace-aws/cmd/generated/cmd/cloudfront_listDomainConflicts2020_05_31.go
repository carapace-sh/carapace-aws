package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listDomainConflicts2020_05_31Cmd = &cobra.Command{
	Use:   "list-domain-conflicts2020_05_31",
	Short: "We recommend that you use the `ListDomainConflicts` API operation to check for domain conflicts, as it supports both standard distributions and distribution tenants.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listDomainConflicts2020_05_31Cmd).Standalone()

	cloudfront_listDomainConflicts2020_05_31Cmd.Flags().String("domain", "", "The domain to check for conflicts.")
	cloudfront_listDomainConflicts2020_05_31Cmd.Flags().String("domain-control-validation-resource", "", "The distribution resource identifier.")
	cloudfront_listDomainConflicts2020_05_31Cmd.Flags().String("marker", "", "The marker for the next set of domain conflicts.")
	cloudfront_listDomainConflicts2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of domain conflicts to return.")
	cloudfront_listDomainConflicts2020_05_31Cmd.MarkFlagRequired("domain")
	cloudfront_listDomainConflicts2020_05_31Cmd.MarkFlagRequired("domain-control-validation-resource")
	cloudfrontCmd.AddCommand(cloudfront_listDomainConflicts2020_05_31Cmd)
}
