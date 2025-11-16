package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var support_describeServicesCmd = &cobra.Command{
	Use:   "describe-services",
	Short: "Returns the current list of Amazon Web Services services and a list of service categories for each service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(support_describeServicesCmd).Standalone()

	support_describeServicesCmd.Flags().String("language", "", "The language in which Amazon Web Services Support handles the case.")
	support_describeServicesCmd.Flags().String("service-code-list", "", "A JSON-formatted list of service codes available for Amazon Web Services services.")
	supportCmd.AddCommand(support_describeServicesCmd)
}
