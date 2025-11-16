package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_updateAutoScalingGroupCmd = &cobra.Command{
	Use:   "update-auto-scaling-group",
	Short: "**We strongly recommend that all Auto Scaling groups use launch templates to ensure full functionality for Amazon EC2 Auto Scaling and Amazon EC2.**\n\nUpdates the configuration for the specified Auto Scaling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_updateAutoScalingGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_updateAutoScalingGroupCmd).Standalone()

		autoscaling_updateAutoScalingGroupCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
		autoscaling_updateAutoScalingGroupCmd.Flags().String("availability-zone-distribution", "", "The instance capacity distribution across Availability Zones.")
		autoscaling_updateAutoScalingGroupCmd.Flags().String("availability-zone-impairment-policy", "", "The policy for Availability Zone impairment.")
		autoscaling_updateAutoScalingGroupCmd.Flags().String("availability-zones", "", "One or more Availability Zones for the group.")
		autoscaling_updateAutoScalingGroupCmd.Flags().String("capacity-rebalance", "", "Enables or disables Capacity Rebalancing.")
		autoscaling_updateAutoScalingGroupCmd.Flags().String("capacity-reservation-specification", "", "The capacity reservation specification for the Auto Scaling group.")
		autoscaling_updateAutoScalingGroupCmd.Flags().String("context", "", "Reserved.")
		autoscaling_updateAutoScalingGroupCmd.Flags().String("default-cooldown", "", "*Only needed if you use simple scaling policies.*")
		autoscaling_updateAutoScalingGroupCmd.Flags().String("default-instance-warmup", "", "The amount of time, in seconds, until a new instance is considered to have finished initializing and resource consumption to become stable after it enters the `InService` state.")
		autoscaling_updateAutoScalingGroupCmd.Flags().String("desired-capacity", "", "The desired capacity is the initial capacity of the Auto Scaling group after this operation completes and the capacity it attempts to maintain.")
		autoscaling_updateAutoScalingGroupCmd.Flags().String("desired-capacity-type", "", "The unit of measurement for the value specified for desired capacity.")
		autoscaling_updateAutoScalingGroupCmd.Flags().String("health-check-grace-period", "", "The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before checking the health status of an EC2 instance that has come into service and marking it unhealthy due to a failed health check.")
		autoscaling_updateAutoScalingGroupCmd.Flags().String("health-check-type", "", "A comma-separated value string of one or more health check types.")
		autoscaling_updateAutoScalingGroupCmd.Flags().String("instance-maintenance-policy", "", "An instance maintenance policy.")
		autoscaling_updateAutoScalingGroupCmd.Flags().String("launch-configuration-name", "", "The name of the launch configuration.")
		autoscaling_updateAutoScalingGroupCmd.Flags().String("launch-template", "", "The launch template and version to use to specify the updates.")
		autoscaling_updateAutoScalingGroupCmd.Flags().String("max-instance-lifetime", "", "The maximum amount of time, in seconds, that an instance can be in service.")
		autoscaling_updateAutoScalingGroupCmd.Flags().String("max-size", "", "The maximum size of the Auto Scaling group.")
		autoscaling_updateAutoScalingGroupCmd.Flags().String("min-size", "", "The minimum size of the Auto Scaling group.")
		autoscaling_updateAutoScalingGroupCmd.Flags().String("mixed-instances-policy", "", "The mixed instances policy.")
		autoscaling_updateAutoScalingGroupCmd.Flags().String("new-instances-protected-from-scale-in", "", "Indicates whether newly launched instances are protected from termination by Amazon EC2 Auto Scaling when scaling in.")
		autoscaling_updateAutoScalingGroupCmd.Flags().String("placement-group", "", "The name of an existing placement group into which to launch your instances.")
		autoscaling_updateAutoScalingGroupCmd.Flags().String("service-linked-role-arn", "", "The Amazon Resource Name (ARN) of the service-linked role that the Auto Scaling group uses to call other Amazon Web Services on your behalf.")
		autoscaling_updateAutoScalingGroupCmd.Flags().String("skip-zonal-shift-validation", "", "If you enable zonal shift with cross-zone disabled load balancers, capacity could become imbalanced across Availability Zones.")
		autoscaling_updateAutoScalingGroupCmd.Flags().String("termination-policies", "", "A policy or a list of policies that are used to select the instances to terminate.")
		autoscaling_updateAutoScalingGroupCmd.Flags().String("vpczone-identifier", "", "A comma-separated list of subnet IDs for a virtual private cloud (VPC).")
		autoscaling_updateAutoScalingGroupCmd.MarkFlagRequired("auto-scaling-group-name")
	})
	autoscalingCmd.AddCommand(autoscaling_updateAutoScalingGroupCmd)
}
