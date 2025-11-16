package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_listTypesCmd = &cobra.Command{
	Use:   "list-types",
	Short: "Returns summary information about all extensions, including your private resource types, modules, and Hooks as well as all public extensions from Amazon Web Services and third-party publishers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_listTypesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_listTypesCmd).Standalone()

		cloudformation_listTypesCmd.Flags().String("deprecated-status", "", "The deprecation status of the extension that you want to get summary information about.")
		cloudformation_listTypesCmd.Flags().String("filters", "", "Filter criteria to use in determining which extensions to return.")
		cloudformation_listTypesCmd.Flags().String("max-results", "", "The maximum number of results to be returned with a single call.")
		cloudformation_listTypesCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		cloudformation_listTypesCmd.Flags().String("provisioning-type", "", "For resource types, the provisioning behavior of the resource type.")
		cloudformation_listTypesCmd.Flags().String("type", "", "The type of extension.")
		cloudformation_listTypesCmd.Flags().String("visibility", "", "The scope at which the extensions are visible and usable in CloudFormation operations.")
	})
	cloudformationCmd.AddCommand(cloudformation_listTypesCmd)
}
