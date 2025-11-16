package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsData_describeGroupCmd = &cobra.Command{
	Use:   "describe-group",
	Short: "Returns information about a specific group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsData_describeGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dsData_describeGroupCmd).Standalone()

		dsData_describeGroupCmd.Flags().String("directory-id", "", "The Identifier (ID) of the directory associated with the group.")
		dsData_describeGroupCmd.Flags().String("other-attributes", "", "One or more attributes to be returned for the group.")
		dsData_describeGroupCmd.Flags().String("realm", "", "The domain name that's associated with the group.")
		dsData_describeGroupCmd.Flags().String("samaccount-name", "", "The name of the group.")
		dsData_describeGroupCmd.MarkFlagRequired("directory-id")
		dsData_describeGroupCmd.MarkFlagRequired("samaccount-name")
	})
	dsDataCmd.AddCommand(dsData_describeGroupCmd)
}
