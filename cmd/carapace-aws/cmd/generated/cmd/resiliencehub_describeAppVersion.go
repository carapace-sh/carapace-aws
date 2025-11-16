package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_describeAppVersionCmd = &cobra.Command{
	Use:   "describe-app-version",
	Short: "Describes the Resilience Hub application version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_describeAppVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resiliencehub_describeAppVersionCmd).Standalone()

		resiliencehub_describeAppVersionCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
		resiliencehub_describeAppVersionCmd.Flags().String("app-version", "", "Resilience Hub application version.")
		resiliencehub_describeAppVersionCmd.MarkFlagRequired("app-arn")
		resiliencehub_describeAppVersionCmd.MarkFlagRequired("app-version")
	})
	resiliencehubCmd.AddCommand(resiliencehub_describeAppVersionCmd)
}
