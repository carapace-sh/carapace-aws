package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_describeActionCmd = &cobra.Command{
	Use:   "describe-action",
	Short: "Retrieves information about an action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_describeActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_describeActionCmd).Standalone()

		iotsitewise_describeActionCmd.Flags().String("action-id", "", "The ID of the action.")
		iotsitewise_describeActionCmd.MarkFlagRequired("action-id")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_describeActionCmd)
}
