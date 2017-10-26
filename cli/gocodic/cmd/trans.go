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
	Short: "Transration API for codic.jp",
	Long:  "Transration API for codic.jp",
	RunE: func(cmd *cobra.Command, args []string) error {
		opts, err := options.NewOptions(options.CmdTrans, viper.GetString("token"))
		if err != nil {
			return err
		}

		if len(args) == 0 {
			var buf bytes.Buffer
			io.Copy(&buf, cui.Reader())
			opts.Add(options.Text(string(buf.Bytes())))
		} else {
			opts.Add(options.Text(strings.Join(args, "\n")))
		}

		jsonFlag := viper.GetBool("json")

		pid := viper.GetInt("projectid")
		if pid > 0 {
			opts.Add(options.ProjectID(pid))
		}
		if casing, ok := options.NewCasingOption(viper.GetString("casing")); ok {
			opts.Add(casing)
		}
		if style, ok := options.NewAcronymStyleOption(viper.GetString("style")); ok {
			opts.Add(style)
		}

		res, err := gocodic.Translate(opts)
		if err != nil {
			return err
		}
		if !res.IsSuccess() {
			errFlag = true
			if jsonFlag {
				cui.WriteErrFrom(res.Body())
				return nil
			}
			ed, err2 := response.DecodeError(res.Body())
			if err != nil {
				return err2
			}
			for _, d := range ed.Errors {
				cui.OutputErrln(d.Message)
			}
			return nil
		}
		if jsonFlag {
			cui.WriteFrom(res.Body())
			return nil
		}
		sd, err := response.DecodeSuccessTrans(res.Body())
		if err != nil {
			return err
		}
		for _, d := range sd {
			cui.Outputln(d.TranslatedText)
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
