package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_deleteResponseHeadersPolicy2020_05_31Cmd = &cobra.Command{
	Use:   "delete-response-headers-policy2020_05_31",
	Short: "Deletes a response headers policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_deleteResponseHeadersPolicy2020_05_31Cmd).Standalone()

	cloudfront_deleteResponseHeadersPolicy2020_05_31Cmd.Flags().String("id", "", "The identifier for the response headers policy that you are deleting.")
	cloudfront_deleteResponseHeadersPolicy2020_05_31Cmd.Flags().String("if-match", "", "The version of the response headers policy that you are deleting.")
	cloudfront_deleteResponseHeadersPolicy2020_05_31Cmd.MarkFlagRequired("id")
	cloudfrontCmd.AddCommand(cloudfront_deleteResponseHeadersPolicy2020_05_31Cmd)
}
