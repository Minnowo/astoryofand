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

func onClickHandler() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_onClickHandler_a72f`,
		Function: `function __templ_onClickHandler_a72f(){var audio = document.getElementsByTagName("audio")[0];
    audio.currentTime = 0;
    audio.volume = 0.5;
    audio.play();
}`,
		Call:       templ.SafeScript(`__templ_onClickHandler_a72f`),
		CallInline: templ.SafeScriptInline(`__templ_onClickHandler_a72f`),
	}
}

func flipCard() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex m-1 w-[256px] h-[192px] perspective flip-hover\"><div class=\"relative w-full h-full transition-3d flip-hover-inner\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func ShowGridPage(show_flat bool) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var3 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>\n            main {\n                background: rgba(255,255,255,0.5);\n            }\n        </style> <main><audio preload=\"auto\"><source src=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs("/static/" + templates.AUDIO)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/pages/grid_page.templ`, Line: 44, Col: 54}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></audio> <center><a class=\"mx-2\" href=\"/preview\">Show as Card</a> <a class=\"mx-2\" href=\"/preview?flat=1\">Show flat</a></center><div class=\"flex flex-wrap justify-center h-full\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for i, item := range templates.IMAGES {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				if !show_flat && len(templates.IMAGES)%2 == 0 {
					if i%2 == 0 {
						templ_7745c5c3_Var5 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
							templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
							if !templ_7745c5c3_IsBuffer {
								templ_7745c5c3_Buffer = templ.GetBuffer()
								defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<img class=\"absolute w-full h-full backface-hidden\" src=\"")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							var templ_7745c5c3_Var6 string
							templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs("/static/images/grid/" + templates.IMAGES[i+1])
							if templ_7745c5c3_Err != nil {
								return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/pages/grid_page.templ`, Line: 70, Col: 87}
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" loading=\"lazy\"> <img class=\"absolute w-full h-full backface-hidden rotate-y-180\" src=\"")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							var templ_7745c5c3_Var7 string
							templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs("/static/images/grid/" + item)
							if templ_7745c5c3_Err != nil {
								return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/pages/grid_page.templ`, Line: 75, Col: 68}
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" loading=\"lazy\">")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							if !templ_7745c5c3_IsBuffer {
								_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
							}
							return templ_7745c5c3_Err
						})
						templ_7745c5c3_Err = flipCard().Render(templ.WithChildren(ctx, templ_7745c5c3_Var5), templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					}
				} else {
					templ_7745c5c3_Var8 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
						templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
						if !templ_7745c5c3_IsBuffer {
							templ_7745c5c3_Buffer = templ.GetBuffer()
							defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<img class=\"absolute w-full h-full backface-hidden\" src=\"")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						var templ_7745c5c3_Var9 string
						templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs("/static/images/grid/" + item)
						if templ_7745c5c3_Err != nil {
							return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/pages/grid_page.templ`, Line: 86, Col: 68}
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" loading=\"lazy\"><div class=\"absolute w-full h-full backface-hidden rotate-y-180 break-words bg-gray-800\"><center><a href=\"")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						var templ_7745c5c3_Var10 templ.SafeURL = templ.SafeURL("/static/images/grid/" + item)
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var10)))
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" target=\"_blank\" rel=\"noopener noreferrer\">image </a> ")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, onClickHandler())
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button class=\"bg-gray-600 hover:bg-gray-800 text-white font-medium py-2 px-4 rounded\" onClick=\"")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						var templ_7745c5c3_Var11 templ.ComponentScript = onClickHandler()
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var11.Call)
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">Click me</button></center></div>")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						if !templ_7745c5c3_IsBuffer {
							_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
						}
						return templ_7745c5c3_Err
					})
					templ_7745c5c3_Err = flipCard().Render(templ.WithChildren(ctx, templ_7745c5c3_Var8), templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></main>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout.Page2("Grid | "+templates.PAGE_TITLE).Render(templ.WithChildren(ctx, templ_7745c5c3_Var3), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
