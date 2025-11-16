package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var synthetics_getCanaryCmd = &cobra.Command{
	Use:   "get-canary",
	Short: "Retrieves complete information about one canary.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(synthetics_getCanaryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(synthetics_getCanaryCmd).Standalone()

		synthetics_getCanaryCmd.Flags().String("dry-run-id", "", "The DryRunId associated with an existing canary’s dry run.")
		synthetics_getCanaryCmd.Flags().String("name", "", "The name of the canary that you want details for.")
		synthetics_getCanaryCmd.MarkFlagRequired("name")
	})
	syntheticsCmd.AddCommand(synthetics_getCanaryCmd)
}
