package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_updateContinuousDeploymentPolicy2020_05_31Cmd = &cobra.Command{
	Use:   "update-continuous-deployment-policy2020_05_31",
	Short: "Updates a continuous deployment policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_updateContinuousDeploymentPolicy2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_updateContinuousDeploymentPolicy2020_05_31Cmd).Standalone()

		cloudfront_updateContinuousDeploymentPolicy2020_05_31Cmd.Flags().String("continuous-deployment-policy-config", "", "The continuous deployment policy configuration.")
		cloudfront_updateContinuousDeploymentPolicy2020_05_31Cmd.Flags().String("id", "", "The identifier of the continuous deployment policy that you are updating.")
		cloudfront_updateContinuousDeploymentPolicy2020_05_31Cmd.Flags().String("if-match", "", "The current version (`ETag` value) of the continuous deployment policy that you are updating.")
		cloudfront_updateContinuousDeploymentPolicy2020_05_31Cmd.MarkFlagRequired("continuous-deployment-policy-config")
		cloudfront_updateContinuousDeploymentPolicy2020_05_31Cmd.MarkFlagRequired("id")
	})
	cloudfrontCmd.AddCommand(cloudfront_updateContinuousDeploymentPolicy2020_05_31Cmd)
}
