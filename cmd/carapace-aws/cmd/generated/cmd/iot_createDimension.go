package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createDimensionCmd = &cobra.Command{
	Use:   "create-dimension",
	Short: "Create a dimension that you can use to limit the scope of a metric used in a security profile for IoT Device Defender.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createDimensionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_createDimensionCmd).Standalone()

		iot_createDimensionCmd.Flags().String("client-request-token", "", "Each dimension must have a unique client request token.")
		iot_createDimensionCmd.Flags().String("name", "", "A unique identifier for the dimension.")
		iot_createDimensionCmd.Flags().String("string-values", "", "Specifies the value or list of values for the dimension.")
		iot_createDimensionCmd.Flags().String("tags", "", "Metadata that can be used to manage the dimension.")
		iot_createDimensionCmd.Flags().String("type", "", "Specifies the type of dimension.")
		iot_createDimensionCmd.MarkFlagRequired("client-request-token")
		iot_createDimensionCmd.MarkFlagRequired("name")
		iot_createDimensionCmd.MarkFlagRequired("string-values")
		iot_createDimensionCmd.MarkFlagRequired("type")
	})
	iotCmd.AddCommand(iot_createDimensionCmd)
}
