package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationAutoscalingCmd = &cobra.Command{
	Use:   "application-autoscaling",
	Short: "With Application Auto Scaling, you can configure automatic scaling for the following resources:",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationAutoscalingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationAutoscalingCmd).Standalone()

	})
	rootCmd.AddCommand(applicationAutoscalingCmd)
}
