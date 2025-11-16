package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2_listEffectiveDeploymentsCmd = &cobra.Command{
	Use:   "list-effective-deployments",
	Short: "Retrieves a paginated list of deployment jobs that IoT Greengrass sends to Greengrass core devices.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2_listEffectiveDeploymentsCmd).Standalone()

	greengrassv2_listEffectiveDeploymentsCmd.Flags().String("core-device-thing-name", "", "The name of the core device.")
	greengrassv2_listEffectiveDeploymentsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per paginated request.")
	greengrassv2_listEffectiveDeploymentsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	greengrassv2_listEffectiveDeploymentsCmd.MarkFlagRequired("core-device-thing-name")
	greengrassv2Cmd.AddCommand(greengrassv2_listEffectiveDeploymentsCmd)
}
