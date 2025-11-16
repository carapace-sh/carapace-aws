package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_updateDeploymentGroupCmd = &cobra.Command{
	Use:   "update-deployment-group",
	Short: "Changes information about a deployment group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_updateDeploymentGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codedeploy_updateDeploymentGroupCmd).Standalone()

		codedeploy_updateDeploymentGroupCmd.Flags().String("alarm-configuration", "", "Information to add or change about Amazon CloudWatch alarms when the deployment group is updated.")
		codedeploy_updateDeploymentGroupCmd.Flags().String("application-name", "", "The application name that corresponds to the deployment group to update.")
		codedeploy_updateDeploymentGroupCmd.Flags().String("auto-rollback-configuration", "", "Information for an automatic rollback configuration that is added or changed when a deployment group is updated.")
		codedeploy_updateDeploymentGroupCmd.Flags().String("auto-scaling-groups", "", "The replacement list of Auto Scaling groups to be included in the deployment group, if you want to change them.")
		codedeploy_updateDeploymentGroupCmd.Flags().String("blue-green-deployment-configuration", "", "Information about blue/green deployment options for a deployment group.")
		codedeploy_updateDeploymentGroupCmd.Flags().String("current-deployment-group-name", "", "The current name of the deployment group.")
		codedeploy_updateDeploymentGroupCmd.Flags().String("deployment-config-name", "", "The replacement deployment configuration name to use, if you want to change it.")
		codedeploy_updateDeploymentGroupCmd.Flags().String("deployment-style", "", "Information about the type of deployment, either in-place or blue/green, you want to run and whether to route deployment traffic behind a load balancer.")
		codedeploy_updateDeploymentGroupCmd.Flags().String("ec2-tag-filters", "", "The replacement set of Amazon EC2 tags on which to filter, if you want to change them.")
		codedeploy_updateDeploymentGroupCmd.Flags().String("ec2-tag-set", "", "Information about groups of tags applied to on-premises instances.")
		codedeploy_updateDeploymentGroupCmd.Flags().String("ecs-services", "", "The target Amazon ECS services in the deployment group.")
		codedeploy_updateDeploymentGroupCmd.Flags().String("load-balancer-info", "", "Information about the load balancer used in a deployment.")
		codedeploy_updateDeploymentGroupCmd.Flags().String("new-deployment-group-name", "", "The new name of the deployment group, if you want to change it.")
		codedeploy_updateDeploymentGroupCmd.Flags().String("on-premises-instance-tag-filters", "", "The replacement set of on-premises instance tags on which to filter, if you want to change them.")
		codedeploy_updateDeploymentGroupCmd.Flags().String("on-premises-tag-set", "", "Information about an on-premises instance tag set.")
		codedeploy_updateDeploymentGroupCmd.Flags().String("outdated-instances-strategy", "", "Indicates what happens when new Amazon EC2 instances are launched mid-deployment and do not receive the deployed application revision.")
		codedeploy_updateDeploymentGroupCmd.Flags().String("service-role-arn", "", "A replacement ARN for the service role, if you want to change it.")
		codedeploy_updateDeploymentGroupCmd.Flags().String("termination-hook-enabled", "", "This parameter only applies if you are using CodeDeploy with Amazon EC2 Auto Scaling.")
		codedeploy_updateDeploymentGroupCmd.Flags().String("trigger-configurations", "", "Information about triggers to change when the deployment group is updated.")
		codedeploy_updateDeploymentGroupCmd.MarkFlagRequired("application-name")
		codedeploy_updateDeploymentGroupCmd.MarkFlagRequired("current-deployment-group-name")
	})
	codedeployCmd.AddCommand(codedeploy_updateDeploymentGroupCmd)
}
