package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_putApplicationGrantCmd = &cobra.Command{
	Use:   "put-application-grant",
	Short: "Creates a configuration for an application to use grants.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_putApplicationGrantCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_putApplicationGrantCmd).Standalone()

		ssoAdmin_putApplicationGrantCmd.Flags().String("application-arn", "", "Specifies the ARN of the application to update.")
		ssoAdmin_putApplicationGrantCmd.Flags().String("grant", "", "Specifies a structure that describes the grant to update.")
		ssoAdmin_putApplicationGrantCmd.Flags().String("grant-type", "", "Specifies the type of grant to update.")
		ssoAdmin_putApplicationGrantCmd.MarkFlagRequired("application-arn")
		ssoAdmin_putApplicationGrantCmd.MarkFlagRequired("grant")
		ssoAdmin_putApplicationGrantCmd.MarkFlagRequired("grant-type")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_putApplicationGrantCmd)
}
