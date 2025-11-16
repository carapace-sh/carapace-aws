package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_getApplicationGrantCmd = &cobra.Command{
	Use:   "get-application-grant",
	Short: "Retrieves details about an application grant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_getApplicationGrantCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_getApplicationGrantCmd).Standalone()

		ssoAdmin_getApplicationGrantCmd.Flags().String("application-arn", "", "Specifies the ARN of the application that contains the grant.")
		ssoAdmin_getApplicationGrantCmd.Flags().String("grant-type", "", "Specifies the type of grant.")
		ssoAdmin_getApplicationGrantCmd.MarkFlagRequired("application-arn")
		ssoAdmin_getApplicationGrantCmd.MarkFlagRequired("grant-type")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_getApplicationGrantCmd)
}
