package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_detectProtectiveEquipmentCmd = &cobra.Command{
	Use:   "detect-protective-equipment",
	Short: "Detects Personal Protective Equipment (PPE) worn by people detected in an image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_detectProtectiveEquipmentCmd).Standalone()

	rekognition_detectProtectiveEquipmentCmd.Flags().String("image", "", "The image in which you want to detect PPE on detected persons.")
	rekognition_detectProtectiveEquipmentCmd.Flags().String("summarization-attributes", "", "An array of PPE types that you want to summarize.")
	rekognition_detectProtectiveEquipmentCmd.MarkFlagRequired("image")
	rekognitionCmd.AddCommand(rekognition_detectProtectiveEquipmentCmd)
}
