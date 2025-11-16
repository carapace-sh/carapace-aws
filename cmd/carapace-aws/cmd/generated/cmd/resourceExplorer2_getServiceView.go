package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_getServiceViewCmd = &cobra.Command{
	Use:   "get-service-view",
	Short: "Retrieves details about a specific Resource Explorer service view.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_getServiceViewCmd).Standalone()

	resourceExplorer2_getServiceViewCmd.Flags().String("service-view-arn", "", "The Amazon Resource Name (ARN) of the service view to retrieve details for.")
	resourceExplorer2_getServiceViewCmd.MarkFlagRequired("service-view-arn")
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_getServiceViewCmd)
}
