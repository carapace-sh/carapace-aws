package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_createPredictorCmd = &cobra.Command{
	Use:   "create-predictor",
	Short: "This operation creates a legacy predictor that does not include all the predictor functionalities provided by Amazon Forecast.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_createPredictorCmd).Standalone()

	forecast_createPredictorCmd.Flags().String("algorithm-arn", "", "The Amazon Resource Name (ARN) of the algorithm to use for model training.")
	forecast_createPredictorCmd.Flags().String("auto-mloverride-strategy", "", "The `LatencyOptimized` AutoML override strategy is only available in private beta.")
	forecast_createPredictorCmd.Flags().String("encryption-config", "", "An Key Management Service (KMS) key and the Identity and Access Management (IAM) role that Amazon Forecast can assume to access the key.")
	forecast_createPredictorCmd.Flags().String("evaluation-parameters", "", "Used to override the default evaluation parameters of the specified algorithm.")
	forecast_createPredictorCmd.Flags().String("featurization-config", "", "The featurization configuration.")
	forecast_createPredictorCmd.Flags().String("forecast-horizon", "", "Specifies the number of time-steps that the model is trained to predict.")
	forecast_createPredictorCmd.Flags().String("forecast-types", "", "Specifies the forecast types used to train a predictor.")
	forecast_createPredictorCmd.Flags().String("hpoconfig", "", "Provides hyperparameter override values for the algorithm.")
	forecast_createPredictorCmd.Flags().String("input-data-config", "", "Describes the dataset group that contains the data to use to train the predictor.")
	forecast_createPredictorCmd.Flags().Bool("no-perform-auto-ml", false, "Whether to perform AutoML.")
	forecast_createPredictorCmd.Flags().Bool("no-perform-hpo", false, "Whether to perform hyperparameter optimization (HPO).")
	forecast_createPredictorCmd.Flags().String("optimization-metric", "", "The accuracy metric used to optimize the predictor.")
	forecast_createPredictorCmd.Flags().Bool("perform-auto-ml", false, "Whether to perform AutoML.")
	forecast_createPredictorCmd.Flags().Bool("perform-hpo", false, "Whether to perform hyperparameter optimization (HPO).")
	forecast_createPredictorCmd.Flags().String("predictor-name", "", "A name for the predictor.")
	forecast_createPredictorCmd.Flags().String("tags", "", "The optional metadata that you apply to the predictor to help you categorize and organize them.")
	forecast_createPredictorCmd.Flags().String("training-parameters", "", "The hyperparameters to override for model training.")
	forecast_createPredictorCmd.MarkFlagRequired("featurization-config")
	forecast_createPredictorCmd.MarkFlagRequired("forecast-horizon")
	forecast_createPredictorCmd.MarkFlagRequired("input-data-config")
	forecast_createPredictorCmd.Flag("no-perform-auto-ml").Hidden = true
	forecast_createPredictorCmd.Flag("no-perform-hpo").Hidden = true
	forecast_createPredictorCmd.MarkFlagRequired("predictor-name")
	forecastCmd.AddCommand(forecast_createPredictorCmd)
}
