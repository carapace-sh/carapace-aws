package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_listStackSetsCmd = &cobra.Command{
	Use:   "list-stack-sets",
	Short: "Returns summary information about StackSets that are associated with the user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_listStackSetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_listStackSetsCmd).Standalone()

		cloudformation_listStackSetsCmd.Flags().String("call-as", "", "\\[Service-managed permissions] Specifies whether you are acting as an account administrator in the management account or as a delegated administrator in a member account.")
		cloudformation_listStackSetsCmd.Flags().String("max-results", "", "The maximum number of results to be returned with a single call.")
		cloudformation_listStackSetsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		cloudformation_listStackSetsCmd.Flags().String("status", "", "The status of the StackSets that you want to get summary information about.")
	})
	cloudformationCmd.AddCommand(cloudformation_listStackSetsCmd)
}
