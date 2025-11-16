package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_enableSharingWithAwsOrganizationCmd = &cobra.Command{
	Use:   "enable-sharing-with-aws-organization",
	Short: "Enables resource sharing within your organization in Organizations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_enableSharingWithAwsOrganizationCmd).Standalone()

	ramCmd.AddCommand(ram_enableSharingWithAwsOrganizationCmd)
}
