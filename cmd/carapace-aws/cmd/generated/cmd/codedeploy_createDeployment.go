package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_createDeploymentCmd = &cobra.Command{
	Use:   "create-deployment",
	Short: "Deploys an application revision through the specified deployment group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_createDeploymentCmd).Standalone()

	codedeploy_createDeploymentCmd.Flags().String("application-name", "", "The name of an CodeDeploy application associated with the user or Amazon Web Services account.")
	codedeploy_createDeploymentCmd.Flags().String("auto-rollback-configuration", "", "Configuration information for an automatic rollback that is added when a deployment is created.")
	codedeploy_createDeploymentCmd.Flags().String("deployment-config-name", "", "The name of a deployment configuration associated with the user or Amazon Web Services account.")
	codedeploy_createDeploymentCmd.Flags().String("deployment-group-name", "", "The name of the deployment group.")
	codedeploy_createDeploymentCmd.Flags().String("description", "", "A comment about the deployment.")
	codedeploy_createDeploymentCmd.Flags().String("file-exists-behavior", "", "Information about how CodeDeploy handles files that already exist in a deployment target location but weren't part of the previous successful deployment.")
	codedeploy_createDeploymentCmd.Flags().Bool("ignore-application-stop-failures", false, "If true, then if an `ApplicationStop`, `BeforeBlockTraffic`, or `AfterBlockTraffic` deployment lifecycle event to an instance fails, then the deployment continues to the next deployment lifecycle event.")
	codedeploy_createDeploymentCmd.Flags().Bool("no-ignore-application-stop-failures", false, "If true, then if an `ApplicationStop`, `BeforeBlockTraffic`, or `AfterBlockTraffic` deployment lifecycle event to an instance fails, then the deployment continues to the next deployment lifecycle event.")
	codedeploy_createDeploymentCmd.Flags().Bool("no-update-outdated-instances-only", false, "Indicates whether to deploy to all instances or only to instances that are not running the latest application revision.")
	codedeploy_createDeploymentCmd.Flags().String("override-alarm-configuration", "", "Allows you to specify information about alarms associated with a deployment.")
	codedeploy_createDeploymentCmd.Flags().String("revision", "", "The type and location of the revision to deploy.")
	codedeploy_createDeploymentCmd.Flags().String("target-instances", "", "Information about the instances that belong to the replacement environment in a blue/green deployment.")
	codedeploy_createDeploymentCmd.Flags().Bool("update-outdated-instances-only", false, "Indicates whether to deploy to all instances or only to instances that are not running the latest application revision.")
	codedeploy_createDeploymentCmd.MarkFlagRequired("application-name")
	codedeploy_createDeploymentCmd.Flag("no-ignore-application-stop-failures").Hidden = true
	codedeploy_createDeploymentCmd.Flag("no-update-outdated-instances-only").Hidden = true
	codedeployCmd.AddCommand(codedeploy_createDeploymentCmd)
}
