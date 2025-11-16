package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscalingPlansCmd = &cobra.Command{
	Use:   "autoscaling-plans",
	Short: "AWS Auto Scaling\n\nUse AWS Auto Scaling to create scaling plans for your applications to automatically scale your scalable AWS resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscalingPlansCmd).Standalone()

	rootCmd.AddCommand(autoscalingPlansCmd)
}
