package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_updateAppVersionCmd = &cobra.Command{
	Use:   "update-app-version",
	Short: "Updates the Resilience Hub application version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_updateAppVersionCmd).Standalone()

	resiliencehub_updateAppVersionCmd.Flags().String("additional-info", "", "Additional configuration parameters for an Resilience Hub application.")
	resiliencehub_updateAppVersionCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
	resiliencehub_updateAppVersionCmd.MarkFlagRequired("app-arn")
	resiliencehubCmd.AddCommand(resiliencehub_updateAppVersionCmd)
}
