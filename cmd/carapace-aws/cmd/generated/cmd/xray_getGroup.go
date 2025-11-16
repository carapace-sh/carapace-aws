package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_getGroupCmd = &cobra.Command{
	Use:   "get-group",
	Short: "Retrieves group resource details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_getGroupCmd).Standalone()

	xray_getGroupCmd.Flags().String("group-arn", "", "The ARN of the group that was generated on creation.")
	xray_getGroupCmd.Flags().String("group-name", "", "The case-sensitive name of the group.")
	xrayCmd.AddCommand(xray_getGroupCmd)
}
