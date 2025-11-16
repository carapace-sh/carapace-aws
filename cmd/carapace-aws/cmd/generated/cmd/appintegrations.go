package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appintegrationsCmd = &cobra.Command{
	Use:   "appintegrations",
	Short: "- [Amazon AppIntegrations actions](https://docs.aws.amazon.com/connect/latest/APIReference/API_Operations_Amazon_AppIntegrations_Service.html)\n- [Amazon AppIntegrations data types](https://docs.aws.amazon.com/connect/latest/APIReference/API_Types_Amazon_AppIntegrations_Service.html)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appintegrationsCmd).Standalone()

	rootCmd.AddCommand(appintegrationsCmd)
}
