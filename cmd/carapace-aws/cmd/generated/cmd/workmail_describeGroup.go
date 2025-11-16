package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_describeGroupCmd = &cobra.Command{
	Use:   "describe-group",
	Short: "Returns the data available for the group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_describeGroupCmd).Standalone()

	workmail_describeGroupCmd.Flags().String("group-id", "", "The identifier for the group to be described.")
	workmail_describeGroupCmd.Flags().String("organization-id", "", "The identifier for the organization under which the group exists.")
	workmail_describeGroupCmd.MarkFlagRequired("group-id")
	workmail_describeGroupCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_describeGroupCmd)
}
