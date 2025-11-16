package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_deleteLaunchConfigurationCmd = &cobra.Command{
	Use:   "delete-launch-configuration",
	Short: "Deletes the specified launch configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_deleteLaunchConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_deleteLaunchConfigurationCmd).Standalone()

		autoscaling_deleteLaunchConfigurationCmd.Flags().String("launch-configuration-name", "", "The name of the launch configuration.")
		autoscaling_deleteLaunchConfigurationCmd.MarkFlagRequired("launch-configuration-name")
	})
	autoscalingCmd.AddCommand(autoscaling_deleteLaunchConfigurationCmd)
}
