package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_listCidrBlocksCmd = &cobra.Command{
	Use:   "list-cidr-blocks",
	Short: "Returns a paginated list of location objects and their CIDR blocks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_listCidrBlocksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_listCidrBlocksCmd).Standalone()

		route53_listCidrBlocksCmd.Flags().String("collection-id", "", "The UUID of the CIDR collection.")
		route53_listCidrBlocksCmd.Flags().String("location-name", "", "The name of the CIDR collection location.")
		route53_listCidrBlocksCmd.Flags().String("max-results", "", "Maximum number of results you want returned.")
		route53_listCidrBlocksCmd.Flags().String("next-token", "", "An opaque pagination token to indicate where the service is to begin enumerating results.")
		route53_listCidrBlocksCmd.MarkFlagRequired("collection-id")
	})
	route53Cmd.AddCommand(route53_listCidrBlocksCmd)
}
