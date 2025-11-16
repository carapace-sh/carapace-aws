package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_listInstanceProfilesCmd = &cobra.Command{
	Use:   "list-instance-profiles",
	Short: "Returns information about all the instance profiles in an AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_listInstanceProfilesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_listInstanceProfilesCmd).Standalone()

		devicefarm_listInstanceProfilesCmd.Flags().String("max-results", "", "An integer that specifies the maximum number of items you want to return in the API response.")
		devicefarm_listInstanceProfilesCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which can be used to return the next set of items in the list.")
	})
	devicefarmCmd.AddCommand(devicefarm_listInstanceProfilesCmd)
}
