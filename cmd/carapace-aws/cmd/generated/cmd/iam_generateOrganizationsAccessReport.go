package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_generateOrganizationsAccessReportCmd = &cobra.Command{
	Use:   "generate-organizations-access-report",
	Short: "Generates a report for service last accessed data for Organizations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_generateOrganizationsAccessReportCmd).Standalone()

	iam_generateOrganizationsAccessReportCmd.Flags().String("entity-path", "", "The path of the Organizations entity (root, OU, or account).")
	iam_generateOrganizationsAccessReportCmd.Flags().String("organizations-policy-id", "", "The identifier of the Organizations service control policy (SCP).")
	iam_generateOrganizationsAccessReportCmd.MarkFlagRequired("entity-path")
	iamCmd.AddCommand(iam_generateOrganizationsAccessReportCmd)
}
