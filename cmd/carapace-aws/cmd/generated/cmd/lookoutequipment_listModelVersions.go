package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_listModelVersionsCmd = &cobra.Command{
	Use:   "list-model-versions",
	Short: "Generates a list of all model versions for a given model, including the model version, model version ARN, and status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_listModelVersionsCmd).Standalone()

	lookoutequipment_listModelVersionsCmd.Flags().String("created-at-end-time", "", "Filter results to return all the model versions created before this time.")
	lookoutequipment_listModelVersionsCmd.Flags().String("created-at-start-time", "", "Filter results to return all the model versions created after this time.")
	lookoutequipment_listModelVersionsCmd.Flags().String("max-model-version", "", "Specifies the highest version of the model to return in the list.")
	lookoutequipment_listModelVersionsCmd.Flags().String("max-results", "", "Specifies the maximum number of machine learning model versions to list.")
	lookoutequipment_listModelVersionsCmd.Flags().String("min-model-version", "", "Specifies the lowest version of the model to return in the list.")
	lookoutequipment_listModelVersionsCmd.Flags().String("model-name", "", "Then name of the machine learning model for which the model versions are to be listed.")
	lookoutequipment_listModelVersionsCmd.Flags().String("next-token", "", "If the total number of results exceeds the limit that the response can display, the response returns an opaque pagination token indicating where to continue the listing of machine learning model versions.")
	lookoutequipment_listModelVersionsCmd.Flags().String("source-type", "", "Filter the results based on the way the model version was generated.")
	lookoutequipment_listModelVersionsCmd.Flags().String("status", "", "Filter the results based on the current status of the model version.")
	lookoutequipment_listModelVersionsCmd.MarkFlagRequired("model-name")
	lookoutequipmentCmd.AddCommand(lookoutequipment_listModelVersionsCmd)
}
