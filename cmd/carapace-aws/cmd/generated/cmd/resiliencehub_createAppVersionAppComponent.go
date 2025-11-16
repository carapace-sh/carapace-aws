package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_createAppVersionAppComponentCmd = &cobra.Command{
	Use:   "create-app-version-app-component",
	Short: "Creates a new Application Component in the Resilience Hub application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_createAppVersionAppComponentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resiliencehub_createAppVersionAppComponentCmd).Standalone()

		resiliencehub_createAppVersionAppComponentCmd.Flags().String("additional-info", "", "Currently, there is no supported additional information for Application Components.")
		resiliencehub_createAppVersionAppComponentCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
		resiliencehub_createAppVersionAppComponentCmd.Flags().String("client-token", "", "Used for an idempotency token.")
		resiliencehub_createAppVersionAppComponentCmd.Flags().String("id", "", "Identifier of the Application Component.")
		resiliencehub_createAppVersionAppComponentCmd.Flags().String("name", "", "Name of the Application Component.")
		resiliencehub_createAppVersionAppComponentCmd.Flags().String("type", "", "Type of Application Component.")
		resiliencehub_createAppVersionAppComponentCmd.MarkFlagRequired("app-arn")
		resiliencehub_createAppVersionAppComponentCmd.MarkFlagRequired("name")
		resiliencehub_createAppVersionAppComponentCmd.MarkFlagRequired("type")
	})
	resiliencehubCmd.AddCommand(resiliencehub_createAppVersionAppComponentCmd)
}
