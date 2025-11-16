package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listAuthorizersCmd = &cobra.Command{
	Use:   "list-authorizers",
	Short: "Lists the authorizers registered in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listAuthorizersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listAuthorizersCmd).Standalone()

		iot_listAuthorizersCmd.Flags().String("ascending-order", "", "Return the list of authorizers in ascending alphabetical order.")
		iot_listAuthorizersCmd.Flags().String("marker", "", "A marker used to get the next set of results.")
		iot_listAuthorizersCmd.Flags().String("page-size", "", "The maximum number of results to return at one time.")
		iot_listAuthorizersCmd.Flags().String("status", "", "The status of the list authorizers request.")
	})
	iotCmd.AddCommand(iot_listAuthorizersCmd)
}
