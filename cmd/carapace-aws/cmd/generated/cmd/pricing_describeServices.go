package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pricing_describeServicesCmd = &cobra.Command{
	Use:   "describe-services",
	Short: "Returns the metadata for one service or a list of the metadata for all services.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pricing_describeServicesCmd).Standalone()

	pricing_describeServicesCmd.Flags().String("format-version", "", "The format version that you want the response to be in.")
	pricing_describeServicesCmd.Flags().String("max-results", "", "The maximum number of results that you want returned in the response.")
	pricing_describeServicesCmd.Flags().String("next-token", "", "The pagination token that indicates the next set of results that you want to retrieve.")
	pricing_describeServicesCmd.Flags().String("service-code", "", "The code for the service whose information you want to retrieve, such as `AmazonEC2`.")
	pricingCmd.AddCommand(pricing_describeServicesCmd)
}
