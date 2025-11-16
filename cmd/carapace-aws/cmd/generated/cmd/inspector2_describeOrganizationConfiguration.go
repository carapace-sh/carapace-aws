package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_describeOrganizationConfigurationCmd = &cobra.Command{
	Use:   "describe-organization-configuration",
	Short: "Describe Amazon Inspector configuration settings for an Amazon Web Services organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_describeOrganizationConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_describeOrganizationConfigurationCmd).Standalone()

	})
	inspector2Cmd.AddCommand(inspector2_describeOrganizationConfigurationCmd)
}
