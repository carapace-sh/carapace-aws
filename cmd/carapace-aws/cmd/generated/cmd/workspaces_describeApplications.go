package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_describeApplicationsCmd = &cobra.Command{
	Use:   "describe-applications",
	Short: "Describes the specified applications by filtering based on their compute types, license availability, operating systems, and owners.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_describeApplicationsCmd).Standalone()

	workspaces_describeApplicationsCmd.Flags().String("application-ids", "", "The identifiers of one or more applications.")
	workspaces_describeApplicationsCmd.Flags().String("compute-type-names", "", "The compute types supported by the applications.")
	workspaces_describeApplicationsCmd.Flags().String("license-type", "", "The license availability for the applications.")
	workspaces_describeApplicationsCmd.Flags().String("max-results", "", "The maximum number of applications to return.")
	workspaces_describeApplicationsCmd.Flags().String("next-token", "", "If you received a `NextToken` from a previous call that was paginated, provide this token to receive the next set of results.")
	workspaces_describeApplicationsCmd.Flags().String("operating-system-names", "", "The operating systems supported by the applications.")
	workspaces_describeApplicationsCmd.Flags().String("owner", "", "The owner of the applications.")
	workspacesCmd.AddCommand(workspaces_describeApplicationsCmd)
}
