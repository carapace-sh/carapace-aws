package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_describeEndpointsCmd = &cobra.Command{
	Use:   "describe-endpoints",
	Short: "Returns the regional endpoint information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_describeEndpointsCmd).Standalone()

	dynamodbCmd.AddCommand(dynamodb_describeEndpointsCmd)
}
