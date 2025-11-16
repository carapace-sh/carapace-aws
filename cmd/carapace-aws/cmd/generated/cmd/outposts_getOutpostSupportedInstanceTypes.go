package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_getOutpostSupportedInstanceTypesCmd = &cobra.Command{
	Use:   "get-outpost-supported-instance-types",
	Short: "Gets the instance types that an Outpost can support in `InstanceTypeCapacity`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_getOutpostSupportedInstanceTypesCmd).Standalone()

	outposts_getOutpostSupportedInstanceTypesCmd.Flags().String("asset-id", "", "The ID of the Outpost asset.")
	outposts_getOutpostSupportedInstanceTypesCmd.Flags().String("max-results", "", "")
	outposts_getOutpostSupportedInstanceTypesCmd.Flags().String("next-token", "", "")
	outposts_getOutpostSupportedInstanceTypesCmd.Flags().String("order-id", "", "The ID for the Amazon Web Services Outposts order.")
	outposts_getOutpostSupportedInstanceTypesCmd.Flags().String("outpost-identifier", "", "The ID or ARN of the Outpost.")
	outposts_getOutpostSupportedInstanceTypesCmd.MarkFlagRequired("outpost-identifier")
	outpostsCmd.AddCommand(outposts_getOutpostSupportedInstanceTypesCmd)
}
