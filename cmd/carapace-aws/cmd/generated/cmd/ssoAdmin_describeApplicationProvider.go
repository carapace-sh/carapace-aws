package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_describeApplicationProviderCmd = &cobra.Command{
	Use:   "describe-application-provider",
	Short: "Retrieves details about a provider that can be used to connect an Amazon Web Services managed application or customer managed application to IAM Identity Center.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_describeApplicationProviderCmd).Standalone()

	ssoAdmin_describeApplicationProviderCmd.Flags().String("application-provider-arn", "", "Specifies the ARN of the application provider for which you want details.")
	ssoAdmin_describeApplicationProviderCmd.MarkFlagRequired("application-provider-arn")
	ssoAdminCmd.AddCommand(ssoAdmin_describeApplicationProviderCmd)
}
