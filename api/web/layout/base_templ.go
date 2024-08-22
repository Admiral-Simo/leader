// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "server/store"
import "time"
import "strconv"

func Base(title string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><link rel=\"icon\" href=\"/public/favicon.ico\" type=\"image/x-icon\"><title>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `api/web/layout/base.templ`, Line: 14, Col: 17}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><!-- Tailwind CSS CDN --><script src=\"https://cdn.tailwindcss.com\"></script><!-- AlpineJS CDN --><script defer src=\"https://cdn.jsdelivr.net/npm/alpinejs@3.14.1/dist/cdn.min.js\"></script><!-- HTMX CDN --><script src=\"https://unpkg.com/htmx.org@2.0.2\" integrity=\"sha384-Y7hw+L/jvKeWIRRkqWYfPcvVxHzVzn5REgzbawhxAuQGwX1XWe70vji+VSeHOThJ\" crossorigin=\"anonymous\"></script></head><body class=\"bg-gradient-to-r from-indigo-100 to-purple-100\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Navbar().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"relative min-h-screen\"><!-- Spinner Overlay --><div id=\"spinner-overlay\" class=\"absolute inset-0 bg-black bg-opacity-75 flex items-center justify-center hidden\"><div class=\"flex space-x-2 justify-center items-center h-screen dark:invert\"><span class=\"sr-only\">Loading...</span><div class=\"h-8 w-8 bg-black rounded-full animate-bounce [animation-delay:-0.3s]\"></div><div class=\"h-8 w-8 bg-black rounded-full animate-bounce [animation-delay:-0.15s]\"></div><div class=\"h-8 w-8 bg-black rounded-full animate-bounce\"></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Footer().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script>\n        // HTMX event listeners to show/hide the spinner\n        document.addEventListener('htmx:configRequest', function () {\n            document.getElementById('spinner-overlay').classList.remove('hidden');\n        });\n\n        document.addEventListener('htmx:afterSwap', function () {\n            document.getElementById('spinner-overlay').classList.add('hidden');\n        });\n    </script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func Navbar() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<nav class=\"bg-gradient-to-r from-indigo-600 to-purple-700 p-4 shadow-lg\" x-data=\"{ open: false, showLogout: false }\"><div class=\"container mx-auto flex justify-between items-center\"><!-- Brand --><a hx-get=\"/\" hx-target=\"body\" class=\"text-white font-extrabold text-2xl tracking-wider cursor-pointer\">EleadGen</a><!-- Desktop Menu --><div class=\"hidden md:flex space-x-8 items-center\"><a hx-get=\"/about\" hx-target=\"body\" class=\"text-white text-lg hover:text-gray-200 transition duration-300 cursor-pointer\">About</a> <a hx-get=\"/services\" hx-target=\"body\" class=\"text-white text-lg hover:text-gray-200 transition duration-300 cursor-pointer\">Services</a> <a hx-get=\"/contact\" hx-target=\"body\" class=\"text-white text-lg hover:text-gray-200 transition duration-300 cursor-pointer\">Contact</a> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if user, ok := ctx.Value("user").(store.User); ok {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Profile Button --> <div @click=\"showLogout = !showLogout\" class=\"relative flex items-center space-x-3 cursor-pointer\"><!-- Avatar --><div class=\"flex items-center justify-center bg-gradient-to-r from-purple-600 to-indigo-600 text-white text-3xl font-bold rounded-full shadow-lg w-14 h-14\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(string(user.Name[0]))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `api/web/layout/base.templ`, Line: 85, Col: 29}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><!-- Logout Button --><div x-show=\"showLogout\" x-transition:enter=\"transition transform ease-out duration-300\" x-transition:enter-start=\"translate-y-4 opacity-0\" x-transition:enter-end=\"translate-y-0 opacity-100\" x-transition:leave=\"transition transform ease-in duration-300\" x-transition:leave-start=\"translate-y-0 opacity-100\" x-transition:leave-end=\"translate-y-4 opacity-0\" class=\"absolute top-10 left-0 mt-2 bg-white text-indigo-600 rounded-md shadow-lg p-2 flex flex-col items-center\"><a hx-get=\"/logout\" hx-target=\"body\" class=\"block px-4 py-2 text-indigo-600 hover:bg-gray-100 transition duration-300 cursor-pointer\">Logout</a></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a hx-get=\"/login\" hx-target=\"body\" class=\"text-white text-lg hover:text-gray-200 transition duration-300 cursor-pointer border border-white px-4 py-2 rounded-md\">Login</a> <a hx-get=\"/signup\" hx-target=\"body\" class=\"bg-white text-indigo-600 font-semibold text-lg px-4 py-2 rounded-md shadow hover:bg-gray-100 transition duration-300 cursor-pointer\">Signup</a>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><!-- Mobile Menu Button --><div class=\"md:hidden\"><button @click=\"open = !open\" aria-expanded=\"open\" aria-controls=\"mobile-menu\" class=\"text-white focus:outline-none\"><svg xmlns=\"http://www.w3.org/2000/svg\" class=\"h-8 w-8\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4 6h16M4 12h16m-7 6h7\"></path></svg></button></div></div><!-- Mobile Menu --><div x-show=\"open\" x-transition @click.away=\"open = false\" id=\"mobile-menu\" class=\"md:hidden mt-4 bg-white rounded-lg shadow-lg\"><a hx-get=\"/about\" hx-target=\"body\" class=\"block px-4 py-3 text-indigo-600 hover:bg-gray-100 transition duration-300 cursor-pointer\">About</a> <a hx-get=\"/services\" hx-target=\"body\" class=\"block px-4 py-3 text-indigo-600 hover:bg-gray-100 transition duration-300 cursor-pointer\">Services</a> <a hx-get=\"/contact\" hx-target=\"body\" class=\"block px-4 py-3 text-indigo-600 hover:bg-gray-100 transition duration-300 cursor-pointer\">Contact</a> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if user, ok := ctx.Value("user").(store.User); ok {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Mobile Profile Button --> <div @click=\"showLogout = !showLogout\" class=\"relative flex items-center space-x-3 px-4 py-3 cursor-pointer\"><!-- Avatar --><div class=\"flex items-center justify-center bg-gradient-to-r from-purple-600 to-indigo-600 text-white text-3xl font-bold rounded-full shadow-lg w-14 h-14\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(string(user.Name[0]))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `api/web/layout/base.templ`, Line: 168, Col: 28}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><!-- Mobile Logout Button --><div x-show=\"showLogout\" x-transition:enter=\"transition transform ease-out duration-300\" x-transition:enter-start=\"translate-y-4 opacity-0\" x-transition:enter-end=\"translate-y-0 opacity-100\" x-transition:leave=\"transition transform ease-in duration-300\" x-transition:leave-start=\"translate-y-0 opacity-100\" x-transition:leave-end=\"translate-y-4 opacity-0\" class=\"absolute top-10 left-0 mt-2 bg-white text-indigo-600 rounded-md shadow-lg p-2 flex flex-col items-center\"><a hx-get=\"/logout\" hx-target=\"body\" class=\"block px-4 py-2 text-indigo-600 hover:bg-gray-100 transition duration-300 cursor-pointer\">Logout</a></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a hx-get=\"/login\" hx-target=\"body\" class=\"block px-4 py-3 text-indigo-600 hover:bg-gray-100 transition duration-300 cursor-pointer\">Login</a> <a hx-get=\"/signup\" hx-target=\"body\" class=\"block px-4 py-3 bg-gradient-to-r from-indigo-600 to-purple-700 text-white font-semibold text-lg rounded-md shadow hover:bg-indigo-700 transition duration-300 cursor-pointer\">Signup</a>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></nav>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func Footer() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var6 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var6 == nil {
			templ_7745c5c3_Var6 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<footer class=\"bg-gradient-to-r from-indigo-600 to-purple-700 text-white py-6\"><div class=\"container mx-auto text-center\"><p class=\"text-sm\">&copy; ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(time.Now().Year()))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `api/web/layout/base.templ`, Line: 207, Col: 62}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" EleadGen. All rights reserved.</p><p class=\"text-xs mt-2\">Built with ❤️ using Tailwind CSS & AlpineJS</p></div></footer>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
