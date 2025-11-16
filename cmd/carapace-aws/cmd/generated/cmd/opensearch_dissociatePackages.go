package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_dissociatePackagesCmd = &cobra.Command{
	Use:   "dissociate-packages",
	Short: "Dissociates multiple packages from a domain simulatneously.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_dissociatePackagesCmd).Standalone()

	opensearch_dissociatePackagesCmd.Flags().String("domain-name", "", "")
	opensearch_dissociatePackagesCmd.Flags().String("package-list", "", "A list of package IDs to be dissociated from a domain.")
	opensearch_dissociatePackagesCmd.MarkFlagRequired("domain-name")
	opensearch_dissociatePackagesCmd.MarkFlagRequired("package-list")
	opensearchCmd.AddCommand(opensearch_dissociatePackagesCmd)
}
