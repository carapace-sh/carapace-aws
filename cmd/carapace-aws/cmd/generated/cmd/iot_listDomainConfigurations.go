package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listDomainConfigurationsCmd = &cobra.Command{
	Use:   "list-domain-configurations",
	Short: "Gets a list of domain configurations for the user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listDomainConfigurationsCmd).Standalone()

	iot_listDomainConfigurationsCmd.Flags().String("marker", "", "The marker for the next set of results.")
	iot_listDomainConfigurationsCmd.Flags().String("page-size", "", "The result page size.")
	iot_listDomainConfigurationsCmd.Flags().String("service-type", "", "The type of service delivered by the endpoint.")
	iotCmd.AddCommand(iot_listDomainConfigurationsCmd)
}
