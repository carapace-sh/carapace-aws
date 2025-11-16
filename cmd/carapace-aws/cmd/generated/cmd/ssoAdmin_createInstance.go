package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_createInstanceCmd = &cobra.Command{
	Use:   "create-instance",
	Short: "Creates an instance of IAM Identity Center for a standalone Amazon Web Services account that is not managed by Organizations or a member Amazon Web Services account in an organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_createInstanceCmd).Standalone()

	ssoAdmin_createInstanceCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive ID that you provide to ensure the idempotency of the request.")
	ssoAdmin_createInstanceCmd.Flags().String("name", "", "The name of the instance of IAM Identity Center.")
	ssoAdmin_createInstanceCmd.Flags().String("tags", "", "Specifies tags to be attached to the instance of IAM Identity Center.")
	ssoAdminCmd.AddCommand(ssoAdmin_createInstanceCmd)
}
