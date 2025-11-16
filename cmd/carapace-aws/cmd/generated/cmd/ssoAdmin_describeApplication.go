package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_describeApplicationCmd = &cobra.Command{
	Use:   "describe-application",
	Short: "Retrieves the details of an application associated with an instance of IAM Identity Center.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_describeApplicationCmd).Standalone()

	ssoAdmin_describeApplicationCmd.Flags().String("application-arn", "", "Specifies the ARN of the application.")
	ssoAdmin_describeApplicationCmd.MarkFlagRequired("application-arn")
	ssoAdminCmd.AddCommand(ssoAdmin_describeApplicationCmd)
}
