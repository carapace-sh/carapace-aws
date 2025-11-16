package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detective_describeOrganizationConfigurationCmd = &cobra.Command{
	Use:   "describe-organization-configuration",
	Short: "Returns information about the configuration for the organization behavior graph.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detective_describeOrganizationConfigurationCmd).Standalone()

	detective_describeOrganizationConfigurationCmd.Flags().String("graph-arn", "", "The ARN of the organization behavior graph.")
	detective_describeOrganizationConfigurationCmd.MarkFlagRequired("graph-arn")
	detectiveCmd.AddCommand(detective_describeOrganizationConfigurationCmd)
}
