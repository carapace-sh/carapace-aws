package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_deleteOrganizationConformancePackCmd = &cobra.Command{
	Use:   "delete-organization-conformance-pack",
	Short: "Deletes the specified organization conformance pack and all of the Config rules and remediation actions from all member accounts in that organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_deleteOrganizationConformancePackCmd).Standalone()

	config_deleteOrganizationConformancePackCmd.Flags().String("organization-conformance-pack-name", "", "The name of organization conformance pack that you want to delete.")
	config_deleteOrganizationConformancePackCmd.MarkFlagRequired("organization-conformance-pack-name")
	configCmd.AddCommand(config_deleteOrganizationConformancePackCmd)
}
