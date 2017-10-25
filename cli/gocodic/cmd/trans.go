package cmd

import (
	"bytes"
	"io"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/spiegel-im-spiegel/gocodic"
	"github.com/spiegel-im-spiegel/gocodic/options"
	"github.com/spiegel-im-spiegel/gocodic/response"
)

// transCmd represents the trans command
var transCmd = &cobra.Command{
	Use:   "trans [flags] [<word>...]",
	Short: "Ttansration API for codic.jp",
	Long:  "Ttansration API for codic.jp",
	RunE: func(cmd *cobra.Command, args []string) error {
		opts, err := options.NewOptions(viper.GetString("token"))
		if err != nil {
			return err
		}

		if len(args) == 0 {
			var buf bytes.Buffer
			io.Copy(&buf, reader)
			opts.Add(options.Text(string(buf.Bytes())))
		} else {
			opts.Add(options.Text(strings.Join(args, "\n")))
		}

		jsonFlag := viper.GetBool("json")

		pid := viper.GetInt("projectid")
		if pid > 0 {
			opts.Add(options.ProjectID(pid))
		}
		casing, err := options.NewCasingOption(viper.GetString("casing"))
		if err == nil {
			opts.Add(casing)
		}
		style, err := options.NewAcronymStyleOption(viper.GetString("style"))
		if err == nil {
			opts.Add(style)
		}

		res, err := gocodic.Translate(opts)
		if err != nil {
			return err
		}
		if !res.IsSuccess() {
			result = ""
			resultErr = ""
			if jsonFlag {
				resultErr = res.String()
				return nil
			}
			ed, err2 := response.DecodeError(res.Body())
			if err != nil {
				return err2
			}
			for _, d := range ed.Errors {
				resultErr += d.Message + "\n"
			}
			return nil
		}
		result = ""
		resultErr = ""
		if jsonFlag {
			result = res.String()
			return nil
		}
		sd, err := response.DecodeSuccessTrans(res.Body())
		if err != nil {
			return err
		}
		for _, d := range sd {
			result += d.TranslatedText + "\n"
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(transCmd)

	transCmd.Flags().IntP("projectid", "p", 0, "project_id parameter")
	transCmd.Flags().StringP("casing", "c", "", "casing parameter")
	transCmd.Flags().StringP("style", "s", "", "acronym_style parameter")
	viper.BindPFlag("projectid", transCmd.Flags().Lookup("projectid"))
	viper.BindPFlag("casing", transCmd.Flags().Lookup("casing"))
	viper.BindPFlag("style", transCmd.Flags().Lookup("style"))

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// transCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// transCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
