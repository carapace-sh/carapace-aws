package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_publishAppVersionCmd = &cobra.Command{
	Use:   "publish-app-version",
	Short: "Publishes a new version of a specific Resilience Hub application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_publishAppVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resiliencehub_publishAppVersionCmd).Standalone()

		resiliencehub_publishAppVersionCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
		resiliencehub_publishAppVersionCmd.Flags().String("version-name", "", "Name of the application version.")
		resiliencehub_publishAppVersionCmd.MarkFlagRequired("app-arn")
	})
	resiliencehubCmd.AddCommand(resiliencehub_publishAppVersionCmd)
}
