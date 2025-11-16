package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_listInferenceSchedulersCmd = &cobra.Command{
	Use:   "list-inference-schedulers",
	Short: "Retrieves a list of all inference schedulers currently available for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_listInferenceSchedulersCmd).Standalone()

	lookoutequipment_listInferenceSchedulersCmd.Flags().String("inference-scheduler-name-begins-with", "", "The beginning of the name of the inference schedulers to be listed.")
	lookoutequipment_listInferenceSchedulersCmd.Flags().String("max-results", "", "Specifies the maximum number of inference schedulers to list.")
	lookoutequipment_listInferenceSchedulersCmd.Flags().String("model-name", "", "The name of the machine learning model used by the inference scheduler to be listed.")
	lookoutequipment_listInferenceSchedulersCmd.Flags().String("next-token", "", "An opaque pagination token indicating where to continue the listing of inference schedulers.")
	lookoutequipment_listInferenceSchedulersCmd.Flags().String("status", "", "Specifies the current status of the inference schedulers.")
	lookoutequipmentCmd.AddCommand(lookoutequipment_listInferenceSchedulersCmd)
}
