package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_attachInstancesCmd = &cobra.Command{
	Use:   "attach-instances",
	Short: "Attaches one or more EC2 instances to the specified Auto Scaling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_attachInstancesCmd).Standalone()

	autoscaling_attachInstancesCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
	autoscaling_attachInstancesCmd.Flags().String("instance-ids", "", "The IDs of the instances.")
	autoscaling_attachInstancesCmd.MarkFlagRequired("auto-scaling-group-name")
	autoscalingCmd.AddCommand(autoscaling_attachInstancesCmd)
}
