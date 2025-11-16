package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscalingCmd = &cobra.Command{
	Use:   "autoscaling",
	Short: "Amazon EC2 Auto Scaling",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscalingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscalingCmd).Standalone()

	})
	rootCmd.AddCommand(autoscalingCmd)
}
