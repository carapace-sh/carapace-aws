package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controltower_getBaselineOperationCmd = &cobra.Command{
	Use:   "get-baseline-operation",
	Short: "Returns the details of an asynchronous baseline operation, as initiated by any of these APIs: `EnableBaseline`, `DisableBaseline`, `UpdateEnabledBaseline`, `ResetEnabledBaseline`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controltower_getBaselineOperationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(controltower_getBaselineOperationCmd).Standalone()

		controltower_getBaselineOperationCmd.Flags().String("operation-identifier", "", "The operation ID returned from mutating asynchronous APIs (Enable, Disable, Update, Reset).")
		controltower_getBaselineOperationCmd.MarkFlagRequired("operation-identifier")
	})
	controltowerCmd.AddCommand(controltower_getBaselineOperationCmd)
}
