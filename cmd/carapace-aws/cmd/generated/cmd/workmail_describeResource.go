package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_describeResourceCmd = &cobra.Command{
	Use:   "describe-resource",
	Short: "Returns the data available for the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_describeResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_describeResourceCmd).Standalone()

		workmail_describeResourceCmd.Flags().String("organization-id", "", "The identifier associated with the organization for which the resource is described.")
		workmail_describeResourceCmd.Flags().String("resource-id", "", "The identifier of the resource to be described.")
		workmail_describeResourceCmd.MarkFlagRequired("organization-id")
		workmail_describeResourceCmd.MarkFlagRequired("resource-id")
	})
	workmailCmd.AddCommand(workmail_describeResourceCmd)
}
