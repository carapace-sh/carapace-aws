package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_getEventSourceMappingCmd = &cobra.Command{
	Use:   "get-event-source-mapping",
	Short: "Returns details about an event source mapping.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_getEventSourceMappingCmd).Standalone()

	lambda_getEventSourceMappingCmd.Flags().String("uuid", "", "The identifier of the event source mapping.")
	lambda_getEventSourceMappingCmd.MarkFlagRequired("uuid")
	lambdaCmd.AddCommand(lambda_getEventSourceMappingCmd)
}
