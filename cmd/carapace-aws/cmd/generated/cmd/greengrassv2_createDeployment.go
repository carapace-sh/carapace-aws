package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2_createDeploymentCmd = &cobra.Command{
	Use:   "create-deployment",
	Short: "Creates a continuous deployment for a target, which is a Greengrass core device or group of core devices.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2_createDeploymentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrassv2_createDeploymentCmd).Standalone()

		greengrassv2_createDeploymentCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you can provide to ensure that the request is idempotent.")
		greengrassv2_createDeploymentCmd.Flags().String("components", "", "The components to deploy.")
		greengrassv2_createDeploymentCmd.Flags().String("deployment-name", "", "The name of the deployment.")
		greengrassv2_createDeploymentCmd.Flags().String("deployment-policies", "", "The deployment policies for the deployment.")
		greengrassv2_createDeploymentCmd.Flags().String("iot-job-configuration", "", "The job configuration for the deployment configuration.")
		greengrassv2_createDeploymentCmd.Flags().String("parent-target-arn", "", "The parent deployment's target [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) within a subdeployment.")
		greengrassv2_createDeploymentCmd.Flags().String("tags", "", "A list of key-value pairs that contain metadata for the resource.")
		greengrassv2_createDeploymentCmd.Flags().String("target-arn", "", "The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the target IoT thing or thing group.")
		greengrassv2_createDeploymentCmd.MarkFlagRequired("target-arn")
	})
	greengrassv2Cmd.AddCommand(greengrassv2_createDeploymentCmd)
}
