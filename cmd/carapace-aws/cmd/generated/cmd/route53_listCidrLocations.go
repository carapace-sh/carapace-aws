package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_listCidrLocationsCmd = &cobra.Command{
	Use:   "list-cidr-locations",
	Short: "Returns a paginated list of CIDR locations for the given collection (metadata only, does not include CIDR blocks).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_listCidrLocationsCmd).Standalone()

	route53_listCidrLocationsCmd.Flags().String("collection-id", "", "The CIDR collection ID.")
	route53_listCidrLocationsCmd.Flags().String("max-results", "", "The maximum number of CIDR collection locations to return in the response.")
	route53_listCidrLocationsCmd.Flags().String("next-token", "", "An opaque pagination token to indicate where the service is to begin enumerating results.")
	route53_listCidrLocationsCmd.MarkFlagRequired("collection-id")
	route53Cmd.AddCommand(route53_listCidrLocationsCmd)
}
