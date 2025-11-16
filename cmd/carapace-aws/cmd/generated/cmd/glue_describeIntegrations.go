package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_describeIntegrationsCmd = &cobra.Command{
	Use:   "describe-integrations",
	Short: "The API is used to retrieve a list of integrations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_describeIntegrationsCmd).Standalone()

	glue_describeIntegrationsCmd.Flags().String("filters", "", "A list of key and values, to filter down the results.")
	glue_describeIntegrationsCmd.Flags().String("integration-identifier", "", "The Amazon Resource Name (ARN) for the integration.")
	glue_describeIntegrationsCmd.Flags().String("marker", "", "A value that indicates the starting point for the next set of response records in a subsequent request.")
	glue_describeIntegrationsCmd.Flags().String("max-records", "", "The total number of items to return in the output.")
	glueCmd.AddCommand(glue_describeIntegrationsCmd)
}
