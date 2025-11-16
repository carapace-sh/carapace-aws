package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_deleteContinuousDeploymentPolicy2020_05_31Cmd = &cobra.Command{
	Use:   "delete-continuous-deployment-policy2020_05_31",
	Short: "Deletes a continuous deployment policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_deleteContinuousDeploymentPolicy2020_05_31Cmd).Standalone()

	cloudfront_deleteContinuousDeploymentPolicy2020_05_31Cmd.Flags().String("id", "", "The identifier of the continuous deployment policy that you are deleting.")
	cloudfront_deleteContinuousDeploymentPolicy2020_05_31Cmd.Flags().String("if-match", "", "The current version (`ETag` value) of the continuous deployment policy that you are deleting.")
	cloudfront_deleteContinuousDeploymentPolicy2020_05_31Cmd.MarkFlagRequired("id")
	cloudfrontCmd.AddCommand(cloudfront_deleteContinuousDeploymentPolicy2020_05_31Cmd)
}
