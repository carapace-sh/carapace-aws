package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_listPermissionAssociationsCmd = &cobra.Command{
	Use:   "list-permission-associations",
	Short: "Lists information about the managed permission and its associations to any resource shares that use this managed permission.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_listPermissionAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ram_listPermissionAssociationsCmd).Standalone()

		ram_listPermissionAssociationsCmd.Flags().String("association-status", "", "Specifies that you want to list only those associations with resource shares that match this status.")
		ram_listPermissionAssociationsCmd.Flags().Bool("default-version", false, "When `true`, specifies that you want to list only those associations with resource shares that use the default version of the specified managed permission.")
		ram_listPermissionAssociationsCmd.Flags().String("feature-set", "", "Specifies that you want to list only those associations with resource shares that have a `featureSet` with this value.")
		ram_listPermissionAssociationsCmd.Flags().String("max-results", "", "Specifies the total number of results that you want included on each page of the response.")
		ram_listPermissionAssociationsCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
		ram_listPermissionAssociationsCmd.Flags().Bool("no-default-version", false, "When `true`, specifies that you want to list only those associations with resource shares that use the default version of the specified managed permission.")
		ram_listPermissionAssociationsCmd.Flags().String("permission-arn", "", "Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the managed permission.")
		ram_listPermissionAssociationsCmd.Flags().String("permission-version", "", "Specifies that you want to list only those associations with resource shares that use this version of the managed permission.")
		ram_listPermissionAssociationsCmd.Flags().String("resource-type", "", "Specifies that you want to list only those associations with resource shares that include at least one resource of this resource type.")
		ram_listPermissionAssociationsCmd.Flag("no-default-version").Hidden = true
	})
	ramCmd.AddCommand(ram_listPermissionAssociationsCmd)
}
