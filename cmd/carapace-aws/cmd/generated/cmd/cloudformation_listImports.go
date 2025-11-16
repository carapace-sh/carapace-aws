package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_listImportsCmd = &cobra.Command{
	Use:   "list-imports",
	Short: "Lists all stacks that are importing an exported output value.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_listImportsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_listImportsCmd).Standalone()

		cloudformation_listImportsCmd.Flags().String("export-name", "", "The name of the exported output value.")
		cloudformation_listImportsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		cloudformation_listImportsCmd.MarkFlagRequired("export-name")
	})
	cloudformationCmd.AddCommand(cloudformation_listImportsCmd)
}
