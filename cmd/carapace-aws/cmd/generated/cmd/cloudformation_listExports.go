package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_listExportsCmd = &cobra.Command{
	Use:   "list-exports",
	Short: "Lists all exported output values in the account and Region in which you call this action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_listExportsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_listExportsCmd).Standalone()

		cloudformation_listExportsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	})
	cloudformationCmd.AddCommand(cloudformation_listExportsCmd)
}
