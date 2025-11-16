package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_createDeploymentGroupCmd = &cobra.Command{
	Use:   "create-deployment-group",
	Short: "Creates a deployment group to which application revisions are deployed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_createDeploymentGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codedeploy_createDeploymentGroupCmd).Standalone()

		codedeploy_createDeploymentGroupCmd.Flags().String("alarm-configuration", "", "Information to add about Amazon CloudWatch alarms when the deployment group is created.")
		codedeploy_createDeploymentGroupCmd.Flags().String("application-name", "", "The name of an CodeDeploy application associated with the user or Amazon Web Services account.")
		codedeploy_createDeploymentGroupCmd.Flags().String("auto-rollback-configuration", "", "Configuration information for an automatic rollback that is added when a deployment group is created.")
		codedeploy_createDeploymentGroupCmd.Flags().String("auto-scaling-groups", "", "A list of associated Amazon EC2 Auto Scaling groups.")
		codedeploy_createDeploymentGroupCmd.Flags().String("blue-green-deployment-configuration", "", "Information about blue/green deployment options for a deployment group.")
		codedeploy_createDeploymentGroupCmd.Flags().String("deployment-config-name", "", "If specified, the deployment configuration name can be either one of the predefined configurations provided with CodeDeploy or a custom deployment configuration that you create by calling the create deployment configuration operation.")
		codedeploy_createDeploymentGroupCmd.Flags().String("deployment-group-name", "", "The name of a new deployment group for the specified application.")
		codedeploy_createDeploymentGroupCmd.Flags().String("deployment-style", "", "Information about the type of deployment, in-place or blue/green, that you want to run and whether to route deployment traffic behind a load balancer.")
		codedeploy_createDeploymentGroupCmd.Flags().String("ec2-tag-filters", "", "The Amazon EC2 tags on which to filter.")
		codedeploy_createDeploymentGroupCmd.Flags().String("ec2-tag-set", "", "Information about groups of tags applied to Amazon EC2 instances.")
		codedeploy_createDeploymentGroupCmd.Flags().String("ecs-services", "", "The target Amazon ECS services in the deployment group.")
		codedeploy_createDeploymentGroupCmd.Flags().String("load-balancer-info", "", "Information about the load balancer used in a deployment.")
		codedeploy_createDeploymentGroupCmd.Flags().String("on-premises-instance-tag-filters", "", "The on-premises instance tags on which to filter.")
		codedeploy_createDeploymentGroupCmd.Flags().String("on-premises-tag-set", "", "Information about groups of tags applied to on-premises instances.")
		codedeploy_createDeploymentGroupCmd.Flags().String("outdated-instances-strategy", "", "Indicates what happens when new Amazon EC2 instances are launched mid-deployment and do not receive the deployed application revision.")
		codedeploy_createDeploymentGroupCmd.Flags().String("service-role-arn", "", "A service role Amazon Resource Name (ARN) that allows CodeDeploy to act on the user's behalf when interacting with Amazon Web Services services.")
		codedeploy_createDeploymentGroupCmd.Flags().String("tags", "", "The metadata that you apply to CodeDeploy deployment groups to help you organize and categorize them.")
		codedeploy_createDeploymentGroupCmd.Flags().String("termination-hook-enabled", "", "This parameter only applies if you are using CodeDeploy with Amazon EC2 Auto Scaling.")
		codedeploy_createDeploymentGroupCmd.Flags().String("trigger-configurations", "", "Information about triggers to create when the deployment group is created.")
		codedeploy_createDeploymentGroupCmd.MarkFlagRequired("application-name")
		codedeploy_createDeploymentGroupCmd.MarkFlagRequired("deployment-group-name")
		codedeploy_createDeploymentGroupCmd.MarkFlagRequired("service-role-arn")
	})
	codedeployCmd.AddCommand(codedeploy_createDeploymentGroupCmd)
}
