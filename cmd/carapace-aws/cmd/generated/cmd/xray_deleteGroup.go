package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_deleteGroupCmd = &cobra.Command{
	Use:   "delete-group",
	Short: "Deletes a group resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_deleteGroupCmd).Standalone()

	xray_deleteGroupCmd.Flags().String("group-arn", "", "The ARN of the group that was generated on creation.")
	xray_deleteGroupCmd.Flags().String("group-name", "", "The case-sensitive name of the group.")
	xrayCmd.AddCommand(xray_deleteGroupCmd)
}
