package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_listCidrCollectionsCmd = &cobra.Command{
	Use:   "list-cidr-collections",
	Short: "Returns a paginated list of CIDR collections in the Amazon Web Services account (metadata only).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_listCidrCollectionsCmd).Standalone()

	route53_listCidrCollectionsCmd.Flags().String("max-results", "", "The maximum number of CIDR collections to return in the response.")
	route53_listCidrCollectionsCmd.Flags().String("next-token", "", "An opaque pagination token to indicate where the service is to begin enumerating results.")
	route53Cmd.AddCommand(route53_listCidrCollectionsCmd)
}
