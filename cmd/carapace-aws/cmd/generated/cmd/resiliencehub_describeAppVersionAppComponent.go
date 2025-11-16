package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_describeAppVersionAppComponentCmd = &cobra.Command{
	Use:   "describe-app-version-app-component",
	Short: "Describes an Application Component in the Resilience Hub application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_describeAppVersionAppComponentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resiliencehub_describeAppVersionAppComponentCmd).Standalone()

		resiliencehub_describeAppVersionAppComponentCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
		resiliencehub_describeAppVersionAppComponentCmd.Flags().String("app-version", "", "Resilience Hub application version.")
		resiliencehub_describeAppVersionAppComponentCmd.Flags().String("id", "", "Identifier of the Application Component.")
		resiliencehub_describeAppVersionAppComponentCmd.MarkFlagRequired("app-arn")
		resiliencehub_describeAppVersionAppComponentCmd.MarkFlagRequired("app-version")
		resiliencehub_describeAppVersionAppComponentCmd.MarkFlagRequired("id")
	})
	resiliencehubCmd.AddCommand(resiliencehub_describeAppVersionAppComponentCmd)
}
