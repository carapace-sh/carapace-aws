package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getContinuousDeploymentPolicyConfig2020_05_31Cmd = &cobra.Command{
	Use:   "get-continuous-deployment-policy-config2020_05_31",
	Short: "Gets configuration information about a continuous deployment policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getContinuousDeploymentPolicyConfig2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_getContinuousDeploymentPolicyConfig2020_05_31Cmd).Standalone()

		cloudfront_getContinuousDeploymentPolicyConfig2020_05_31Cmd.Flags().String("id", "", "The identifier of the continuous deployment policy whose configuration you are getting.")
		cloudfront_getContinuousDeploymentPolicyConfig2020_05_31Cmd.MarkFlagRequired("id")
	})
	cloudfrontCmd.AddCommand(cloudfront_getContinuousDeploymentPolicyConfig2020_05_31Cmd)
}
