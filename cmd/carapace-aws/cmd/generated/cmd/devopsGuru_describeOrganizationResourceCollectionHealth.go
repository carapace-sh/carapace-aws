package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_describeOrganizationResourceCollectionHealthCmd = &cobra.Command{
	Use:   "describe-organization-resource-collection-health",
	Short: "Provides an overview of your system's health.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_describeOrganizationResourceCollectionHealthCmd).Standalone()

	devopsGuru_describeOrganizationResourceCollectionHealthCmd.Flags().String("account-ids", "", "The ID of the Amazon Web Services account.")
	devopsGuru_describeOrganizationResourceCollectionHealthCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	devopsGuru_describeOrganizationResourceCollectionHealthCmd.Flags().String("next-token", "", "The pagination token to use to retrieve the next page of results for this operation.")
	devopsGuru_describeOrganizationResourceCollectionHealthCmd.Flags().String("organization-resource-collection-type", "", "An Amazon Web Services resource collection type.")
	devopsGuru_describeOrganizationResourceCollectionHealthCmd.Flags().String("organizational-unit-ids", "", "The ID of the organizational unit.")
	devopsGuru_describeOrganizationResourceCollectionHealthCmd.MarkFlagRequired("organization-resource-collection-type")
	devopsGuruCmd.AddCommand(devopsGuru_describeOrganizationResourceCollectionHealthCmd)
}
