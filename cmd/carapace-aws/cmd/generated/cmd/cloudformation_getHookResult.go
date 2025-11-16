package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_getHookResultCmd = &cobra.Command{
	Use:   "get-hook-result",
	Short: "Retrieves detailed information and remediation guidance for a Hook invocation result.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_getHookResultCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_getHookResultCmd).Standalone()

		cloudformation_getHookResultCmd.Flags().String("hook-result-id", "", "The unique identifier (ID) of the Hook invocation result that you want details about.")
	})
	cloudformationCmd.AddCommand(cloudformation_getHookResultCmd)
}
