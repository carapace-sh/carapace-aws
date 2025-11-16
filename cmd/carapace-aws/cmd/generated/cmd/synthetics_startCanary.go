package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var synthetics_startCanaryCmd = &cobra.Command{
	Use:   "start-canary",
	Short: "Use this operation to run a canary that has already been created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(synthetics_startCanaryCmd).Standalone()

	synthetics_startCanaryCmd.Flags().String("name", "", "The name of the canary that you want to run.")
	synthetics_startCanaryCmd.MarkFlagRequired("name")
	syntheticsCmd.AddCommand(synthetics_startCanaryCmd)
}
