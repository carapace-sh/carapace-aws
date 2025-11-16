package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_deleteOriginRequestPolicy2020_05_31Cmd = &cobra.Command{
	Use:   "delete-origin-request-policy2020_05_31",
	Short: "Deletes an origin request policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_deleteOriginRequestPolicy2020_05_31Cmd).Standalone()

	cloudfront_deleteOriginRequestPolicy2020_05_31Cmd.Flags().String("id", "", "The unique identifier for the origin request policy that you are deleting.")
	cloudfront_deleteOriginRequestPolicy2020_05_31Cmd.Flags().String("if-match", "", "The version of the origin request policy that you are deleting.")
	cloudfront_deleteOriginRequestPolicy2020_05_31Cmd.MarkFlagRequired("id")
	cloudfrontCmd.AddCommand(cloudfront_deleteOriginRequestPolicy2020_05_31Cmd)
}
