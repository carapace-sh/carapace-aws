package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_deleteApplicationCmd = &cobra.Command{
	Use:   "delete-application",
	Short: "Deletes the association with the application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_deleteApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_deleteApplicationCmd).Standalone()

		ssoAdmin_deleteApplicationCmd.Flags().String("application-arn", "", "Specifies the ARN of the application.")
		ssoAdmin_deleteApplicationCmd.MarkFlagRequired("application-arn")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_deleteApplicationCmd)
}
