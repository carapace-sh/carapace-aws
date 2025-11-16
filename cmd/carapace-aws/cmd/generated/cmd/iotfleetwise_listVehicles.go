package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_listVehiclesCmd = &cobra.Command{
	Use:   "list-vehicles",
	Short: "Retrieves a list of summaries of created vehicles.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_listVehiclesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotfleetwise_listVehiclesCmd).Standalone()

		iotfleetwise_listVehiclesCmd.Flags().String("attribute-names", "", "The fully qualified names of the attributes.")
		iotfleetwise_listVehiclesCmd.Flags().String("attribute-values", "", "Static information about a vehicle attribute value in string format.")
		iotfleetwise_listVehiclesCmd.Flags().String("list-response-scope", "", "When you set the `listResponseScope` parameter to `METADATA_ONLY`, the list response includes: vehicle name, Amazon Resource Name (ARN), creation time, and last modification time.")
		iotfleetwise_listVehiclesCmd.Flags().String("max-results", "", "The maximum number of items to return, between 1 and 100, inclusive.")
		iotfleetwise_listVehiclesCmd.Flags().String("model-manifest-arn", "", "The Amazon Resource Name (ARN) of a vehicle model (model manifest).")
		iotfleetwise_listVehiclesCmd.Flags().String("next-token", "", "A pagination token for the next set of results.")
	})
	iotfleetwiseCmd.AddCommand(iotfleetwise_listVehiclesCmd)
}
