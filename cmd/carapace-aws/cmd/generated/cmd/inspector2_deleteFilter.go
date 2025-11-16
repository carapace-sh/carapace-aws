package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_deleteFilterCmd = &cobra.Command{
	Use:   "delete-filter",
	Short: "Deletes a filter resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_deleteFilterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_deleteFilterCmd).Standalone()

		inspector2_deleteFilterCmd.Flags().String("arn", "", "The Amazon Resource Number (ARN) of the filter to be deleted.")
		inspector2_deleteFilterCmd.MarkFlagRequired("arn")
	})
	inspector2Cmd.AddCommand(inspector2_deleteFilterCmd)
}
