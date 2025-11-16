package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var synthetics_stopCanaryCmd = &cobra.Command{
	Use:   "stop-canary",
	Short: "Stops the canary to prevent all future runs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(synthetics_stopCanaryCmd).Standalone()

	synthetics_stopCanaryCmd.Flags().String("name", "", "The name of the canary that you want to stop.")
	synthetics_stopCanaryCmd.MarkFlagRequired("name")
	syntheticsCmd.AddCommand(synthetics_stopCanaryCmd)
}
