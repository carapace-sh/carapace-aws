package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_updateAppVersionAppComponentCmd = &cobra.Command{
	Use:   "update-app-version-app-component",
	Short: "Updates an existing Application Component in the Resilience Hub application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_updateAppVersionAppComponentCmd).Standalone()

	resiliencehub_updateAppVersionAppComponentCmd.Flags().String("additional-info", "", "Currently, there is no supported additional information for Application Components.")
	resiliencehub_updateAppVersionAppComponentCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
	resiliencehub_updateAppVersionAppComponentCmd.Flags().String("id", "", "Identifier of the Application Component.")
	resiliencehub_updateAppVersionAppComponentCmd.Flags().String("name", "", "Name of the Application Component.")
	resiliencehub_updateAppVersionAppComponentCmd.Flags().String("type", "", "Type of Application Component.")
	resiliencehub_updateAppVersionAppComponentCmd.MarkFlagRequired("app-arn")
	resiliencehub_updateAppVersionAppComponentCmd.MarkFlagRequired("id")
	resiliencehubCmd.AddCommand(resiliencehub_updateAppVersionAppComponentCmd)
}
