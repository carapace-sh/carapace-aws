package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_describeAppVersionResourcesResolutionStatusCmd = &cobra.Command{
	Use:   "describe-app-version-resources-resolution-status",
	Short: "Returns the resolution status for the specified resolution identifier for an application version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_describeAppVersionResourcesResolutionStatusCmd).Standalone()

	resiliencehub_describeAppVersionResourcesResolutionStatusCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
	resiliencehub_describeAppVersionResourcesResolutionStatusCmd.Flags().String("app-version", "", "The version of the application.")
	resiliencehub_describeAppVersionResourcesResolutionStatusCmd.Flags().String("resolution-id", "", "The identifier for a specific resolution.")
	resiliencehub_describeAppVersionResourcesResolutionStatusCmd.MarkFlagRequired("app-arn")
	resiliencehub_describeAppVersionResourcesResolutionStatusCmd.MarkFlagRequired("app-version")
	resiliencehubCmd.AddCommand(resiliencehub_describeAppVersionResourcesResolutionStatusCmd)
}
