package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var braket_searchDevicesCmd = &cobra.Command{
	Use:   "search-devices",
	Short: "Searches for devices using the specified filters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(braket_searchDevicesCmd).Standalone()

	braket_searchDevicesCmd.Flags().String("filters", "", "Array of SearchDevicesFilter objects to use when searching for devices.")
	braket_searchDevicesCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	braket_searchDevicesCmd.Flags().String("next-token", "", "A token used for pagination of results returned in the response.")
	braket_searchDevicesCmd.MarkFlagRequired("filters")
	braketCmd.AddCommand(braket_searchDevicesCmd)
}
