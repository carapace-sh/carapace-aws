package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_describeLaunchConfigurationsCmd = &cobra.Command{
	Use:   "describe-launch-configurations",
	Short: "Gets information about the launch configurations in the account and Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_describeLaunchConfigurationsCmd).Standalone()

	autoscaling_describeLaunchConfigurationsCmd.Flags().String("launch-configuration-names", "", "The launch configuration names.")
	autoscaling_describeLaunchConfigurationsCmd.Flags().String("max-records", "", "The maximum number of items to return with this call.")
	autoscaling_describeLaunchConfigurationsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	autoscalingCmd.AddCommand(autoscaling_describeLaunchConfigurationsCmd)
}
