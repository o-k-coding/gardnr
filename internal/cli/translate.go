package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"okcoding.com/gardnr/internal/translate"
)

// TODO prompt support would be awesome with a searchable table of langs
func newTranslateCmd() *cobra.Command {
	var fromLanguage string
	var toLanguage string
	translateCmd := &cobra.Command{
		Use:   "translate",
		Short: "Command for translating words using the google translate API.",
		Long:  "Support for translating text from one language to another using the google translate API.",
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) == 0 {
				return fmt.Errorf("text is required in order to translate")
			}
			text := args[0]
			result, err := translate.TranslateText(fromLanguage, toLanguage, text)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	translateCmd.Flags().StringVarP(&fromLanguage, "from-lang", "f", "en-us", "Language code to translate from, defaults to en-us. Code Reference: https://cloud.google.com/translate/docs/languages")
	translateCmd.Flags().StringVarP(&toLanguage, "to-lang", "l", "es", "Language code to translate to, defaults to es. Code Reference: https://cloud.google.com/translate/docs/languages")

	return translateCmd
}
