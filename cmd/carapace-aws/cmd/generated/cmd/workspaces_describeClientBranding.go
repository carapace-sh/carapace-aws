package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_describeClientBrandingCmd = &cobra.Command{
	Use:   "describe-client-branding",
	Short: "Describes the specified client branding.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_describeClientBrandingCmd).Standalone()

	workspaces_describeClientBrandingCmd.Flags().String("resource-id", "", "The directory identifier of the WorkSpace for which you want to view client branding information.")
	workspaces_describeClientBrandingCmd.MarkFlagRequired("resource-id")
	workspacesCmd.AddCommand(workspaces_describeClientBrandingCmd)
}
