package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_describeApplicationAssociationsCmd = &cobra.Command{
	Use:   "describe-application-associations",
	Short: "Describes the associations between the application and the specified associated resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_describeApplicationAssociationsCmd).Standalone()

	workspaces_describeApplicationAssociationsCmd.Flags().String("application-id", "", "The identifier of the specified application.")
	workspaces_describeApplicationAssociationsCmd.Flags().String("associated-resource-types", "", "The resource type of the associated resources.")
	workspaces_describeApplicationAssociationsCmd.Flags().String("max-results", "", "The maximum number of associations to return.")
	workspaces_describeApplicationAssociationsCmd.Flags().String("next-token", "", "If you received a `NextToken` from a previous call that was paginated, provide this token to receive the next set of results.")
	workspaces_describeApplicationAssociationsCmd.MarkFlagRequired("application-id")
	workspaces_describeApplicationAssociationsCmd.MarkFlagRequired("associated-resource-types")
	workspacesCmd.AddCommand(workspaces_describeApplicationAssociationsCmd)
}
