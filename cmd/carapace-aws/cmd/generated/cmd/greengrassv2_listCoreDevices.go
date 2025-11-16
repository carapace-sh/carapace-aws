package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2_listCoreDevicesCmd = &cobra.Command{
	Use:   "list-core-devices",
	Short: "Retrieves a paginated list of Greengrass core devices.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2_listCoreDevicesCmd).Standalone()

	greengrassv2_listCoreDevicesCmd.Flags().String("max-results", "", "The maximum number of results to be returned per paginated request.")
	greengrassv2_listCoreDevicesCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	greengrassv2_listCoreDevicesCmd.Flags().String("runtime", "", "The runtime to be used by the core device.")
	greengrassv2_listCoreDevicesCmd.Flags().String("status", "", "The core device status by which to filter.")
	greengrassv2_listCoreDevicesCmd.Flags().String("thing-group-arn", "", "The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the IoT thing group by which to filter.")
	greengrassv2Cmd.AddCommand(greengrassv2_listCoreDevicesCmd)
}
