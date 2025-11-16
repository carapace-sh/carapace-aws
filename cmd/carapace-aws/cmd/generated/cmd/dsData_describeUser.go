package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsData_describeUserCmd = &cobra.Command{
	Use:   "describe-user",
	Short: "Returns information about a specific user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsData_describeUserCmd).Standalone()

	dsData_describeUserCmd.Flags().String("directory-id", "", "The identifier (ID) of the directory that's associated with the user.")
	dsData_describeUserCmd.Flags().String("other-attributes", "", "One or more attribute names to be returned for the user.")
	dsData_describeUserCmd.Flags().String("realm", "", "The domain name that's associated with the user.")
	dsData_describeUserCmd.Flags().String("samaccount-name", "", "The name of the user.")
	dsData_describeUserCmd.MarkFlagRequired("directory-id")
	dsData_describeUserCmd.MarkFlagRequired("samaccount-name")
	dsDataCmd.AddCommand(dsData_describeUserCmd)
}
