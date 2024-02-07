// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package order

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import (
	"bytes"
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/minnowo/astoryofand/util"
	"github.com/minnowo/astoryofand/view"
	"github.com/minnowo/astoryofand/view/components"
	"github.com/minnowo/astoryofand/view/layout"
)

func ShowOrderPage(boxPrice, stickerPrice float32) templ.Component {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>\n\n            #place_order_form > div ,\n            #place_order_form select,\n            #place_order_form input {\n                width: 100%;\n            }\n\n        </style> <form id=\"place_order_form\" class=\"hpad30p\" action=\"/order/place\" method=\"POST\"><input type=\"hidden\" name=\"boxpricevalue\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(util.F32TS(boxPrice)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"> <input type=\"hidden\" name=\"stickerpricevalue\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(util.F32TS(stickerPrice)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"> <center><h1>Order ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(view.TITLE)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/pages/order/order_page.templ`, Line: 30, Col: 38}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1></center><br><div><h3 class=\"no-margin\">What is your Email? </h3><input type=\"email\" name=\"email\" placeholder=\"Semple@Email.com\" required></div><br><div><h3 class=\"no-margin\">What is your Name? </h3><input type=\"text\" name=\"fullname\" placeholder=\"FirstName LastName\" required></div><br><div><h3 class=\"no-margin\">How many A Story Of And Box Sets do you want? </h3><p class=\"no-margin\">The price is ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = components.CodeHighlightFloat(boxPrice).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("&nbsp;per. </p><input type=\"number\" name=\"boxsetcount\" value=\"1\" min=\"0\" max=\"999\" required></div><br><div><h3 class=\"no-margin\">How many #VeryAsian Waterproof Vinyl Stickers do you want? </h3><p class=\"no-margin\">The price is ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = components.CodeHighlightFloat(stickerPrice).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("&nbsp;per. </p><input type=\"number\" name=\"cardsetcount\" value=\"1\" min=\"0\" max=\"999\" required></div><br><div><h3 class=\"no-margin\">How would you like to pay? </h3><select name=\"paymethod\" id=\"paymethod\" class=\"db\"><option value=\"cash\">Cash  </option> <option value=\"emt\">EMT   </option> <option value=\"other\">Other </option></select><div id=\"other_pay_row\" style=\"margin: 0.3rem;\"><label for=\"other\"><noscript>(If other) </noscript>Please specify:</label><br><textarea class=\"no-margin\" type=\"text\" id=\"otherPay\" name=\"otherpay\"></textarea></div></div><br><div><h3 class=\"no-margin\">How would you like to get your package? </h3><select name=\"deliverymethod\" id=\"deliverymethod\"><option value=\"delivery\">Delivery                               </option> <option value=\"pickup\">Pick Up                                </option> <option value=\"ship\">Ship to me (may incur additional costs)</option> <option value=\"other\">Other                                  </option></select><br><table style=\" margin: 0.3rem;\"><tr id=\"address_row\"><td><label for=\"address\">Address:</label></td><td><input type=\"text\" id=\"address\" name=\"address\"></td></tr><tr id=\"city_row\"><td><label for=\"city\">City:</label></td><td><input type=\"text\" id=\"city\" name=\"city\"></td></tr><tr id=\"zipcode_row\"><td><label for=\"zipcode\">Zip Code:</label></td><td><input type=\"text\" id=\"zipcode\" name=\"zipcode\"></td></tr></table><div id=\"other_row\" style=\"margin: 0.3rem;\"><label for=\"otherdelivery\"><noscript>(If other) </noscript>Please specify:</label><br><textarea type=\"text\" id=\"otherdelivery\" name=\"otherdelivery\"></textarea></div></div><div style=\"justify-content: end; display: flex;\"><input style=\"width:33%;\" type=\"submit\" value=\"Submit Order\" required></div></form><script>\n\n            function handleDeliveryMethodChange() {\n                const selectedValue = document.getElementById(\"deliverymethod\").value;\n\n                const addressInput = document.getElementById(\"address_row\");\n                const cityInput    = document.getElementById(\"city_row\");\n                const zipcodeInput = document.getElementById(\"zipcode_row\");\n                const otherInput   = document.getElementById(\"other_row\");\n\n                for (var j of addressInput.querySelectorAll(\"input\")) {\n                    j.required = selectedValue !== \"other\";\n                }\n                for (var j of cityInput.querySelectorAll(\"input\")) {\n                    j.required = selectedValue !== \"other\";\n                }\n                for (var j of zipcodeInput.querySelectorAll(\"input\")) {\n                    j.required = selectedValue === \"ship\";\n                }\n                for (var j of otherInput.querySelectorAll(\"textarea\")) {\n                    j.required = selectedValue === \"other\";\n                }\n\n                if (selectedValue === \"delivery\") {\n                    addressInput.style.display = \"inherit\";\n                    cityInput.style.display = \"inherit\";\n                    zipcodeInput.style.display = \"none\";\n                    otherInput.style.display = \"none\";\n                } \n                else if (selectedValue === \"ship\"){\n                    addressInput.style.display = \"inherit\";\n                    cityInput.style.display = \"inherit\";\n                    zipcodeInput.style.display = \"inherit\";\n                    otherInput.style.display = \"none\";\n                }\n                else if (selectedValue === \"pickup\"){\n                    addressInput.style.display = \"inherit\";\n                    cityInput.style.display = \"inherit\";\n                    zipcodeInput.style.display = \"none\";\n                    otherInput.style.display = \"none\";\n                }\n                else if (selectedValue === \"other\"){\n                    addressInput.style.display = \"none\";\n                    cityInput.style.display = \"none\";\n                    zipcodeInput.style.display = \"none\";\n                    otherInput.style.display = \"inherit\";\n                }\n            }\n\n            function handlePayMethodChange() {\n\n                const selectedValue = document.getElementById(\"paymethod\").value;\n\n                const otherPay = document.getElementById(\"other_pay_row\");\n\n                for (var j of otherPay.querySelectorAll(\"textarea\")) {\n                    j.required = selectedValue === \"other\";\n                }\n\n                if (selectedValue === \"other\") {\n                    otherPay.style.display = \"inherit\";\n                } else {\n                    otherPay.style.display = \"none\";\n                }\n            }\n\n            document.getElementById(\"deliverymethod\").addEventListener(\"change\", handleDeliveryMethodChange);\n            document.getElementById(\"paymethod\").addEventListener(\"change\", handlePayMethodChange);\n            handleDeliveryMethodChange();\n            handlePayMethodChange();\n\n\n        </script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout.Page("Order | "+view.PAGE_TITLE).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
