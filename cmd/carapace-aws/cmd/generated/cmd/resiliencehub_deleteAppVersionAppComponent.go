package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_deleteAppVersionAppComponentCmd = &cobra.Command{
	Use:   "delete-app-version-app-component",
	Short: "Deletes an Application Component from the Resilience Hub application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_deleteAppVersionAppComponentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resiliencehub_deleteAppVersionAppComponentCmd).Standalone()

		resiliencehub_deleteAppVersionAppComponentCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
		resiliencehub_deleteAppVersionAppComponentCmd.Flags().String("client-token", "", "Used for an idempotency token.")
		resiliencehub_deleteAppVersionAppComponentCmd.Flags().String("id", "", "Identifier of the Application Component.")
		resiliencehub_deleteAppVersionAppComponentCmd.MarkFlagRequired("app-arn")
		resiliencehub_deleteAppVersionAppComponentCmd.MarkFlagRequired("id")
	})
	resiliencehubCmd.AddCommand(resiliencehub_deleteAppVersionAppComponentCmd)
}
