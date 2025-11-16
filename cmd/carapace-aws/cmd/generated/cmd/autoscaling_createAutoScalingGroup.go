package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_createAutoScalingGroupCmd = &cobra.Command{
	Use:   "create-auto-scaling-group",
	Short: "**We strongly recommend using a launch template when calling this operation to ensure full functionality for Amazon EC2 Auto Scaling and Amazon EC2.**\n\nCreates an Auto Scaling group with the specified name and attributes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_createAutoScalingGroupCmd).Standalone()

	autoscaling_createAutoScalingGroupCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("availability-zone-distribution", "", "The instance capacity distribution across Availability Zones.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("availability-zone-impairment-policy", "", "The policy for Availability Zone impairment.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("availability-zones", "", "A list of Availability Zones where instances in the Auto Scaling group can be created.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("capacity-rebalance", "", "Indicates whether Capacity Rebalancing is enabled.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("capacity-reservation-specification", "", "The capacity reservation specification for the Auto Scaling group.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("context", "", "Reserved.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("default-cooldown", "", "*Only needed if you use simple scaling policies.*")
	autoscaling_createAutoScalingGroupCmd.Flags().String("default-instance-warmup", "", "The amount of time, in seconds, until a new instance is considered to have finished initializing and resource consumption to become stable after it enters the `InService` state.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("desired-capacity", "", "The desired capacity is the initial capacity of the Auto Scaling group at the time of its creation and the capacity it attempts to maintain.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("desired-capacity-type", "", "The unit of measurement for the value specified for desired capacity.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("health-check-grace-period", "", "The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before checking the health status of an EC2 instance that has come into service and marking it unhealthy due to a failed health check.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("health-check-type", "", "A comma-separated value string of one or more health check types.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("instance-id", "", "The ID of the instance used to base the launch configuration on.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("instance-maintenance-policy", "", "An instance maintenance policy.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("launch-configuration-name", "", "The name of the launch configuration to use to launch instances.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("launch-template", "", "Information used to specify the launch template and version to use to launch instances.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("lifecycle-hook-specification-list", "", "One or more lifecycle hooks to add to the Auto Scaling group before instances are launched.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("load-balancer-names", "", "A list of Classic Load Balancers associated with this Auto Scaling group.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("max-instance-lifetime", "", "The maximum amount of time, in seconds, that an instance can be in service.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("max-size", "", "The maximum size of the group.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("min-size", "", "The minimum size of the group.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("mixed-instances-policy", "", "The mixed instances policy.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("new-instances-protected-from-scale-in", "", "Indicates whether newly launched instances are protected from termination by Amazon EC2 Auto Scaling when scaling in.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("placement-group", "", "The name of the placement group into which to launch your instances.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("service-linked-role-arn", "", "The Amazon Resource Name (ARN) of the service-linked role that the Auto Scaling group uses to call other Amazon Web Services service on your behalf.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("skip-zonal-shift-validation", "", "If you enable zonal shift with cross-zone disabled load balancers, capacity could become imbalanced across Availability Zones.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("tags", "", "One or more tags.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("target-group-arns", "", "The Amazon Resource Names (ARN) of the Elastic Load Balancing target groups to associate with the Auto Scaling group.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("termination-policies", "", "A policy or a list of policies that are used to select the instance to terminate.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("traffic-sources", "", "The list of traffic sources to attach to this Auto Scaling group.")
	autoscaling_createAutoScalingGroupCmd.Flags().String("vpczone-identifier", "", "A comma-separated list of subnet IDs for a virtual private cloud (VPC) where instances in the Auto Scaling group can be created.")
	autoscaling_createAutoScalingGroupCmd.MarkFlagRequired("auto-scaling-group-name")
	autoscaling_createAutoScalingGroupCmd.MarkFlagRequired("max-size")
	autoscaling_createAutoScalingGroupCmd.MarkFlagRequired("min-size")
	autoscalingCmd.AddCommand(autoscaling_createAutoScalingGroupCmd)
}
