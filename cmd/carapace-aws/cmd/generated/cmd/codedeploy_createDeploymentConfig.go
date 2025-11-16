package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_createDeploymentConfigCmd = &cobra.Command{
	Use:   "create-deployment-config",
	Short: "Creates a deployment configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_createDeploymentConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codedeploy_createDeploymentConfigCmd).Standalone()

		codedeploy_createDeploymentConfigCmd.Flags().String("compute-platform", "", "The destination platform type for the deployment (`Lambda`, `Server`, or `ECS`).")
		codedeploy_createDeploymentConfigCmd.Flags().String("deployment-config-name", "", "The name of the deployment configuration to create.")
		codedeploy_createDeploymentConfigCmd.Flags().String("minimum-healthy-hosts", "", "The minimum number of healthy instances that should be available at any time during the deployment.")
		codedeploy_createDeploymentConfigCmd.Flags().String("traffic-routing-config", "", "The configuration that specifies how the deployment traffic is routed.")
		codedeploy_createDeploymentConfigCmd.Flags().String("zonal-config", "", "Configure the `ZonalConfig` object if you want CodeDeploy to deploy your application to one [Availability Zone](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-availability-zones) at a time, within an Amazon Web Services Region.")
		codedeploy_createDeploymentConfigCmd.MarkFlagRequired("deployment-config-name")
	})
	codedeployCmd.AddCommand(codedeploy_createDeploymentConfigCmd)
}
