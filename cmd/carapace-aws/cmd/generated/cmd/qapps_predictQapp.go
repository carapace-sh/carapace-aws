package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_predictQappCmd = &cobra.Command{
	Use:   "predict-qapp",
	Short: "Generates an Amazon Q App definition based on either a conversation or a problem statement provided as input.The resulting app definition can be used to call `CreateQApp`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_predictQappCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qapps_predictQappCmd).Standalone()

		qapps_predictQappCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
		qapps_predictQappCmd.Flags().String("options", "", "The input to generate the Q App definition from, either a conversation or problem statement.")
		qapps_predictQappCmd.MarkFlagRequired("instance-id")
	})
	qappsCmd.AddCommand(qapps_predictQappCmd)
}
