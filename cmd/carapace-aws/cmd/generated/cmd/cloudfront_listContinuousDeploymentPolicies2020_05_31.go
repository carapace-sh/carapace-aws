package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listContinuousDeploymentPolicies2020_05_31Cmd = &cobra.Command{
	Use:   "list-continuous-deployment-policies2020_05_31",
	Short: "Gets a list of the continuous deployment policies in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listContinuousDeploymentPolicies2020_05_31Cmd).Standalone()

	cloudfront_listContinuousDeploymentPolicies2020_05_31Cmd.Flags().String("marker", "", "Use this field when paginating results to indicate where to begin in your list of continuous deployment policies.")
	cloudfront_listContinuousDeploymentPolicies2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of continuous deployment policies that you want returned in the response.")
	cloudfrontCmd.AddCommand(cloudfront_listContinuousDeploymentPolicies2020_05_31Cmd)
}
