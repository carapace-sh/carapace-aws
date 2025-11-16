package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getCloudFormationStackRecordsCmd = &cobra.Command{
	Use:   "get-cloud-formation-stack-records",
	Short: "Returns the CloudFormation stack record created as a result of the `create cloud formation stack` operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getCloudFormationStackRecordsCmd).Standalone()

	lightsail_getCloudFormationStackRecordsCmd.Flags().String("page-token", "", "The token to advance to the next page of results from your request.")
	lightsailCmd.AddCommand(lightsail_getCloudFormationStackRecordsCmd)
}
