package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_listCrossAccountResourcesCmd = &cobra.Command{
	Use:   "list-cross-account-resources",
	Short: "List the cross-account resources available to work with.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_listCrossAccountResourcesCmd).Standalone()

	globalaccelerator_listCrossAccountResourcesCmd.Flags().String("accelerator-arn", "", "The Amazon Resource Name (ARN) of an accelerator in a cross-account attachment.")
	globalaccelerator_listCrossAccountResourcesCmd.Flags().String("max-results", "", "The number of cross-account resource objects that you want to return with this call.")
	globalaccelerator_listCrossAccountResourcesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	globalaccelerator_listCrossAccountResourcesCmd.Flags().String("resource-owner-aws-account-id", "", "The account ID of a resource owner in a cross-account attachment.")
	globalaccelerator_listCrossAccountResourcesCmd.MarkFlagRequired("resource-owner-aws-account-id")
	globalacceleratorCmd.AddCommand(globalaccelerator_listCrossAccountResourcesCmd)
}
