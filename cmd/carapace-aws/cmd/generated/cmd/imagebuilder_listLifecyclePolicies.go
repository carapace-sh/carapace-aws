package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_listLifecyclePoliciesCmd = &cobra.Command{
	Use:   "list-lifecycle-policies",
	Short: "Get a list of lifecycle policies in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_listLifecyclePoliciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_listLifecyclePoliciesCmd).Standalone()

		imagebuilder_listLifecyclePoliciesCmd.Flags().String("filters", "", "Streamline results based on one of the following values: `Name`, `Status`.")
		imagebuilder_listLifecyclePoliciesCmd.Flags().String("max-results", "", "Specify the maximum number of items to return in a request.")
		imagebuilder_listLifecyclePoliciesCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
	})
	imagebuilderCmd.AddCommand(imagebuilder_listLifecyclePoliciesCmd)
}
