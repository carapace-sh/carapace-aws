package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_createOrganizationCmd = &cobra.Command{
	Use:   "create-organization",
	Short: "Creates an Amazon Web Services organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_createOrganizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(organizations_createOrganizationCmd).Standalone()

		organizations_createOrganizationCmd.Flags().String("feature-set", "", "Specifies the feature set supported by the new organization.")
	})
	organizationsCmd.AddCommand(organizations_createOrganizationCmd)
}
