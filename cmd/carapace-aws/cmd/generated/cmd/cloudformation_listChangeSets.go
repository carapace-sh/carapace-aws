package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_listChangeSetsCmd = &cobra.Command{
	Use:   "list-change-sets",
	Short: "Returns the ID and status of each active change set for a stack.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_listChangeSetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_listChangeSetsCmd).Standalone()

		cloudformation_listChangeSetsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		cloudformation_listChangeSetsCmd.Flags().String("stack-name", "", "The name or the Amazon Resource Name (ARN) of the stack for which you want to list change sets.")
		cloudformation_listChangeSetsCmd.MarkFlagRequired("stack-name")
	})
	cloudformationCmd.AddCommand(cloudformation_listChangeSetsCmd)
}
