package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_deactivateOrganizationsAccessCmd = &cobra.Command{
	Use:   "deactivate-organizations-access",
	Short: "Deactivates trusted access with Organizations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_deactivateOrganizationsAccessCmd).Standalone()

	cloudformationCmd.AddCommand(cloudformation_deactivateOrganizationsAccessCmd)
}
