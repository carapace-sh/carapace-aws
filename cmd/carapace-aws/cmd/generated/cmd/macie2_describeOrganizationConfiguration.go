package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_describeOrganizationConfigurationCmd = &cobra.Command{
	Use:   "describe-organization-configuration",
	Short: "Retrieves the Amazon Macie configuration settings for an organization in Organizations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_describeOrganizationConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_describeOrganizationConfigurationCmd).Standalone()

	})
	macie2Cmd.AddCommand(macie2_describeOrganizationConfigurationCmd)
}
