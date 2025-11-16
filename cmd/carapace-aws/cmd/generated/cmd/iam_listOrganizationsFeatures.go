package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listOrganizationsFeaturesCmd = &cobra.Command{
	Use:   "list-organizations-features",
	Short: "Lists the centralized root access features enabled for your organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listOrganizationsFeaturesCmd).Standalone()

	iamCmd.AddCommand(iam_listOrganizationsFeaturesCmd)
}
