package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_setInstanceProtectionCmd = &cobra.Command{
	Use:   "set-instance-protection",
	Short: "Updates the instance protection settings of the specified instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_setInstanceProtectionCmd).Standalone()

	autoscaling_setInstanceProtectionCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
	autoscaling_setInstanceProtectionCmd.Flags().String("instance-ids", "", "One or more instance IDs.")
	autoscaling_setInstanceProtectionCmd.Flags().String("protected-from-scale-in", "", "Indicates whether the instance is protected from termination by Amazon EC2 Auto Scaling when scaling in.")
	autoscaling_setInstanceProtectionCmd.MarkFlagRequired("auto-scaling-group-name")
	autoscaling_setInstanceProtectionCmd.MarkFlagRequired("instance-ids")
	autoscaling_setInstanceProtectionCmd.MarkFlagRequired("protected-from-scale-in")
	autoscalingCmd.AddCommand(autoscaling_setInstanceProtectionCmd)
}
