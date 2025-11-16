package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_describePublisherCmd = &cobra.Command{
	Use:   "describe-publisher",
	Short: "Returns information about a CloudFormation extension publisher.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_describePublisherCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_describePublisherCmd).Standalone()

		cloudformation_describePublisherCmd.Flags().String("publisher-id", "", "The ID of the extension publisher.")
	})
	cloudformationCmd.AddCommand(cloudformation_describePublisherCmd)
}
