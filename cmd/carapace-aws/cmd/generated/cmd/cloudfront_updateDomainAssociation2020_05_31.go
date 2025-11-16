package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_updateDomainAssociation2020_05_31Cmd = &cobra.Command{
	Use:   "update-domain-association2020_05_31",
	Short: "We recommend that you use the `UpdateDomainAssociation` API operation to move a domain association, as it supports both standard distributions and distribution tenants.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_updateDomainAssociation2020_05_31Cmd).Standalone()

	cloudfront_updateDomainAssociation2020_05_31Cmd.Flags().String("domain", "", "The domain to update.")
	cloudfront_updateDomainAssociation2020_05_31Cmd.Flags().String("if-match", "", "The value of the `ETag` identifier for the standard distribution or distribution tenant that will be associated with the domain.")
	cloudfront_updateDomainAssociation2020_05_31Cmd.Flags().String("target-resource", "", "The target standard distribution or distribution tenant resource for the domain.")
	cloudfront_updateDomainAssociation2020_05_31Cmd.MarkFlagRequired("domain")
	cloudfront_updateDomainAssociation2020_05_31Cmd.MarkFlagRequired("target-resource")
	cloudfrontCmd.AddCommand(cloudfront_updateDomainAssociation2020_05_31Cmd)
}
