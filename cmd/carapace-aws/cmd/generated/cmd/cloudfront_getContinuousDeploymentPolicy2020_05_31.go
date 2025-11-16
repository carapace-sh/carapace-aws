package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getContinuousDeploymentPolicy2020_05_31Cmd = &cobra.Command{
	Use:   "get-continuous-deployment-policy2020_05_31",
	Short: "Gets a continuous deployment policy, including metadata (the policy's identifier and the date and time when the policy was last modified).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getContinuousDeploymentPolicy2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_getContinuousDeploymentPolicy2020_05_31Cmd).Standalone()

		cloudfront_getContinuousDeploymentPolicy2020_05_31Cmd.Flags().String("id", "", "The identifier of the continuous deployment policy that you are getting.")
		cloudfront_getContinuousDeploymentPolicy2020_05_31Cmd.MarkFlagRequired("id")
	})
	cloudfrontCmd.AddCommand(cloudfront_getContinuousDeploymentPolicy2020_05_31Cmd)
}
