package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_activateOrganizationsAccessCmd = &cobra.Command{
	Use:   "activate-organizations-access",
	Short: "Activate trusted access with Organizations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_activateOrganizationsAccessCmd).Standalone()

	cloudformationCmd.AddCommand(cloudformation_activateOrganizationsAccessCmd)
}
