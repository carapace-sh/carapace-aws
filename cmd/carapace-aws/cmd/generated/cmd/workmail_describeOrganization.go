package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_describeOrganizationCmd = &cobra.Command{
	Use:   "describe-organization",
	Short: "Provides more information regarding a given organization based on its identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_describeOrganizationCmd).Standalone()

	workmail_describeOrganizationCmd.Flags().String("organization-id", "", "The identifier for the organization to be described.")
	workmail_describeOrganizationCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_describeOrganizationCmd)
}
