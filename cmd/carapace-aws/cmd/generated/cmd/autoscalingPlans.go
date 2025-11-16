package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscalingPlansCmd = &cobra.Command{
	Use:   "autoscaling-plans",
	Short: "AWS Auto Scaling",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscalingPlansCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscalingPlansCmd).Standalone()

	})
	rootCmd.AddCommand(autoscalingPlansCmd)
}
