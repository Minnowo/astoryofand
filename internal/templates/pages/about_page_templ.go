// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.696
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/minnowo/astoryofand/internal/templates"
	"github.com/minnowo/astoryofand/internal/templates/layout"
)

func ShowAboutPage() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>\n\n            main {\n                background: rgba(255,255,255,0.5);\n            }\n\n            .special_form {\n                height: 100%;\n            }\n\n            .title {\n                max-width: 100%;\n                height: auto;\n                color: black;\n            }\n\n            .body {\n                width: 80% !important;\n                padding-top: 100px;\n                padding-bottom: 100px;\n                font-weight: bold;\n            }\n            .image1, .image2, .image3 {\n                display: inline-block;\n                height:fit-content!important;\n                max-width: none !important;\n                max-height: none !important;\n            }\n\n            .image1 {\n                width: 15%  !important;\n            }\n\n            .image2 {\n                width: 55%  !important;\n            }\n\n            .image3 {\n                width: 50px  !important;\n                margin-left: 0;\n            }\n\n\n            .quote_container {\n                display: flex;\n            }\n\n\n        </style> <main><form class=\"special_form\"><div class=\"quote_container\"><img class=\"title nobg noborder\" src=\"/static/images/title.png\" style=\"margin-bottom: -50px;\" alt=\"Have You Ever Been &#39;Butted&#39;?\"></div><div class=\"quote_container bg-bubble\"><center><h1 class=\"body title nobg noborder text-3xl\">A Story of And asks you to consider replacing the word 'But' with the word 'And' as a way to connect, reflect and have meaningful conversations.  Using 'east Asian' racism as an entry point, this handheld poster series brings together symbols, visuals, analogies and prompting questions to introduce concepts related to First Nations and  east Asian worldviews and beyond.  Suitable for all ages.</h1></center></div><div class=\"quote_container\"><img class=\"image3 flipim nobg noborder\" src=\"/static/images/megaphone_man.png\" alt=\"Man with a megaphone blowing the speech bubble from the above image\"><div style=\"max-width: 850px;\"><img class=\"image1 nobg noborder\" src=\"/static/images/mich.png\" alt=\"Picture of author &#39;Mich&#39;\"> <img class=\"image2 nobg noborder\" src=\"/static/images/text.png\" alt=\"We hope that &#39;A Story of And&#39; helps you to see how everything is connected, including the impact racism has on individuals, communities, and non-humans.\"> <img class=\"image1 nobg noborder\" src=\"/static/images/jo.png\" style=\"margin-top: 10%;\" alt=\"Picture of author &#39;Jo&#39;\"></div></div></form></main>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout.Page2("About | "+templates.PAGE_TITLE).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
