package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_createContinuousDeploymentPolicy2020_05_31Cmd = &cobra.Command{
	Use:   "create-continuous-deployment-policy2020_05_31",
	Short: "Creates a continuous deployment policy that distributes traffic for a custom domain name to two different CloudFront distributions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_createContinuousDeploymentPolicy2020_05_31Cmd).Standalone()

	cloudfront_createContinuousDeploymentPolicy2020_05_31Cmd.Flags().String("continuous-deployment-policy-config", "", "Contains the configuration for a continuous deployment policy.")
	cloudfront_createContinuousDeploymentPolicy2020_05_31Cmd.MarkFlagRequired("continuous-deployment-policy-config")
	cloudfrontCmd.AddCommand(cloudfront_createContinuousDeploymentPolicy2020_05_31Cmd)
}
