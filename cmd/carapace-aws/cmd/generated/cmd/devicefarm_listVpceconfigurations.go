package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_listVpceconfigurationsCmd = &cobra.Command{
	Use:   "list-vpceconfigurations",
	Short: "Returns information about all Amazon Virtual Private Cloud (VPC) endpoint configurations in the AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_listVpceconfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_listVpceconfigurationsCmd).Standalone()

		devicefarm_listVpceconfigurationsCmd.Flags().String("max-results", "", "An integer that specifies the maximum number of items you want to return in the API response.")
		devicefarm_listVpceconfigurationsCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which can be used to return the next set of items in the list.")
	})
	devicefarmCmd.AddCommand(devicefarm_listVpceconfigurationsCmd)
}
